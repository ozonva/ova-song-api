package saver

import (
	"sync"
	"time"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/ozonva/ova-song-api/internal/mocks"
	. "github.com/ozonva/ova-song-api/internal/models"
)

const (
	capacity = 7
	n        = 5
)

var _ = Describe("Saver", func() {

	var (
		svr       Saver
		songChan  chan Song
		closeChan chan struct{}

		someSong = Song{Id: 1, Author: "1", Name: "1", Year: 1}
	)

	BeforeEach(func() {
		songChan = make(chan Song, capacity)
		closeChan = make(chan struct{})

		svr = &saver{songChan, closeChan}
	})

	When("Save() called", func() {
		BeforeEach(func() {
			svr.Save(someSong)
		})

		It("will push song into the songs channel", func() {
			Expect(songChan).To(Receive(Equal(someSong)))
		})
	})

	When("Close() called", func() {

		BeforeEach(func() {
			defer GinkgoRecover()
			// called in a separate goroutine to avoid blocking due to using an unbuffered channel
			go svr.Close()
		})

		It("will close songs and `close` channel", func() {
			Eventually(closeChan).Should(BeClosed())
			Eventually(songChan).Should(BeClosed())
		})
	})
})

var _ = Describe("SaverBackend", func() {

	var (
		ctrl        *gomock.Controller
		mockFlusher *mocks.MockFlusher
		mockTicker  *mocks.MockTicker
		tickerChan  chan time.Time
		sendTick    func()

		backend         saverBackend
		songChan        chan Song
		closeChan       chan struct{}
		sendCloseSignal func()

		period   = time.Hour // whatever
		someSong = Song{Id: 100, Author: "100", Name: "100", Year: 100}
		songs    = []Song{
			{1, "Author 1", "Name 1", 2001},
			{2, "Author 2", "Name 2", 2002},
			{3, "Author 3", "Name 3", 2003},
			{4, "Author 4", "Name 4", 2004},
			{5, "Author 5", "Name 5", 2005},
			{6, "Author 6", "Name 6", 2006},
			{7, "Author 7", "Name 7", 2007},
			{8, "Author 8", "Name 8", 2008},
			{9, "Author 9", "Name 9", 2009},
			{10, "Author 10", "Name 10", 2010},
		}

		// we have to wait for the backend to complete in order to be sure that the gomock can record all calls
		wg sync.WaitGroup
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mockFlusher = mocks.NewMockFlusher(ctrl)

		tickerChan = make(chan time.Time)
		mockTicker = mocks.NewMockTicker(ctrl)
		mockTicker.EXPECT().C().Return(tickerChan).AnyTimes()
		sendTick = func() {
			tickerChan <- time.Time{}
		}

		songChan = make(chan Song)      // unbuffered to make tests more deterministic
		closeChan = make(chan struct{}) // unbuffered

		backend = saverBackend{
			flusher:     mockFlusher,
			ticker:      mockTicker,
			songsBuffer: make([]Song, 0, capacity),
			period:      period,
			songChan:    songChan,
			closeChan:   closeChan,
		}

		wg = sync.WaitGroup{}
		sendCloseSignal = func() {
			close(songChan)
			close(closeChan)
		}
	})

	AfterEach(func() {
		sendCloseSignal()
	})

	AfterEach(func() {
		wg.Wait() // waiting for the backend to complete before calling `Finish` on mock controller
		ctrl.Finish()
		close(tickerChan)
	})

	BeforeEach(func() {
		wg.Add(1)
		go func() {
			defer GinkgoRecover()
			defer wg.Done()

			backend.Serve()
		}()
	})

	When("closed without songs", func() {
		It("will close ticker", func() {
			mockTicker.EXPECT().Stop().Times(1)
		})
	})

	When("one song received", func() {
		BeforeEach(func() {
			mockFlusher.EXPECT().Flush([]Song{someSong}).Times(1)
		})

		AfterEach(func() {
			mockTicker.EXPECT().Stop().Times(1)
		})

		It("can receive song", func() {
			Eventually(songChan).Should(BeSent(someSong))
		})
	})

	When("some songs arrived", func() {
		BeforeEach(func() {
			for i := 0; i < n; i++ {
				songChan <- songs[i]
			}
		})

		AfterEach(func() {
			mockTicker.EXPECT().Stop().Times(1)
		})

		When("the close signals arrived", func() {
			It("will flush all songs and stop the timer", func() {
				mockFlusher.EXPECT().Flush(songs[:n]).Times(1)
			})
		})

		When("the ticker does tick", func() {
			AfterEach(func() {
				sendTick()
			})

			It("will flush contained songs and reset the ticker", func() {
				mockFlusher.EXPECT().Flush(songs[:n]).Times(1)
				mockTicker.EXPECT().Reset(period).Times(1)
			})
		})

		Context("flusher will fail to save the second and the fourth songs", func() {
			BeforeEach(func() {
				mockFlusher.EXPECT().Flush(songs[:n]).Return([]Song{songs[1], songs[3]}).Times(1)
			})

			When("the ticker does tick", func() {
				AfterEach(func() {
					sendTick()
				})

				It("will flush contained songs, reset the ticker and flush remaining songs later", func() {
					mockFlusher.EXPECT().Flush([]Song{songs[1], songs[3]}).Times(1)
					mockTicker.EXPECT().Reset(period).Times(1)
				})
			})
		})

		When("the total number of songs exceeds capacity", func() {
			var expected []Song
			BeforeEach(func() {
				expected = append([]Song{}, songs[:capacity]...)
				expected = append(expected, someSong)

				// note that there are already n songs in the buffer
				for i := n; i < capacity; i++ {
					songChan <- songs[i]
				}
				songChan <- someSong
			})

			It("should save the all the songs(when closed)", func() {
				mockFlusher.EXPECT().Flush(expected).Times(1)
			})

			When("the ticker does tick", func() {
				AfterEach(func() {
					sendTick()
				})

				It("will flush contained songs, and reset the ticker", func() {
					mockFlusher.EXPECT().Flush(expected).Times(1)
					mockTicker.EXPECT().Reset(period).Times(1)
				})
			})
		})
	})
})
