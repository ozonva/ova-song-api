package utils

import (
	"errors"
)

func DivideSlice(s []int, batchSize int) ([][]int, error) {
	if batchSize == 0 {
		return nil, errors.New("batch size must be non-zero")
	}

	if s == nil {
		return nil, errors.New("slice must be not nil")
	}

	result := make([][]int, divCeil(len(s), batchSize))

	fullBatches := len(s) / batchSize
	pos := 0
	for i := 0; i < fullBatches; i++ {
		result[i] = s[pos : pos+batchSize]
		pos += batchSize
	}

	if pos != len(s) {
		result[len(result)-1] = s[pos:]
	}

	return result, nil
}

func divCeil(divisor, dividend int) int {
	// Notice: you may want to use another form to prevent integer overflow:
	// divCeil = 1 + (divisor - 1) / dividend
	return (divisor + dividend - 1) / dividend
}
