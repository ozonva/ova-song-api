package flusher

import (
	. "context"

	"github.com/ozonva/ova-song-api/internal/models"
	"github.com/ozonva/ova-song-api/internal/repo"
	"github.com/ozonva/ova-song-api/internal/utils"
)

type Flusher interface {
	Flush(ctx Context, songs []models.Song) []models.Song // no error; see slack
}

func NewFlusher(chunkSize int, songsRepo repo.Repo) Flusher {
	return &flusher{
		chunkSize: chunkSize,
		songsRepo: songsRepo,
	}
}

type flusher struct {
	chunkSize int
	songsRepo repo.Repo
}

func (f *flusher) Flush(ctx Context, songs []models.Song) []models.Song {
	if f.chunkSize <= 0 {
		return songs
	}

	chunks, err := utils.DivideSliceOfSongs(songs, f.chunkSize)
	if err != nil {
		return songs
	}

	var failedSongs []models.Song
	for _, chunk := range chunks {
		_, err := f.songsRepo.AddSongs(ctx, chunk)
		if err != nil {
			failedSongs = append(failedSongs, chunk...)
		}
	}

	return failedSongs
}
