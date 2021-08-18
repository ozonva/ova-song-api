package utils

import (
	"reflect"
	"testing"

	"github.com/ozonva/ova-song-api/internal/models"
)

func TestSongSliceToMapNil(t *testing.T) {
	m, err := SongSliceToMap(nil)

	if err != nil {
		t.Fatalf("SongSliceToMap failed. In: nil, expected: empty map, actual: error: %v", err)
	}

	if len(m) != 0 {
		t.Fatalf("SongSliceToMap failed. In: nil, expected: empty map, actual: %v", m)
	}
}

func TestSongSliceToMapSuccess(t *testing.T) {
	in := []models.Song{
		*createSongWithId(6),
		*createSongWithId(8),
		*createSongWithId(4),
		*createSongWithId(2),
	}
	expected := map[uint64]models.Song{
		6: *createSongWithId(6),
		8: *createSongWithId(8),
		4: *createSongWithId(4),
		2: *createSongWithId(2),
	}

	actual, actualErr := SongSliceToMap(in)
	if actualErr != nil {
		t.Fatalf("SongSliceToMap failed. In: %v, expected error == null, actual: %v", in, actualErr)
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("SongSliceToMap failed. In: %v, expected: %v, actual: %v", in, expected, actual)
	}
}

func TestSongSliceToMapNonUniqueId(t *testing.T) {
	in := []models.Song{
		*createSongWithId(6),
		*createSongWithId(8),
		*createSongWithId(8),
		*createSongWithId(2),
	}

	_, err := SongSliceToMap(in)
	if err == nil {
		t.Fatalf("SongSliceToMap failed. In: %v, expected error != null, actual: nil", in)
	}
}

func createSongWithId(id uint64) *models.Song {
	s := new(models.Song)
	s.Id = id
	return s
}
