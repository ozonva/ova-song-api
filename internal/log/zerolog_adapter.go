package zerolog_adapter

import "github.com/rs/zerolog/log"

// JaegerAdapter is implementation of the jaeger.Logger interface that delegates all calls to the zerolog.
var JaegerAdapter = &jaegerAdapter{}

type jaegerAdapter struct{}

func (*jaegerAdapter) Error(msg string) {
	log.Error().Msg(msg)
}

func (*jaegerAdapter) Infof(msg string, args ...interface{}) {
	log.Info().Msgf(msg, args...)
}

func (*jaegerAdapter) Debugf(msg string, args ...interface{}) {
	log.Debug().Msgf(msg, args...)
}
