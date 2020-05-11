package postgres

import (
	"context"

	"github.com/goldenCRM.git/lib/models"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
)

type Postgres struct {
	pool *pgxpool.Pool
}

func New(ctx context.Context, conf *Config) (*Postgres, error) {

	pgxConf, err := pgxpool.ParseConfig(conf.PoolConnURL())
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse connection URL")
	}

	p, err := pgxpool.ConnectConfig(ctx, pgxConf)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create connections pool")
	}

	return &Postgres{
		pool: p,
	}, nil
}

func (p *Postgres) Add(flat models.Flat) error {
	return nil
}
func (p *Postgres) List() ([]models.Flat, error) {
	return []models.Flat{}, nil
}
