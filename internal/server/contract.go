package server

import (
	"context"

	"go.uber.org/zap"
)

type CityExistence interface {
	AddNewCity(logger *zap.Logger, ctx context.Context, c string) error
}
