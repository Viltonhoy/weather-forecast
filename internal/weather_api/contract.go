package weatherapi

import (
	"context"
	"weather-forecast/internal/generated"

	"go.uber.org/zap"
)

type Servicer interface {
	AddWeatherInfo(logger *zap.Logger, ctx context.Context, g generated.CurrentInfo, city string) error
	AddNewCity(logger *zap.Logger, ctx context.Context, c string) error
}
