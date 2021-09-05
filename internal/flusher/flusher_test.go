package flusher_test

import (
	"context"
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
		ctx      context.Context

		songs = []models.Song{
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

		chunkSize   int
		input       []models.Song
		failedSongs []models.Song
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mockRepo = mocks.NewMockRepo(ctrl)

		ctx = context.Background()
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	JustBeforeEach(func() {
		fls = flusher.NewFlusher(chunkSize, mockRepo)
		failedSongs = fls.Flush(ctx, input)
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
				mockRepo.EXPECT().AddSongs(ctx, songs).Return(int64(0), nil).Times(1)
			})

			It("should return empty", func() {
				Expect(failedSongs).To(BeNil())
			})
		})

		Context("with a chunk size equals to songs size", func() {
			BeforeEach(func() {
				chunkSize = len(songs)
				mockRepo.EXPECT().AddSongs(ctx, songs).Return(int64(0), nil).Times(1)
			})

			It("should return empty", func() {
				Expect(failedSongs).To(BeNil())
			})
		})

		Context("with a chunk size less than songs size", func() {
			BeforeEach(func() {
				chunkSize = len(songs) / 3

				gomock.InOrder(
					mockRepo.EXPECT().AddSongs(ctx, songs[chunkSize*0:chunkSize*0+chunkSize]).Return(int64(0), nil),
					mockRepo.EXPECT().AddSongs(ctx, songs[chunkSize*1:chunkSize*1+chunkSize]).Return(int64(0), nil),
					mockRepo.EXPECT().AddSongs(ctx, songs[chunkSize*2:chunkSize*2+chunkSize]).Return(int64(0), nil),
					mockRepo.EXPECT().AddSongs(ctx, songs[chunkSize*3:]).Return(int64(0), nil),
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
					mockRepo.EXPECT().AddSongs(ctx, songs[chunkSize*0:chunkSize*0+chunkSize]).Return(int64(0), errors.New("whatever")),
					mockRepo.EXPECT().AddSongs(ctx, songs[chunkSize*1:chunkSize*1+chunkSize]).Return(int64(0), nil),
					mockRepo.EXPECT().AddSongs(ctx, songs[chunkSize*2:chunkSize*2+chunkSize]).Return(int64(0), nil),
					mockRepo.EXPECT().AddSongs(ctx, songs[chunkSize*3:]).Return(int64(0), nil),
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
						mockRepo.EXPECT().AddSongs(ctx, songs[chunkSize*0:chunkSize*0+chunkSize]).Return(int64(0), nil),
						mockRepo.EXPECT().AddSongs(ctx, songs[chunkSize*1:chunkSize*1+chunkSize]).Return(int64(0), errors.New("whatever")),
						mockRepo.EXPECT().AddSongs(ctx, songs[chunkSize*2:chunkSize*2+chunkSize]).Return(int64(0), nil),
						mockRepo.EXPECT().AddSongs(ctx, songs[chunkSize*3:]).Return(int64(0), nil),
					)
				})

				It("should return second chunk", func() {
					Expect(failedSongs).To(BeEquivalentTo(songs[chunkSize : chunkSize*2]))
				})
			})

			Context("on the second chunk", func() {
				BeforeEach(func() {
					gomock.InOrder(
						mockRepo.EXPECT().AddSongs(ctx, songs[chunkSize*0:chunkSize*0+chunkSize]).Return(int64(0), nil),
						mockRepo.EXPECT().AddSongs(ctx, songs[chunkSize*1:chunkSize*1+chunkSize]).Return(int64(0), errors.New("whatever")),
						mockRepo.EXPECT().AddSongs(ctx, songs[chunkSize*2:chunkSize*2+chunkSize]).Return(int64(0), nil),
						mockRepo.EXPECT().AddSongs(ctx, songs[chunkSize*3:]).Return(int64(0), nil),
					)
				})

				It("should return second chunk", func() {
					Expect(failedSongs).To(BeEquivalentTo(songs[chunkSize : chunkSize*2]))
				})
			})

			Context("on the first and the third chunks", func() {
				BeforeEach(func() {
					gomock.InOrder(
						mockRepo.EXPECT().AddSongs(ctx, songs[chunkSize*0:chunkSize*0+chunkSize]).Return(int64(0), errors.New("whatever")),
						mockRepo.EXPECT().AddSongs(ctx, songs[chunkSize*1:chunkSize*1+chunkSize]).Return(int64(0), nil),
						mockRepo.EXPECT().AddSongs(ctx, songs[chunkSize*2:chunkSize*2+chunkSize]).Return(int64(0), errors.New("another whatever")),
						mockRepo.EXPECT().AddSongs(ctx, songs[chunkSize*3:]).Return(int64(0), nil),
					)
				})

				It("should return first and third chunks concatenated", func() {
					Expect(failedSongs).To(BeEquivalentTo(append(songs[:chunkSize], songs[chunkSize*2:chunkSize*3]...)))
				})
			})

			Context("on the all chunks", func() {
				BeforeEach(func() {
					gomock.InOrder(
						mockRepo.EXPECT().AddSongs(ctx, songs[chunkSize*0:chunkSize*0+chunkSize]).Return(int64(0), errors.New("whatever")),
						mockRepo.EXPECT().AddSongs(ctx, songs[chunkSize*1:chunkSize*1+chunkSize]).Return(int64(0), errors.New("whatever")),
						mockRepo.EXPECT().AddSongs(ctx, songs[chunkSize*2:chunkSize*2+chunkSize]).Return(int64(0), errors.New("another whatever")),
						mockRepo.EXPECT().AddSongs(ctx, songs[chunkSize*3:]).Return(int64(0), errors.New("whatever")),
					)
				})

				It("should return all songs", func() {
					Expect(failedSongs).To(BeEquivalentTo(songs))
				})
			})
		})
	})
})
