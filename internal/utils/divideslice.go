package utils

import (
	"errors"
)

func DivideSlice(s []int, batchSize int) ([][]int, error) {
	if batchSize == 0 {
		panic("batch size must be non-zero")
	}

	if s == nil {
		return nil, errors.New("slice can't be nil")
	}

	count := len(s) / batchSize
	result := make([][]int, count)

	pos := 0
	for i := 0; i < count; i++ {
		result[i] = s[pos : pos+batchSize]
		pos += batchSize
	}

	if pos != len(s) {
		result = append(result, s[pos:])
	}

	return result, nil
}
