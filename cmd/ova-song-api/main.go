package main

import (
	"net"
	"net/http"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	api "github.com/ozonva/ova-song-api/internal/api"
	br "github.com/ozonva/ova-song-api/internal/broker"
	"github.com/ozonva/ova-song-api/internal/configuration"
	rp "github.com/ozonva/ova-song-api/internal/repo"
	"github.com/ozonva/ova-song-api/internal/startup"
	descHealth "github.com/ozonva/ova-song-api/pkg/health-probe"
	desc "github.com/ozonva/ova-song-api/pkg/ova-song-api"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func main() {
	conf := ReadConfig()

	tracerCloser, err := startup.InitJaegerTracer(conf.Jaeger.ServiceName, conf.Jaeger.Host, conf.Jaeger.Port)
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
		http.Handle(conf.Metrics.HandleName, promhttp.Handler())
		if err := http.ListenAndServe(conf.Metrics.ListenPort, nil); err != nil {
			log.Error().Err(err).Msg("Failed to start listen to metric requests")
		}
	}()

	broker, err := br.NewKafkaBroker(conf.Kafka.Brokers, conf.Kafka.TopicName)
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

	lis, err := net.Listen("tcp", conf.Grpc.Port)
	if err != nil {
		log.Fatal().Err(err).Msgf("failed to listen grpc port")
	}

	s := grpc.NewServer()
	descHealth.RegisterHealthServer(s, api.NewHealthApi())
	desc.RegisterOvaSongApiServer(s, createSongApi(conf.Database, broker))
	if err := s.Serve(lis); err != nil {
		log.Fatal().Err(err).Msg("api failed to serve")
	}
	log.Fatal().Msgf("api failed to serve")
}

func ReadConfig() configuration.App {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal().Err(err).Msg("Error reading config file")
	}

	var conf configuration.App
	err := viper.Unmarshal(&conf)
	if err != nil {
		log.Fatal().Err(err).Msg("Unable to decode configuration into struct")
	}

	return conf
}

func createSongApi(config configuration.Database, broker br.Broker) desc.OvaSongApiServer {
	dsn := "postgres://" + config.Username + ":" + config.Password + "@" +
		config.Host + ":" + config.Port + "/" + config.DbName + "?sslmode=disable"

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
