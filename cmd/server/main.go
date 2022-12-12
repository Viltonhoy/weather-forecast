package main

import (
	"context"
	"log"
	"weather-forecast/internal/server"
	"weather-forecast/internal/storage/postgresql"
	weatherapi "weather-forecast/internal/weather_api"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

var CityLoc = make(chan string)

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("zap.NewDevelopment: %v", err)
	}
	defer logger.Sync()

	if err := godotenv.Load("../../.env"); err != nil {
		logger.Debug("No .env file found", zap.Error(err))
	}

	ctx := context.Background()

	storage, err := postgresql.NewStorage(ctx, logger)
	if err != nil {
		logger.Fatal("failed to create storage instance", zap.Error(err))
	}

	w := weatherapi.New()

	srv, err := server.New(
		logger,
		storage,
		storage.Close,
		w,
	)

	if err != nil {
		logger.Fatal("failed to create http server instance", zap.Error(err))
	}

	err = srv.Start()
	if err != nil {
		logger.Fatal("failed to start or shutdown server", zap.Error(err))
	}

	err = w.CallAt(logger, CityLoc, w.WeatherRates)
	if err != nil {
		logger.Fatal("failed to apiweather client", zap.Error(err))
	}
}
