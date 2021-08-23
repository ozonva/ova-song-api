package flusher_test

import (
	"errors"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/ozonva/ova-song-api/internal/flusher"
	"github.com/ozonva/ova-song-api/internal/mocks"
	"github.com/ozonva/ova-song-api/internal/models"
)

var _ = Describe("Flusher", func() {

	var (
		ctrl     *gomock.Controller
		mockRepo *mocks.MockRepo
		fls      flusher.Flusher

		songs = []models.Song{
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

		chunkSize   int
		input       []models.Song
		failedSongs []models.Song
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mockRepo = mocks.NewMockRepo(ctrl)
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	JustBeforeEach(func() {
		fls = flusher.NewFlusher(chunkSize, mockRepo)
		failedSongs = fls.Flush(input)
	})

	Describe("with nil input", func() {
		BeforeEach(func() {
			chunkSize = 10
			input = nil
		})

		It("should return empty", func() {
			Expect(failedSongs).To(BeNil())
		})
	})

	Describe("with empty input", func() {
		BeforeEach(func() {
			chunkSize = 10
			input = []models.Song{}
		})

		It("should return empty", func() {
			Expect(failedSongs).To(BeNil())
		})
	})

	Describe("with normal input", func() {
		BeforeEach(func() {
			input = songs
		})

		Context("with negative chunk size", func() {
			BeforeEach(func() {
				chunkSize = -10
			})

			It("should return all songs", func() {
				Expect(failedSongs).To(BeEquivalentTo(songs))
			})
		})

		Context("with zero chunk size", func() {
			BeforeEach(func() {
				chunkSize = 0
			})

			It("should return all songs", func() {
				Expect(failedSongs).To(BeEquivalentTo(songs))
			})
		})

		Context("with a chunk size greater than songs size", func() {
			BeforeEach(func() {
				chunkSize = len(songs) + 5
				mockRepo.EXPECT().AddSongs(songs).Return(nil).Times(1)
			})

			It("should return empty", func() {
				Expect(failedSongs).To(BeNil())
			})
		})

		Context("with a chunk size equals to songs size", func() {
			BeforeEach(func() {
				chunkSize = len(songs)
				mockRepo.EXPECT().AddSongs(songs).Return(nil).Times(1)
			})

			It("should return empty", func() {
				Expect(failedSongs).To(BeNil())
			})
		})

		Context("with a chunk size less than songs size", func() {
			BeforeEach(func() {
				chunkSize = len(songs) / 3

				gomock.InOrder(
					mockRepo.EXPECT().AddSongs(songs[chunkSize*0:chunkSize*0+chunkSize]).Return(nil),
					mockRepo.EXPECT().AddSongs(songs[chunkSize*1:chunkSize*1+chunkSize]).Return(nil),
					mockRepo.EXPECT().AddSongs(songs[chunkSize*2:chunkSize*2+chunkSize]).Return(nil),
					mockRepo.EXPECT().AddSongs(songs[chunkSize*3:]).Return(nil),
				)
			})

			It("should return empty", func() {
				Expect(failedSongs).To(BeNil())
			})
		})

		Context("when repo fails on the first chunk", func() {
			BeforeEach(func() {
				chunkSize = len(songs) / 3

				gomock.InOrder(
					mockRepo.EXPECT().AddSongs(songs[chunkSize*0:chunkSize*0+chunkSize]).Return(errors.New("whatever")),
					mockRepo.EXPECT().AddSongs(songs[chunkSize*1:chunkSize*1+chunkSize]).Return(nil),
					mockRepo.EXPECT().AddSongs(songs[chunkSize*2:chunkSize*2+chunkSize]).Return(nil),
					mockRepo.EXPECT().AddSongs(songs[chunkSize*3:]).Return(nil),
				)
			})

			It("should return first chunk", func() {
				Expect(failedSongs).To(BeEquivalentTo(songs[:chunkSize]))
			})
		})

		Describe("when repo fails", func() {
			BeforeEach(func() {
				chunkSize = len(songs) / 3
			})

			Context("on the second chunk", func() {
				BeforeEach(func() {
					gomock.InOrder(
						mockRepo.EXPECT().AddSongs(songs[chunkSize*0:chunkSize*0+chunkSize]).Return(nil),
						mockRepo.EXPECT().AddSongs(songs[chunkSize*1:chunkSize*1+chunkSize]).Return(errors.New("whatever")),
						mockRepo.EXPECT().AddSongs(songs[chunkSize*2:chunkSize*2+chunkSize]).Return(nil),
						mockRepo.EXPECT().AddSongs(songs[chunkSize*3:]).Return(nil),
					)
				})

				It("should return second chunk", func() {
					Expect(failedSongs).To(BeEquivalentTo(songs[chunkSize : chunkSize*2]))
				})
			})

			Context("on the second chunk", func() {
				BeforeEach(func() {
					gomock.InOrder(
						mockRepo.EXPECT().AddSongs(songs[chunkSize*0:chunkSize*0+chunkSize]).Return(nil),
						mockRepo.EXPECT().AddSongs(songs[chunkSize*1:chunkSize*1+chunkSize]).Return(errors.New("whatever")),
						mockRepo.EXPECT().AddSongs(songs[chunkSize*2:chunkSize*2+chunkSize]).Return(nil),
						mockRepo.EXPECT().AddSongs(songs[chunkSize*3:]).Return(nil),
					)
				})

				It("should return second chunk", func() {
					Expect(failedSongs).To(BeEquivalentTo(songs[chunkSize : chunkSize*2]))
				})
			})

			Context("on the first and the third chunks", func() {
				BeforeEach(func() {
					gomock.InOrder(
						mockRepo.EXPECT().AddSongs(songs[chunkSize*0:chunkSize*0+chunkSize]).Return(errors.New("whatever")),
						mockRepo.EXPECT().AddSongs(songs[chunkSize*1:chunkSize*1+chunkSize]).Return(nil),
						mockRepo.EXPECT().AddSongs(songs[chunkSize*2:chunkSize*2+chunkSize]).Return(errors.New("another whatever")),
						mockRepo.EXPECT().AddSongs(songs[chunkSize*3:]).Return(nil),
					)
				})

				It("should return first and third chunks concatenated", func() {
					Expect(failedSongs).To(BeEquivalentTo(append(songs[:chunkSize], songs[chunkSize*2:chunkSize*3]...)))
				})
			})

			Context("on the all chunks", func() {
				BeforeEach(func() {
					gomock.InOrder(
						mockRepo.EXPECT().AddSongs(songs[chunkSize*0:chunkSize*0+chunkSize]).Return(errors.New("whatever")),
						mockRepo.EXPECT().AddSongs(songs[chunkSize*1:chunkSize*1+chunkSize]).Return(errors.New("whatever")),
						mockRepo.EXPECT().AddSongs(songs[chunkSize*2:chunkSize*2+chunkSize]).Return(errors.New("another whatever")),
						mockRepo.EXPECT().AddSongs(songs[chunkSize*3:]).Return(errors.New("whatever")),
					)
				})

				It("should return all songs", func() {
					Expect(failedSongs).To(BeEquivalentTo(songs))
				})
			})
		})
	})
})
