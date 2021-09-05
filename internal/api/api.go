package api

import (
	"context"
	"errors"

	"github.com/opentracing/opentracing-go"
	olog "github.com/opentracing/opentracing-go/log"
	br "github.com/ozonva/ova-song-api/internal/broker"
	"github.com/ozonva/ova-song-api/internal/metrics"
	. "github.com/ozonva/ova-song-api/internal/models"
	rp "github.com/ozonva/ova-song-api/internal/repo"
	"github.com/ozonva/ova-song-api/internal/utils"
	desc "github.com/ozonva/ova-song-api/pkg/ova-song-api"
	log "github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const songIdTag = "song_id"

type api struct {
	desc.UnimplementedOvaSongApiServer

	repo      rp.Repo
	batchSize int
	broker    br.Broker
}

func (s *api) CreateSongV1(
	ctx context.Context,
	req *desc.CreateSongV1Request,
) (
	*desc.CreateSongV1Response,
	error,
) {
	const methodName = "CreateSongV1"
	log.Debug().
		Str("Method name", methodName).
		Str("name", req.Name).
		Str("Author", req.Author).
		Int32("Year", req.Year).
		Msg("Method called")

	tracer := opentracing.GlobalTracer()
	span := tracer.StartSpan(methodName)
	defer span.Finish()

	newId, err := s.repo.AddSong(ctx, CreateSong(0, req.Author, req.Name, int(req.Year)))
	if err != nil {
		log.Error().Str("Method name", methodName).Err(err).Msg("Error returned from repo")
		return nil, err
	}
	span.SetTag(songIdTag, newId)

	metrics.Counters.AddSucceeds.Inc()

	err = s.broker.SendEvent(br.NewCreateEvent(newId))
	if err != nil {
		span.LogFields(olog.Error(err))
		log.Error().Str("Method name", methodName).Err(err).Msg("Failed to send event")
		// graceful; no `return` here
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
	const methodName = "CreateSongMultiV1"

	log.Debug().
		Str("Method name", methodName).
		Int("Total count", len(req.Songs)).
		Msg("Method called")

	tracer := opentracing.GlobalTracer()
	span := tracer.StartSpan(methodName)
	defer span.Finish()

	if len(req.Songs) == 0 {
		log.Error().Str("Method name", "CreateSongMulti").Msg("songs should be provided")
		return nil, status.Error(codes.InvalidArgument, "songs should be provided")
	}

	var songs []Song
	for i := range req.Songs {
		songs = append(songs, CreateSong(0, req.Songs[i].Author, req.Songs[i].Name, int(req.Songs[i].Year)))
	}

	span.LogFields(olog.Int("Songs count", len(songs)))

	batch, err := utils.DivideSliceOfSongs(songs, s.batchSize)
	if err != nil {
		return nil, err
	}

	span.LogFields(olog.Int("Chunks count", len(batch)))
	span.LogFields(olog.Int("Chunk size", s.batchSize))

	var failedCount int
	var lastId int64
	for i := range batch {
		bSpan := opentracing.StartSpan(methodName+" Batch", opentracing.ChildOf(span.Context()))
		bSpan.LogFields(olog.Int("Batch number", i))
		bSpan.LogFields(olog.Int("Songs count", len(batch[i])))

		id, err := s.repo.AddSongs(ctx, batch[i])
		if err != nil {
			failedCount += len(batch[i])
			bSpan.LogFields(olog.Error(err))
		} else {
			lastId = id
			err = s.broker.SendEvent(br.NewCreateMultiEvent(id))
			if err != nil {
				bSpan.LogFields(olog.Error(err))
				log.Error().Str("Method name", methodName).Err(err).Int("Batch no", i).Msg("Failed to send event")
				// graceful; no `return` here
			}
		}

		bSpan.Finish()
	}

	if failedCount > 0 {
		span.LogFields(olog.Int("Songs failed", failedCount))
		log.Warn().
			Str("Method name", "CreateSongMultiV1").
			Int("count", failedCount).
			Msg("Failed so save some songs")
	} else {
		// succeed only if no one failed
		metrics.Counters.AddMultiSucceeds.Inc()
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
	const methodName = "DescribeSongV1"

	log.Debug().
		Str("Method name", methodName).
		Uint64("SongId", req.SongId).
		Msg("Method called")

	tracer := opentracing.GlobalTracer()
	span := tracer.StartSpan(methodName)
	defer span.Finish()
	span.SetTag(songIdTag, req.SongId)
	span.SetTag("song_id2", nil)

	song, err := s.repo.DescribeSong(ctx, req.SongId)
	if err != nil {
		span.LogFields(olog.Error(err))
		log.Error().
			Str("Method name", methodName).
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
	const methodName = "UpdateSongV1"

	tracer := opentracing.GlobalTracer()
	span := tracer.StartSpan(methodName)
	defer span.Finish()

	if req.Song == nil {
		log.Error().Str("Method name", methodName).Msg("song should be provided")
		return nil, status.Error(codes.InvalidArgument, "song should be provided")
	}

	log.Debug().
		Str("Method name", methodName).
		Uint64("id", req.Song.Id).
		Str("Name", req.Song.Name).
		Str("Author", req.Song.Author).
		Int32("Year", req.Song.Year).
		Msg("Method called")

	span.SetTag(songIdTag, req.Song.Id)

	updated, err := s.repo.UpdateSong(ctx, CreateSong(req.Song.Id, req.Song.Author, req.Song.Name, int(req.Song.Year)))
	if err != nil {
		span.LogFields(olog.Error(err))
		log.Error().
			Str("Method name", "UpdateSongV1").
			Err(err).
			Msg("Error returned from repo")
		return nil, err
	}

	if updated {
		metrics.Counters.UpdateSucceeds.Inc()

		err = s.broker.SendEvent(br.NewUpdateEvent(int64(req.Song.Id)))
		if err != nil {
			span.LogFields(olog.Error(err))
			log.Error().Str("Method name", methodName).Err(err).Msg("Failed to send event")
			// graceful; no `return` here
		}
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
	const methodName = "ListSongsV1"

	tracer := opentracing.GlobalTracer()
	span := tracer.StartSpan(methodName)
	defer span.Finish()

	log.Debug().
		Str("Method name", methodName).
		Uint64("Offset", req.Offset).
		Uint64("Limit", req.Limit).
		Msg("Method called")

	span.LogFields(olog.Uint64("Offset", req.Offset))
	span.LogFields(olog.Uint64("Limit", req.Limit))

	songs, err := s.repo.ListSongs(ctx, req.Limit, req.Offset)

	if err != nil {
		span.LogFields(olog.Error(err))
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
	const methodName = "RemoveSongV1"

	tracer := opentracing.GlobalTracer()
	span := tracer.StartSpan(methodName)
	defer span.Finish()

	log.Debug().
		Str("Method name", methodName).
		Uint64("SongId", req.SongId).
		Msg("Method called")
	span.SetTag(songIdTag, req.SongId)

	removed, err := s.repo.RemoveSong(ctx, req.SongId)
	if err != nil {
		span.LogFields(olog.Error(err))
		log.Error().
			Str("Method name", "RemoveSongV1").
			Err(err).
			Msg("Error returned from repo")
		return nil, err
	}

	if removed {
		metrics.Counters.DeleteSucceeds.Inc()

		err = s.broker.SendEvent(br.NewRemoveEvent(int64(req.SongId)))
		if err != nil {
			span.LogFields(olog.Error(err))
			log.Error().Str("Method name", methodName).Err(err).Msg("Failed to send event")
			// graceful; no `return` here
		}
	}

	return &desc.RemoveSongV1Response{Removed: removed}, nil
}

func NewSongApi(repo rp.Repo, batchSize int, broker br.Broker) (desc.OvaSongApiServer, error) {
	if batchSize <= 0 {
		return nil, errors.New("chunkSize must be positive")
	}
	return &api{repo: repo, batchSize: batchSize, broker: broker}, nil
}
