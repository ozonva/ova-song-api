package saver

import (
	"time"

	"github.com/ozonva/ova-song-api/internal/models"
)

type Saver interface {
	Save(song models.Song) // no error returned; see slack
	Close()
}

type saver struct {
	songChan  chan<- models.Song
	closeChan chan<- struct{}
}

func (s *saver) Save(song models.Song) {
	s.songChan <- song
}

func (s *saver) Close() {
	// note that order matters
	close(s.songChan)
	close(s.closeChan)
}

type saverBackend struct {
	period      time.Duration
	flusher     Flusher
	songsBuffer []models.Song

	songChan  <-chan models.Song
	closeChan <-chan struct{}
}

func (s *saverBackend) Serve() {
	t := time.NewTicker(s.period)
	for {
		select {
		case newSong, ok := <-s.songChan:
			if !ok {
				s.performClosing(t)
				return
			}
			s.songsBuffer = append(s.songsBuffer, newSong)

		case <-t.C:
			s.doSave()
			// resetting timer to prevent too frequent calls(e.g. in cases if doSave took > s.period time)
			t.Reset(s.period)
		}
	}
}

func (s *saverBackend) performClosing(t *time.Ticker) {
	t.Stop()

	<-s.closeChan
	s.doSave()
}

type Flusher interface {
	Flush(songs []models.Song) []models.Song
}

func (s *saverBackend) doSave() {
	println("doSave called; len, cap = ", len(s.songsBuffer), cap(s.songsBuffer))
	if len(s.songsBuffer) == 0 {
		return
	}

	failed := s.flusher.Flush(s.songsBuffer)
	s.songsBuffer = s.songsBuffer[:0]
	println("buff cap = ", cap(s.songsBuffer))
	s.songsBuffer = append(s.songsBuffer, failed...)
}

func NewSaver(capacity uint, period time.Duration, flusher Flusher) Saver {
	songsChan := make(chan models.Song, capacity)
	closeChan := make(chan struct{}) // intentionally unbuffered

	saver := saver{
		songChan:  songsChan,
		closeChan: closeChan,
	}

	backend := saverBackend{
		period:      period,
		flusher:     flusher,
		songsBuffer: make([]models.Song, 0, capacity),
		songChan:    songsChan,
		closeChan:   closeChan,
	}

	go backend.Serve()
	return &saver
}
