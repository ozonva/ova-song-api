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
		*models.CreateSongWithId(6),
		*models.CreateSongWithId(8),
		*models.CreateSongWithId(4),
		*models.CreateSongWithId(2),
	}
	expected := map[uint64]models.Song{
		6: *models.CreateSongWithId(6),
		8: *models.CreateSongWithId(8),
		4: *models.CreateSongWithId(4),
		2: *models.CreateSongWithId(2),
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
		*models.CreateSongWithId(6),
		*models.CreateSongWithId(8),
		*models.CreateSongWithId(8),
		*models.CreateSongWithId(2),
	}

	_, err := SongSliceToMap(in)
	if err == nil {
		t.Fatalf("SongSliceToMap failed. In: %v, expected error != null, actual: nil", in)
	}
}
