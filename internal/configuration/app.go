package configuration

type App struct {
	Log      Log
	Grpc     Grpc
	Database Database
	Jaeger   Jaeger
	Metrics  Metrics
	Kafka    Kafka
}
