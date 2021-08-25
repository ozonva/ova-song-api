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

	const (
		capacity = 10
	)

	var (
		ctrl        *gomock.Controller
		mockFlusher *mocks.MockFlusher

		backend         saverBackend
		songChan        chan Song
		closeChan       chan struct{}
		sendCloseSignal func()

		someSong = Song{Id: 100, Author: "100", Name: "100", Year: 100}
		_        = someSong // TODO: remove me
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

		songChan = make(chan Song, capacity)
		closeChan = make(chan struct{})

		period := time.Hour

		backend = saverBackend{
			flusher:     mockFlusher,
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
		wg.Wait() // waiting for the backend to complete before calling `Finish` on mock controller
		ctrl.Finish()
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
		AfterEach(func() {
			sendCloseSignal()
		})

		It("returns", func() {})
	})

	When("one song received", func() {
		BeforeEach(func() {
			mockFlusher.EXPECT().Flush([]Song{someSong}).AnyTimes() // todo: here
		})

		AfterEach(func() {
			sendCloseSignal()
		})

		It("can receive song", func() {
			Eventually(songChan).Should(BeSent(someSong))
		})
	})

	Context("some songs arrived", func() {
		BeforeEach(func() {
			for i := 0; i < n; i++ {
				songChan <- songs[i]
			}
		})

		When("close signals arrived", func() {
			BeforeEach(func() {
				println("setting expectations")
				mockFlusher.EXPECT().Flush(songs[:n]).Times(1)

				sendCloseSignal()
			})

			It("will flush all songs", func() {
				// verified through flusher mock
			})

			It("will stop the timer", func() {
				// todo
			})
		})

		When("the total number of songs exceeds capacity", func() {
			AfterEach(func() {
				// note that there are already n songs in the buffer
				for i := n; i < capacity; i++ {
					songChan <- songs[i]
				}

				sendCloseSignal()
			})

			It("should save the `capacity` of songs, than flush the remaining", func() {
				mockFlusher.EXPECT().Flush(songs).Times(1)
			})
		})
	})

	//mockFlusher.EXPECT().Flush(gomock.Any()).DoAndReturn(func(songs []Song) interface{} {
	//	fmt.Printf("flusher called with: %v", songs)
	//	return nil
	//}).MinTimes(10)
	//
	//saver := saver.NewSaver(10, time.Second, mockFlusher)
	//
	//saver.Save(Song{Id: 1, Author: "1", Name: "1", Year: 1})
	//saver.Save(Song{Id: 2, Author: "1", Name: "1", Year: 1})
	//saver.Save(Song{Id: 3, Author: "1", Name: "1", Year: 1})
	//time.Sleep(time.Second * 3)
	//saver.Save(Song{Id: 4, Author: "1", Name: "1", Year: 1})
	//time.Sleep(time.Second * 3)
})
