package configuration

type App struct {
	Grpc     Grpc
	Database Database
	Jaeger   Jaeger
	Metrics  Metrics
	Kafka    Kafka
}
