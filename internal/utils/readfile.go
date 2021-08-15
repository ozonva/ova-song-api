package utils

import (
	"log"
	"os"
)

const times int = 10

func ReadFile(path string) {
	for i := 0; i < times; i++ {
		err := doReadFile(path)
		if err != nil {
			return
		}
	}
}

func doReadFile(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}

	defer func() {
		if err := f.Close(); err != nil {
			log.Fatalf("failed to close file: \"%v\", with error: %v", path, err)
		}
	}()

	// no need to read file
	return nil
}
