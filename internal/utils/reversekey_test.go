package utils

import (
	"reflect"
	"testing"
)

func TestReverseKeyNil(t *testing.T) {
	_, err := ReverseKey(nil)
	if err == nil {
		t.Errorf("ReverseKey failed. In: nil, expected error != null, actual: nil")
	}
}

func TestReverseKeySuccess(t *testing.T) {
	in := map[int]string{1: "a1", 2: "a2", 3: "a3", 4: "a4"}
	expected := map[string]int{"a1": 1, "a2": 2, "a3": 3, "a4": 4}

	actual, actualErr := ReverseKey(in)
	if actualErr != nil {
		t.Fatalf("ReverseKey failed. In: %v, expected error == null, actual: %v", in, actualErr)
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("ReverseKey failed. In: %v, expected: %v, actual: %v", in, expected, actual)
	}
}

func TestReverseKeyNonUniqueValues(t *testing.T) {
	in := map[int]string{1: "a1", 2: "a1", 3: "a3", 4: "a5"}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("ReverseKey failed. Expected: panic, actual: no panic")
		}
	}()

	ignored, _ := ReverseKey(in)
	_ = ignored
}
