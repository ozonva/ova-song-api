package repo

import "github.com/ozonva/ova-song-api/internal/models"

type Repo interface {
	AddSongs(songs []models.Song) error
	ListSongs(limit, offset uint64) ([]models.Song, error)
	DescribeSong(songId uint64) (*models.Song, error)
}
