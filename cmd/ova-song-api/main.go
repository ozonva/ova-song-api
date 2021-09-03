package main

import (
	"net"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	api "github.com/ozonva/ova-song-api/internal/api"
	rp "github.com/ozonva/ova-song-api/internal/repo"
	desc "github.com/ozonva/ova-song-api/pkg/ova-song-api"
	log "github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal().Msgf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	desc.RegisterOvaSongApiServer(s, createSongApi())
	if err := s.Serve(lis); err != nil {
		log.Fatal().Msgf("failed to serve: %v", err)
	}
	log.Fatal().Msgf("11failed to serve: %v", err)
}

func createSongApi() desc.OvaSongApiServer {
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
	songApi := api.NewSongApi(repo)
	return songApi
}
