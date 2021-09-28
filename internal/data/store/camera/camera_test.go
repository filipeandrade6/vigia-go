package camera_test

import (
	"database/sql"
	"fmt"
	"net/url"
	"os"
	"testing"
	"time"

	"github.com/filipeandrade6/vigia-go/internal/sys/logger"
	_ "github.com/lib/pq"
	"github.com/ory/dockertest/v3"
	"go.uber.org/zap"
)

var (
	log   *zap.SugaredLogger
	pgURL *url.URL
)

func TestMain(m *testing.M) {
	code := 0
	defer func() {
		os.Exit(code)
	}()

	var err error
	log, err = logger.New("TEST")
	if err != nil {
		fmt.Printf("logger error: %s", err)
	}

	pgURL = &url.URL{
		Scheme: "postgres",
		User:   url.UserPassword("postgres", "secret"),
		Path:   "vigia",
	}
	q := pgURL.Query()
	q.Add("sslmode", "disable")
	pgURL.RawQuery = q.Encode()

	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalw("could not connect to docker", "ERROR", err)
	}

	pw, _ := pgURL.User.Password()
	runOpts := dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "latest",
		Env: []string{
			"POSTGRES_USER=" + pgURL.User.Username(),
			"POSTGRES_PASSWORD=" + pw,
			"POSTGRES_DB=" + pgURL.Path,
		},
	}

	resource, err := pool.RunWithOptions(&runOpts)
	if err != nil {
		log.Fatalw("could not start postgres container", "ERROR", err)
	}
	defer func() {
		err = pool.Purge(resource)
		if err != nil {
			log.Errorw("could not purge resource", "ERROR", err)
		}
	}()

	pgURL.Host = resource.Container.NetworkSettings.IPAddress

	pool.MaxWait = 10 * time.Second
	err = pool.Retry(func() error {
		db, err := sql.Open("postgres", pgURL.String())
		if err != nil {
			return err
		}
		return db.Ping()
	})
	if err != nil {
		log.Fatalw("could not connect to postgres server", "ERROR", err)
	}

	code = m.Run()
}

func TestRealbob(t *testing.T) {
	// all tests
}

// tests.NewUnit(t, dbc) onde t é t *testing.T (parametro da func) e dbc é as config do container
// defer cleanup(teardown)

// t.Log("given the need to work with User records")
// {
// 	testID := 0
// }
