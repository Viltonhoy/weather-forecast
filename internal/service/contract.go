package service

import (
	"context"
	"weather-forecast/internal/generated"
)

type Storager interface {
	LocationInfo(ctx context.Context) error
	WeatherInfo(ctx context.Context, c generated.CurrentInfo, name string) error
}
