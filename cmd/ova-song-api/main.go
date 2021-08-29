package main

import (
	"net"

	api "github.com/ozonva/ova-song-api/internal/api"
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
	desc.RegisterOvaSongApiServer(s, api.NewApi())
	if err := s.Serve(lis); err != nil {
		log.Fatal().Msgf("failed to serve: %v", err)
	}
	log.Fatal().Msgf("11failed to serve: %v", err)
}
