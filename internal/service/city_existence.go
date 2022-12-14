package service

import (
	"context"

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
		err := s.Storager.AddLocationInfo(ctx)
		if err != nil {
			//
			return err
		}
	}
	s.CityMap[c] = s.Counter
	s.Counter += 1

	return nil
}
