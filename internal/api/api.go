package api

import (
	"context"

	. "github.com/ozonva/ova-song-api/internal/models"
	rp "github.com/ozonva/ova-song-api/internal/repo"
	desc "github.com/ozonva/ova-song-api/pkg/ova-song-api"
	log "github.com/rs/zerolog/log"
)

type api struct {
	desc.UnimplementedOvaSongApiServer

	repo rp.Repo
}

func (s *api) CreateSongV1(
	ctx context.Context,
	req *desc.CreateSongV1Request,
) (
	*desc.CreateSongV1Response,
	error,
) {
	log.Debug().
		Str("Method name", "CreateSongV1").
		Str("name", req.Name).
		Str("Author", req.Author).
		Int32("Year", req.Year).
		Msg("Method called")

	newId, err := s.repo.AddSong(CreateSong(0, req.Author, req.Name, int(req.Year)))

	if err != nil {
		log.Error().
			Str("Method name", "CreateSongV1").
			Err(err).
			Msg("Error returned from repo")
		return nil, err
	}
	return &desc.CreateSongV1Response{SongId: uint64(newId)}, nil
}

func (s *api) DescribeSongV1(
	ctx context.Context,
	req *desc.DescribeSongV1Request,
) (
	*desc.DescribeSongV1Response,
	error,
) {
	log.Info().
		Str("Method name", "DescribeSongV1").
		Uint64("SongId", req.SongId).
		Msg("Method called")

	song, err := s.repo.DescribeSong(req.SongId)

	if err != nil {
		log.Error().
			Str("Method name", "DescribeSongV1").
			Err(err).
			Msg("Error returned from repo")
		return nil, err
	}

	return &desc.DescribeSongV1Response{Song: &desc.Song{
		Id:     song.Id,
		Name:   song.Name,
		Author: song.Author,
		Year:   int32(song.Year),
	}}, nil
}

func (s *api) ListSongsV1(
	ctx context.Context,
	req *desc.ListSongsV1Request,
) (
	*desc.ListSongsV1Response,
	error,
) {
	log.Info().
		Str("Method name", "ListSongsV1").
		Uint64("Offset", req.Offset).
		Uint64("Limit", req.Limit).
		Msg("Method called")

	songs, err := s.repo.ListSongs(req.Limit, req.Offset)

	if err != nil {
		log.Error().
			Str("Method name", "ListSongsV1").
			Err(err).
			Msg("Error returned from repo")
		return nil, err
	}

	res := &desc.ListSongsV1Response{Songs: make([]*desc.Song, 0, len(songs))}
	for i := range songs {
		res.Songs = append(res.Songs,
			&desc.Song{
				Id:     songs[i].Id,
				Name:   songs[i].Name,
				Author: songs[i].Author,
				Year:   int32(songs[i].Year),
			})
	}
	return res, nil
}

func (s *api) RemoveSongV1(
	ctx context.Context,
	req *desc.RemoveSongV1Request,
) (
	*desc.RemoveSongV1Response,
	error,
) {
	log.Info().
		Str("Method name", "RemoveSongV1").
		Uint64("SongId", req.SongId).
		Msg("Method called")

	removed, err := s.repo.RemoveSong(req.SongId)

	if err != nil {
		log.Error().
			Str("Method name", "RemoveSongV1").
			Err(err).
			Msg("Error returned from repo")
		return nil, err
	}
	return &desc.RemoveSongV1Response{Removed: removed}, nil
}

func NewSongApi(repo rp.Repo) desc.OvaSongApiServer {
	return &api{repo: repo}
}
