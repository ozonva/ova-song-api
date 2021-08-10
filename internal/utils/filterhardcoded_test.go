package utils

import (
	"reflect"
	"testing"
)

func TestFilterHardcodedNil(t *testing.T) {
	_, err := FilterHardcoded(nil)
	if err == nil {
		t.Errorf("FilterHardcoded failed. In: nil, expected error != null, actual: nil")
	}
}

func TestFilterHardcodedSuccess(t *testing.T) {
	in := []int{6, 7, 8, 9, 0, 1, 2, 3, 4, 5}
	expected := []int{6, 8, 0, 2, 4}

	actual, actualErr := FilterHardcoded(in)
	if actualErr != nil {
		t.Fatalf("FilterHardcoded failed. In: %v, expected error == null, actual: %v", in, actualErr)
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("FilterHardcoded failed. In: %v, expected: %v, actual: %v", in, expected, actual)
	}
}
