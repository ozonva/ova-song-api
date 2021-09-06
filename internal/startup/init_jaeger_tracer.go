package startup

import (
	"io"

	"github.com/ozonva/ova-song-api/internal/log"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"github.com/uber/jaeger-lib/metrics"
)

func InitJaegerTracer(serviceName, host, port string) (io.Closer, error) {
	cfg := config.Configuration{
		Sampler: &config.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: host + port,
		}}

	jLogger := zerolog_adapter.JaegerAdapter
	jMetricsFactory := metrics.NullFactory

	closer, err := cfg.InitGlobalTracer(
		serviceName,
		config.Logger(jLogger),
		config.Metrics(jMetricsFactory),
	)
	return closer, err
}
