package utils

import (
	"reflect"
	"strconv"
	"testing"

	"github.com/ozonva/ova-song-api/internal/models"
)

type divideSliceTest struct {
	in        []int
	batchSize int
	expected  [][]int
	error     bool
}

var divideSliceTests = []divideSliceTest{
	{nil, 10, nil, true},
	{[]int{1, 2, 3}, 0, nil, true},
	{[]int{1, 2, 3}, -1, nil, true},
	{[]int{}, -1, nil, true},
	{[]int{}, 2, [][]int{}, false},
	{[]int{9, 1, 2, 3, 4, 5, 6, 7, 8}, 2, [][]int{{9, 1}, {2, 3}, {4, 5}, {6, 7}, {8}}, false},
	{[]int{9, 1, 2, 3, 4, 5, 6, 7, 8}, 1, [][]int{{9}, {1}, {2}, {3}, {4}, {5}, {6}, {7}, {8}}, false},
	{[]int{9, 1, 2, 3, 4, 5, 6, 7, 8}, 25, [][]int{{9, 1, 2, 3, 4, 5, 6, 7, 8}}, false},
	{[]int{9, 1, 2, 3, 4, 5, 6, 7, 8}, 10, [][]int{{9, 1, 2, 3, 4, 5, 6, 7, 8}}, false},
}

func TestDivideSlice(t *testing.T) {
	for i, test := range divideSliceTests {
		actual, err := DivideSlice(test.in, test.batchSize)

		if (test.error && err == nil) || (!test.error && err != nil) {
			t.Errorf("DivideSlice failed. Case: %v. In: (%v, %v). Expected error == nil, actual: %v",
				i, test.in, test.batchSize, err)
			continue
		}

		if !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("DivideSlice failed. Case: %v. In: (%v, %v). Expected %v, actual: %v",
				i, test.in, test.batchSize, test.expected, actual)
		}
	}
}

func TestDivideSliceOfSongs(t *testing.T) {
	for i, test := range divideSliceTests {
		in, expected := prepareTestDataForSongs(test.in, test.expected)

		actual, err := DivideSliceOfSongs(in, test.batchSize)

		if (test.error && err == nil) || (!test.error && err != nil) {
			t.Errorf("DivideSlice failed. Case: %v. In: (%v, %v). Expected error == nil, actual: %v",
				i, in, test.batchSize, err)
			continue
		}

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("DivideSlice failed. Case: %v. In: (%v, %v). Expected %v, actual: %v",
				i, test.in, test.batchSize, expected, actual)
		}
	}
}

func prepareTestDataForSongs(inInts []int, expectedInts [][]int) (in []models.Song, expected [][]models.Song) {
	in = songSliceFromInts(inInts)

	if expectedInts == nil {
		return in, nil
	}

	expected = make([][]models.Song, len(expectedInts))
	for i, ints := range expectedInts {
		expected[i] = songSliceFromInts(ints)
	}
	return in, expected
}

func songSliceFromInts(slice []int) []models.Song {
	if slice == nil {
		return nil
	}

	result := make([]models.Song, len(slice))

	for i, e := range slice {
		result[i] = models.CreateSong(0, strconv.Itoa(e), "", 0)
	}

	return result
}
