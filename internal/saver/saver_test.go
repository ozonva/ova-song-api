package saver

import (
	"context"
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
		ctx       context.Context

		someSong = Song{Id: 1, Author: "1", Name: "1", Year: 1}
	)

	BeforeEach(func() {
		songChan = make(chan Song, capacity)
		closeChan = make(chan struct{})
		ctx = context.Background()

		svr = &saver{songChan, closeChan}
	})

	When("Save() called", func() {
		BeforeEach(func() {
			svr.Save(ctx, someSong)
		})

		It("will push song into the songs channel", func() {
			Expect(songChan).To(Receive(Equal(someSong)))
		})
	})

	When("Context cancelled while Save() called", func() {
		var (
			cancelContext context.Context
			cancelFun     context.CancelFunc
			unblocked     chan struct{}
		)

		BeforeEach(func() {
			cancelContext, cancelFun = context.WithCancel(ctx)
			unblocked = make(chan struct{})

			// should block so called in separate goroutine
			go func() {
				defer GinkgoRecover()
				for i := 0; i < capacity+1; i++ {
					svr.Save(cancelContext, someSong)
				}

				close(unblocked)
			}()
		})

		It("will cancel saving the last song", func() {
			Eventually(songChan).Should(HaveLen(capacity))
			cancelFun()
			Eventually(unblocked).Should(BeClosed())
		})
	})

	When("Close() called", func() {
		BeforeEach(func() {
			// called in a separate goroutine to avoid blocking due to using an unbuffered channel
			go func() {
				defer GinkgoRecover()
				svr.Close()
			}()
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
			{Id: 1, Author: "Author 1", Name: "Name 1", Year: 2001},
			{Id: 2, Author: "Author 2", Name: "Name 2", Year: 2002},
			{Id: 3, Author: "Author 3", Name: "Name 3", Year: 2003},
			{Id: 4, Author: "Author 4", Name: "Name 4", Year: 2004},
			{Id: 5, Author: "Author 5", Name: "Name 5", Year: 2005},
			{Id: 6, Author: "Author 6", Name: "Name 6", Year: 2006},
			{Id: 7, Author: "Author 7", Name: "Name 7", Year: 2007},
			{Id: 8, Author: "Author 8", Name: "Name 8", Year: 2008},
			{Id: 9, Author: "Author 9", Name: "Name 9", Year: 2009},
			{Id: 10, Author: "Author 10", Name: "Name 10", Year: 2010},
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
			mockFlusher.EXPECT().Flush(notNil(), []Song{someSong}).Times(1)
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
				mockFlusher.EXPECT().Flush(notNil(), songs[:n]).Times(1)
			})
		})

		When("the ticker does tick", func() {
			AfterEach(func() {
				sendTick()
			})

			It("will flush contained songs and reset the ticker", func() {
				mockFlusher.EXPECT().Flush(notNil(), songs[:n]).Times(1)
				mockTicker.EXPECT().Reset(period).Times(1)
			})
		})

		Context("flusher will fail to save the second and the fourth songs", func() {
			BeforeEach(func() {
				mockFlusher.EXPECT().Flush(notNil(), songs[:n]).Return([]Song{songs[1], songs[3]}).Times(1)
			})

			When("the ticker does tick", func() {
				AfterEach(func() {
					sendTick()
				})

				It("will flush contained songs, reset the ticker and flush remaining songs later", func() {
					mockFlusher.EXPECT().Flush(notNil(), []Song{songs[1], songs[3]}).Times(1)
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
				mockFlusher.EXPECT().Flush(notNil(), expected).Times(1)
			})

			When("the ticker does tick", func() {
				AfterEach(func() {
					sendTick()
				})

				It("will flush contained songs, and reset the ticker", func() {
					mockFlusher.EXPECT().Flush(notNil(), expected).Times(1)
					mockTicker.EXPECT().Reset(period).Times(1)
				})
			})
		})
	})
})

// notNil matcher for gomock
func notNil() gomock.Matcher {
	return gomock.Not(gomock.Nil())
}
