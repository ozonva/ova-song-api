package api

import (
	"context"

	desc "github.com/ozonva/ova-song-api/pkg/ova-song-api"
	log "github.com/rs/zerolog/log"
)

type api struct {
	desc.UnimplementedOvaSongApiServer
}

func (s *api) CreateSongV1(
	ctx context.Context,
	req *desc.CreateSongV1Request,
) (
	*desc.CreateSongV1Response,
	error,
) {
	log.Info().
		Str("Method name", "RemoveSongV1").
		Str("name", req.Name).
		Str("Author", req.Author).
		Int32("Year", req.Year).
		Msg("Method called")

	return &desc.CreateSongV1Response{SongId: 10}, nil
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

	return &desc.DescribeSongV1Response{Song: &desc.Song{
		Id:     req.SongId,
		Name:   "name",
		Author: "author",
		Year:   2007,
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

	return &desc.ListSongsV1Response{Songs: []*desc.Song{
		{
			Id:     req.Offset,
			Name:   "name",
			Author: "author",
			Year:   2007},
	}}, nil
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

	log.Printf("RemoveSongV1 called with id: %v", req.SongId)
	return &desc.RemoveSongV1Response{Removed: req.SongId%2 == 0}, nil
}

func NewApi() desc.OvaSongApiServer {
	return &api{}
}
