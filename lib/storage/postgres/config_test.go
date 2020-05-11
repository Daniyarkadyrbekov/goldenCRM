package postgres

import (
	"net"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestConfigCheck(t *testing.T) {

	require.EqualError(t,
		(&Config{}).Check(),
		"'db.host' was not set")

	require.EqualError(t,
		(&Config{Host: "h"}).Check(),
		"'db.port' invalid value")

	require.EqualError(t,
		(&Config{Host: "h", Port: "1"}).Check(),
		"'db.name' was not set")

	require.EqualError(t,
		(&Config{Host: "h", Port: "1", Name: "n"}).Check(),
		"'db.user' was not set")

	require.EqualError(t,
		(&Config{Host: "h", Port: "1", Name: "n", User: "u"}).Check(),
		"'db.password' was not set")

	require.EqualError(t,
		(&Config{Host: "h", Port: "1", Name: "n", User: "u", Password: "p"}).Check(),
		"'db.ssl-mode' was not set")

	require.NoError(t,
		(&Config{Host: "h", Port: "1", Name: "n", User: "u", Password: "p", SslMode: "s"}).Check(),
		"test ok")

	require.NoError(t,
		(&Config{Host: "h", Port: "1", Name: "n", User: "u", Password: "p", SslMode: "s", HealthCheckPeriod: time.Second, MaxConnections: 2}).Check(),
		"test ok")
}

func TestConfigConnURL(t *testing.T) {

	fnCheckUrl := func(host string, u *url.URL) {
		require.Equal(t,
			&url.URL{
				Scheme:     "postgres",
				Opaque:     "",
				User:       url.UserPassword("user name", "password data"),
				Host:       net.JoinHostPort(host, "3000"),
				Path:       "/base+name",
				RawPath:    "",
				ForceQuery: false,
				RawQuery:   "",
				Fragment:   ""},
			u)
	}

	for _, host := range []string{
		"::1", // ipv6
		"localhost",
	} {

		src := &Config{
			Host:              host,
			Port:              "3000",
			Name:              "base name",
			User:              "user name",
			Password:          "password data",
			SslMode:           "disable",
			HealthCheckPeriod: time.Second,
			MaxConnections:    10,
		}

		{
			u, err := url.Parse(src.PoolConnURL())
			require.NoError(t, err)

			q := url.Values{}
			q.Set("pool_health_check_period", "1s")
			q.Set("pool_max_conns", "10")
			q.Set("sslmode", "disable")
			q.Set("statement_cache_capacity", "13")
			q.Set("statement_cache_mode", "describe")
			require.Equal(t, q, u.Query())

			u.RawQuery = ""
			fnCheckUrl(host, u)
		}

		{
			u, err := url.Parse(src.ConnURL())
			require.NoError(t, err)

			q := url.Values{}
			q.Set("sslmode", "disable")
			require.Equal(t, q, u.Query())

			u.RawQuery = ""
			fnCheckUrl(host, u)
		}
	}
}
