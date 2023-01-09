package service

import (
	"context"
	"weather-forecast/internal/generated"

	"go.uber.org/zap"
)

type CityExistencer struct {
	Storager
	CityMap map[string]int
	CityLoc chan<- string
	Counter int
}

func (s *CityExistencer) AddNewCity(logger *zap.Logger, ctx context.Context, c string) error {
	logger.Debug("")

	if s.CityMap[c] == 0 {
		err := s.Storager.LocationInfo(ctx)
		if err != nil {
			//
			return err
		}
		s.CityMap[c] = s.Counter
		s.Counter += 1
	}

	return nil
}

func (s *CityExistencer) AddWeatherInfo(logger *zap.Logger, ctx context.Context, g generated.CurrentInfo, city string) error {
	logger.Debug("")

	err := s.Storager.WeatherInfo(ctx, g, city)
	if err != nil {
		//
		return err
	}

	return nil
}
