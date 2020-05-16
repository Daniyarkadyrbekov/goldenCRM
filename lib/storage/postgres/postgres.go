package postgres

import (
	"context"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	bindata "github.com/golang-migrate/migrate/v4/source/go_bindata"
	"github.com/goldenCRM.git/lib/models"
	"github.com/goldenCRM.git/lib/storage/postgres/migrations"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
)

const (
	_QueryGetFlats = `
		SELECT
		id, street
		FROM flats
		LIMIT 5000`
	_QuerySaveFlat = `
		INSERT INTO flats (id, street)
		VALUES ($1, $2)`
)

type Postgres struct {
	pool *pgxpool.Pool
}

func New(ctx context.Context, connUrl string) (*Postgres, error) {

	if err := runMigrations(connUrl); err != nil {
		return nil, errors.Wrap(err, "failed to run migrations")
	}

	pgxConf, err := pgxpool.ParseConfig(connUrl)
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
	_, err := p.pool.Exec(context.Background(), _QuerySaveFlat, flat.ID, flat.Street)
	return err
}

func (p *Postgres) List() ([]models.Flat, error) {

	flatsRows, err := p.pool.Query(context.Background(), _QueryGetFlats)
	if err != nil {
		return nil, err
	}

	flats, err := readFlats(flatsRows)
	if err != nil {
		return nil, err
	}

	return flats, nil
}

func readFlats(rows pgx.Rows) ([]models.Flat, error) {
	flats := make([]models.Flat, 0)
	for rows.Next() {
		flat := models.Flat{}
		err := rows.Scan(&flat.ID, &flat.Street)
		if err != nil {
			return nil, err
		}

		flats = append(flats, flat)
	}

	return flats, nil
}

func runMigrations(url string) error {

	m, err := getMigrations(url)
	if err != nil {
		return errors.Wrap(err, "failed to get migrations")
	}

	if err = m.Up(); err != migrate.ErrNoChange {
		return errors.Wrap(err, "failed to run migrations")
	}

	return nil
}

func getMigrations(url string) (*migrate.Migrate, error) {

	s := bindata.Resource(migrations.AssetNames(),
		func(name string) (data []byte, err error) {
			data, err = migrations.Asset(name)
			err = errors.Wrap(err, "failed to get migration data: "+name)
			return
		})
	d, err := bindata.WithInstance(s)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create migrations data driver")
	}

	m, err := migrate.NewWithSourceInstance("migrations", d, url)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create migrations instance")
	}

	return m, nil
}
