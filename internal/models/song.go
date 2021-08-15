package models

import "fmt"

type Song struct {
	Id     uint64
	Author string
	Name   string
	Year   int
}

func (s *Song) String() string {
	return fmt.Sprintf("%v: %v - %v (%v)", s.Id, s.Author, s.Name, s.Year)
}

func CreateSong(id uint64, author string, name string, year int) *Song {
	return &Song{id, author, name, year}
}

func CreateSongWithId(id uint64) *Song {
	s := new(Song)
	s.Id = id
	return s
}
