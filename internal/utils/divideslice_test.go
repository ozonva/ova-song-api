package utils

import (
	"reflect"
	"testing"
)

type divideSliceTest struct {
	in        []int
	batchSize int
	expected  [][]int
	error     bool
}

var divideSliceTests = []divideSliceTest{
	{nil, 10, nil, true},
	{[]int{}, 2, [][]int{}, false},
	{[]int{9, 1, 2, 3, 4, 5, 6, 7, 8}, 2, [][]int{{9, 1}, {2, 3}, {4, 5}, {6, 7}, {8}}, false},
	{[]int{9, 1, 2, 3, 4, 5, 6, 7, 8}, 1, [][]int{{9}, {1}, {2}, {3}, {4}, {5}, {6}, {7}, {8}}, false},
	{[]int{9, 1, 2, 3, 4, 5, 6, 7, 8}, 25, [][]int{{9, 1, 2, 3, 4, 5, 6, 7, 8}}, false},
	{[]int{9, 1, 2, 3, 4, 5, 6, 7, 8}, 10, [][]int{{9, 1, 2, 3, 4, 5, 6, 7, 8}}, false},
}

func TestDivideSlice(t *testing.T) {
	for _, test := range divideSliceTests {
		actual, err := DivideSlice(test.in, test.batchSize)

		if (test.error && err == nil) || (!test.error && err != nil) {
			t.Errorf("DivideSlice failed. In: (%v, %v). Expected error == nil, actual: %v",
				test.in, test.batchSize, err)
			continue
		}

		if !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("DivideSlice failed. In: (%v, %v). Expected %v, actual: %v",
				test.in, test.batchSize, test.expected, actual)
		}
	}
}

func TestDivideSliceZeroSizedBatch(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("DivideSlice failed. Expected: panic, actual: no panic")
		}
	}()

	ignored, _ := DivideSlice([]int{1, 2, 3}, 0)
	_ = ignored
}
