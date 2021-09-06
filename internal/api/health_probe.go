package api

import (
	"context"

	desc "github.com/ozonva/ova-song-api/pkg/health-probe"
)

type healthApi struct {
	desc.UnsafeHealthServer
}

func (*healthApi) CheckHealthV1(
	_ context.Context,
	request *desc.CheckHealthV1Request,
) (
	*desc.CheckHealthV1Response,
	error,
) {
	return &desc.CheckHealthV1Response{EchoText: request.Text}, nil
}

func NewHealthApi() desc.HealthServer {
	return &healthApi{}
}
