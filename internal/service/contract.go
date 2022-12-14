package service

import (
	"context"
	"weather-forecast/internal/generated"

	"go.uber.org/zap"
)

type Storager interface {
	AddLocationInfo(ctx context.Context) error
}

type ApiClient interface {
	CallAt(logger *zap.Logger, ctx context.Context, loc <-chan string, f func(*zap.Logger, string) (generated.WeatherResult, error)) error
}
