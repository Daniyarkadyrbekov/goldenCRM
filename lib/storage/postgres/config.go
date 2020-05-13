package postgres

import (
	"net"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type Config struct {
	Host              string        `mapstructure:"host"`
	Port              string        `mapstructure:"port"`
	Name              string        `mapstructure:"name"`
	User              string        `mapstructure:"user"`
	Password          string        `mapstructure:"password"`
	SslMode           string        `mapstructure:"ssl-mode"`
	Schema            string        `mapstructure:"schema"` // database schema. example: 'files'. default: 'public'
	HealthCheckPeriod time.Duration `mapstructure:"health-check"`
	MaxConnections    int           `mapstructure:"max-connections"`
}

func NewConfig(src *viper.Viper) (*Config, error) {

	conf := &Config{}
	if err := src.Unmarshal(conf); err != nil {
		return nil, errors.Wrap(err, "failed to parse DB config")
	}

	return conf, nil
}

func (c *Config) Check() error {

	if len(c.Host) == 0 {
		return errors.New("'db.host' was not set")
	}

	if len(c.Port) == 0 {
		return errors.New("'db.port' invalid value")
	}

	if len(c.Name) == 0 {
		return errors.New("'db.name' was not set")
	}

	if val := strings.TrimSpace(c.Schema); val != c.Schema {
		return errors.Errorf("'db.schema' invalid value: '%s'", c.Schema)
	}

	if len(c.User) == 0 {
		return errors.New("'db.user' was not set")
	}

	if len(c.Password) == 0 {
		return errors.New("'db.password' was not set")
	}

	if len(c.SslMode) == 0 {
		return errors.New("'db.ssl-mode' was not set")
	}

	return nil
}

func (c *Config) PoolConnURL() string {
	return c.URL(true, true).String()
}

func (c *Config) Addr() string {
	return net.JoinHostPort(c.Host, c.Port)
}

func (c *Config) ConnURL() string {
	return c.URL(false, true).String()
}

func (c *Config) ConnURLWithoutSchema() string {
	return c.URL(false, false).String()
}

func (c *Config) URL(poll, withSchema bool) *url.URL {

	q := url.Values{}
	q.Set("sslmode", c.SslMode)
	if c.Schema != "" && withSchema {
		q.Set("search_path", c.Schema)
	}

	if poll {
		q.Set("statement_cache_mode", "describe")
		q.Set("statement_cache_capacity", "13") // WARN: change it, if will add new query

		// in the pgx library exist default value
		if c.MaxConnections > 0 {
			q.Set("pool_max_conns", strconv.Itoa(c.MaxConnections))
		}

		// in the pgx library exist default value
		if c.HealthCheckPeriod > 0 {
			q.Set("pool_health_check_period", c.HealthCheckPeriod.String())
		}
	}

	return &url.URL{
		Scheme:   "postgres",
		Host:     c.Addr(),
		User:     url.UserPassword(c.User, c.Password),
		Path:     url.QueryEscape(c.Name),
		RawQuery: q.Encode(),
	}
}
