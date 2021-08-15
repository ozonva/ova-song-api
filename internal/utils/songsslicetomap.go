package utils

import (
	"errors"
	"fmt"

	"github.com/ozonva/ova-song-api/internal/models"
)

func SongSliceToMap(songs []models.Song) (map[uint64]models.Song, error) {
	m := make(map[uint64]models.Song, len(songs))

	for i := range songs {
		id := songs[i].Id

		if _, duplicate := m[id]; duplicate {
			return nil, errors.New(fmt.Sprintf("when converting to a map, a duplicate was found, id: %v", id))
		}
		m[id] = songs[i]
	}

	return m, nil
}
