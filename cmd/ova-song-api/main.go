package main

import (
	"net"
	"net/http"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	api "github.com/ozonva/ova-song-api/internal/api"
	br "github.com/ozonva/ova-song-api/internal/broker"
	rp "github.com/ozonva/ova-song-api/internal/repo"
	"github.com/ozonva/ova-song-api/internal/startup"
	desc "github.com/ozonva/ova-song-api/pkg/ova-song-api"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

func main() {
	tracerCloser, err := startup.InitJaegerTracer("ova-song-api", "127.0.0.1", ":6831")
	if err != nil {
		log.Error().Err(err).Msg("Could not initialize jaeger tracer")
	} else {
		defer func() {
			if err := tracerCloser.Close(); err != nil {
				log.Error().Err(err).Msg("Failed to close tracer")
			}
		}()
		log.Info().Msg("Tracer started")
	}

	go func() {
		http.Handle("/metrics", promhttp.Handler())
		if err := http.ListenAndServe(":8056", nil); err != nil {
			log.Fatal().Err(err).Msgf("Failed to start listen to metric requests, error %s", err)
		}
	}()

	broker, err := br.NewKafkaBroker([]string{"kafka:9092"}, "songs")
	if err != nil {
		log.Error().Err(err).Msg("Failed to start kafka producer")
		broker = br.NewNullBroker() // graceful
	} else {
		defer func() {
			if err := broker.Close(); err != nil {
				log.Error().Err(err).Msg("Failed to close broker")
			}
		}()
	}

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal().Err(err).Msgf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	desc.RegisterOvaSongApiServer(s, createSongApi(broker))
	if err := s.Serve(lis); err != nil {
		log.Fatal().Msgf("failed to serve: %v", err)
	}
	log.Fatal().Msgf("11failed to serve: %v", err)
}

func createSongApi(broker br.Broker) desc.OvaSongApiServer {
	dsn := "postgres://ova_song_db_user:ova_song_db_user@localhost:5432/ova_song_db?sslmode=disable"

	db, err := sqlx.Open("pgx", dsn)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to database")
	}

	if err := db.Ping(); err != nil {
		log.Fatal().Err(err).Msg("Failed to ping database after connect")
	}

	log.Info().Msg("Connected to db")

	repo := rp.NewRepo(db)
	const batchSize = 2
	songApi, err := api.NewSongApi(repo, batchSize, broker)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create song api")
	}
	return songApi
}
