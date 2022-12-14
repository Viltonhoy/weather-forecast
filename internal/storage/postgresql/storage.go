package postgresql

import (
	"context"
	"errors"
	"weather-forecast/internal/generated"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/log/zapadapter"
	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
)

type Storage struct {
	Logger *zap.Logger
	DB     *pgxpool.Pool
}

var LocationInfoError = errors.New("City already exists")

func NewStorage(ctx context.Context, logger *zap.Logger) (*Storage, error) {
	if logger == nil {
		return nil, errors.New("no logger provided")
	}

	// taking connect info from environment variables
	config, _ := pgxpool.ParseConfig("")

	config.ConnConfig.Logger = zapadapter.NewLogger(logger)
	config.ConnConfig.LogLevel = pgx.LogLevelError

	// create a pool connection
	pool, err := pgxpool.ConnectConfig(ctx, config)
	if err != nil {
		logger.Error("error database connection", zap.Error(err))
		return nil, err
	}

	err = pool.Ping(ctx)
	if err != nil {
		logger.Error("connection was not established", zap.Error(err))
		return &Storage{}, err
	}

	return &Storage{
		Logger: logger,
		DB:     pool,
	}, err
}

func (s *Storage) Close() {
	s.Logger.Info("closing Storage connection")
	s.DB.Close()
}

func (s *Storage) AddLocationInfo(ctx context.Context, l generated.LocationInfo, c generated.CurrentInfo) error {
	logger := s.Logger.With()
	logger.Debug("")

	firstInsertExec := `insert into location_info (name, region, country, lat, lon, tz_id, localtime_epoch, localtime) values ($1, $2, $3, $4, $5, $6, $7, $8);`

	_, err := s.DB.Exec(
		ctx,
		firstInsertExec,
		l.Name,
		l.Region,
		l.Country,
		l.Lat,
		l.Lon,
		l.Tz_id,
		l.Localtime_epoch,
		l.Localtime,
	)

	if err != nil {
		logger.Error("failed to insert record", zap.Error(LocationInfoError))
		return LocationInfoError
	}

	return err
}
