package utils

import (
	"errors"
)

var filtered = []int{1, 3, 5, 7, 9}

func FilterHardcoded(s []int) ([]int, error) {
	if s == nil {
		return nil, errors.New("slice can't be nil")
	}

	index := make(map[int]interface{}, len(filtered))
	for _, v := range filtered {
		index[v] = nil
	}

	result := make([]int, 0)
	for _, v := range s {
		if _, present := index[v]; !present {
			result = append(result, v)
		}
	}

	return result, nil
}
