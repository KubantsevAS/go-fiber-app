package database

import (
	"context"
	"demo/go-fiber/config"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
)

func CreateDbPool(config *config.DataBaseConfig, logger *zerolog.Logger) *pgxpool.Pool {
	dbpool, err := pgxpool.New(context.Background(), config.Url)
	if err != nil {
		logger.Error().Msg("Filed connection to DB")
		panic(err)
	}
	logger.Info().Msg("Connected to DB")
	return dbpool
}
