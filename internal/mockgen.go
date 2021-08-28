package internal

//go:generate mockgen -destination=./mocks/repo_mock.go -package=mocks github.com/ozonva/ova-song-api/internal/repo Repo

//go:generate mockgen -destination=./mocks/flusher_mock.go -package=mocks github.com/ozonva/ova-song-api/internal/flusher Flusher

//go:generate mockgen -destination=./mocks/ticker_mock.go -package=mocks github.com/ozonva/ova-song-api/internal/time Ticker
