package utils

import (
	"errors"
)

func ReverseKey(m map[int]string) (map[string]int, error) {
	if m == nil {
		return nil, errors.New("map can't be nil")
	}

	result := make(map[string]int, len(m))

	for k, v := range m {
		if _, ok := result[v]; ok {
			return nil, errors.New("non unique value find")
		}
		result[v] = k
	}

	return result, nil
}
