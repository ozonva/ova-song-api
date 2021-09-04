package api

import (
	"context"
	"errors"

	. "github.com/ozonva/ova-song-api/internal/models"
	rp "github.com/ozonva/ova-song-api/internal/repo"
	"github.com/ozonva/ova-song-api/internal/utils"
	desc "github.com/ozonva/ova-song-api/pkg/ova-song-api"
	log "github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type api struct {
	desc.UnimplementedOvaSongApiServer

	repo      rp.Repo
	batchSize int
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

func (s *api) CreateSongMultiV1(
	ctx context.Context,
	req *desc.CreateSongMultiV1Request,
) (
	*desc.CreateSongMultiV1Response,
	error,
) {
	log.Debug().
		Str("Method name", "CreateSongMultiV1").
		Int("Total count", len(req.Songs)).
		Msg("Method called")

	if len(req.Songs) == 0 {
		log.Error().Str("Method name", "CreateSongMulti").Msg("songs should be provided")
		return nil, status.Error(codes.InvalidArgument, "songs should be provided")
	}

	var songs []Song
	for i := range req.Songs {
		songs = append(songs, CreateSong(0, req.Songs[i].Author, req.Songs[i].Name, int(req.Songs[i].Year)))
	}

	chunks, err := utils.DivideSliceOfSongs(songs, s.batchSize)
	if err != nil {
		return nil, err
	}

	var failedCount int
	var lastId int64
	for _, chunk := range chunks {
		id, err := s.repo.AddSongs(chunk)
		if err != nil {
			failedCount += len(chunk)
		} else {
			lastId = id
		}
	}

	if failedCount > 0 {
		log.Warn().
			Str("Method name", "CreateSongMultiV1").
			Int("count", failedCount).
			Msg("Failed so save some songs")
	}

	if err != nil {
		log.Error().
			Str("Method name", "CreateSongMultiV1").
			Err(err).
			Msg("Error returned from repo")
		return nil, err
	}
	return &desc.CreateSongMultiV1Response{LastInsertedId: uint64(lastId)}, nil
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

func (s *api) UpdateSongV1(
	ctx context.Context,
	req *desc.UpdateSongV1Request,
) (
	*desc.UpdateSongV1Response,
	error,
) {
	if req.Song == nil {
		log.Error().Str("Method name", "UpdateSongV1").Msg("song should be provided")
		return nil, status.Error(codes.InvalidArgument, "song should be provided")
	}

	log.Debug().
		Str("Method name", "UpdateSongV1").
		Uint64("id", req.Song.Id).
		Str("Name", req.Song.Name).
		Str("Author", req.Song.Author).
		Int32("Year", req.Song.Year).
		Msg("Method called")

	updated, err := s.repo.UpdateSong(CreateSong(req.Song.Id, req.Song.Author, req.Song.Name, int(req.Song.Year)))

	if err != nil {
		log.Error().
			Str("Method name", "UpdateSongV1").
			Err(err).
			Msg("Error returned from repo")
		return nil, err
	}

	return &desc.UpdateSongV1Response{Updated: updated}, nil
}

func (s *api) ListSongsV1(
	ctx context.Context,
	req *desc.ListSongsV1Request,
) (
	*desc.ListSongsV1Response,
	error,
) {
	log.Debug().
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
	log.Debug().
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

func NewSongApi(repo rp.Repo, batchSize int) (desc.OvaSongApiServer, error) {
	if batchSize <= 0 {
		return nil, errors.New("chunkSize must be positive")
	}
	return &api{repo: repo, batchSize: batchSize}, nil
}
