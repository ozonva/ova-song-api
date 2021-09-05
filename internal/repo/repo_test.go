package repo_test

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/ozonva/ova-song-api/internal/models"
	rp "github.com/ozonva/ova-song-api/internal/repo"
)

var _ = Describe("Repo", func() {
	var (
		db     *sql.DB
		sqlxDB *sqlx.DB
		mock   sqlmock.Sqlmock
		ctx    context.Context

		repo rp.Repo

		songs = []models.Song{
			{1, "Author 1", "Name 1", 2001},
			{2, "Author 2", "Name 2", 2002},
			{3, "Author 3", "Name 3", 2003},
			{4, "Author 4", "Name 4", 2004},
		}
		someSong = songs[0]
	)

	BeforeEach(func() {
		var err error
		db, mock, err = sqlmock.New()
		Expect(err).To(BeNil())
		sqlxDB = sqlx.NewDb(db, "sqlmock")
		ctx = context.Background()

		repo = rp.NewRepo(sqlxDB)
	})

	AfterEach(func() {
		mock.ExpectClose()
		err := db.Close()
		Expect(err).To(BeNil())
	})

	Describe("Add song", func() {
		var expectedRowId int64

		BeforeEach(func() {
			expectedRowId = 7
			mock.ExpectQuery("INSERT INTO songs").
				WithArgs(songs[0].Name, songs[0].Author, songs[0].Year).
				WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(expectedRowId))
		})

		It("Should succeed and return correct id", func() {
			song := songs[0]
			newId, err := repo.AddSong(ctx, song)
			Expect(err).To(BeNil())
			Expect(newId).To(Equal(expectedRowId))
		})
	})

	Describe("Add songs", func() {
		BeforeEach(func() {
			query := mock.ExpectQuery("INSERT INTO songs").
				WithArgs(
					songs[0].Name, songs[0].Author, songs[0].Year,
					songs[1].Name, songs[1].Author, songs[1].Year,
					songs[2].Name, songs[2].Author, songs[2].Year,
					songs[3].Name, songs[3].Author, songs[3].Year,
				)
			query.WillReturnRows(sqlmock.
				NewRows([]string{"id"}).
				AddRow(1).AddRow(2).AddRow(3).AddRow(4),
			)
		})

		It("Should succeed and return last inserted id", func() {
			id, err := repo.AddSongs(ctx, songs)
			Expect(err).To(BeNil())
			Expect(id).To(Equal(int64(4)))
		})
	})

	Describe("Describe song", func() {
		BeforeEach(func() {
			mock.ExpectQuery("SELECT id, name, author, year FROM songs WHERE").
				WithArgs(someSong.Id).
				WillReturnRows(sqlmock.
					NewRows([]string{"id", "name", "author", "year"}).
					AddRow(someSong.Id, someSong.Name, someSong.Author, someSong.Year))
		})

		It("Should succeed and return correct song", func() {
			song, err := repo.DescribeSong(ctx, someSong.Id)
			Expect(err).To(BeNil())
			Expect(song).To(BeEquivalentTo(&someSong))
		})
	})

	Describe("Update song", func() {
		var rowsAffected int64

		JustBeforeEach(func() {
			mock.ExpectExec("UPDATE songs SET").
				WithArgs(someSong.Name, someSong.Author, someSong.Year, someSong.Id).
				WillReturnResult(sqlmock.NewResult(0, rowsAffected))
		})

		Context("the song to be updated is present", func() {
			BeforeEach(func() {
				rowsAffected = 1
			})

			It("Should succeed and return true", func() {
				succeed, err := repo.UpdateSong(ctx, someSong)
				Expect(err).To(BeNil())
				Expect(succeed).To(BeTrue())
			})
		})

		Context("the song to be updated is absent", func() {
			BeforeEach(func() {
				rowsAffected = 0
			})

			It("Should succeed and return false", func() {
				succeed, err := repo.UpdateSong(ctx, someSong)
				Expect(err).To(BeNil())
				Expect(succeed).To(BeFalse())
			})
		})
	})

	Describe("List songs", func() {
		const (
			limit  uint64 = 2
			offset uint64 = 1
		)

		BeforeEach(func() {
			mock.ExpectQuery(fmt.Sprintf("SELECT id, name, author, year FROM songs LIMIT %v OFFSET %v", limit, offset)).
				WillReturnRows(sqlmock.
					NewRows([]string{"id", "name", "author", "year"}).
					AddRow(songs[offset].Id, songs[offset].Name, songs[offset].Author, songs[offset].Year).
					AddRow(songs[offset+1].Id, songs[offset+1].Name, songs[offset+1].Author, songs[offset+1].Year),
				)
		})

		It("Should succeed and return correct songs", func() {
			actualSongs, err := repo.ListSongs(ctx, limit, offset)
			Expect(err).To(BeNil())
			Expect(actualSongs).To(BeEquivalentTo(songs[offset : offset+limit]))
		})
	})

	Describe("Remove song", func() {
		Context("the song to be removed is present", func() {
			BeforeEach(func() {
				query := mock.ExpectExec("DELETE FROM songs WHERE id").
					WithArgs(someSong.Id)
				query.WillReturnResult(sqlmock.NewResult(int64(someSong.Id), 1))
			})

			It("Should succeed and return true", func() {
				deleted, err := repo.RemoveSong(ctx, someSong.Id)
				Expect(err).To(BeNil())
				Expect(deleted).To(BeTrue())
			})
		})

		Context("the song to be removed is absent", func() {
			BeforeEach(func() {
				query := mock.ExpectExec("DELETE FROM songs WHERE id").
					WithArgs(someSong.Id)
				query.WillReturnResult(sqlmock.NewResult(0, 0))
			})

			It("Should succeed and return false", func() {
				deleted, err := repo.RemoveSong(ctx, someSong.Id)
				Expect(err).To(BeNil())
				Expect(deleted).To(BeFalse())
			})
		})
	})
})
