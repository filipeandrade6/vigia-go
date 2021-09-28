package tests

import (
	"net/url"
	"os"
	"testing"
	"time"

	"github.com/filipeandrade6/vigia-go/internal/sys/logger"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/ory/dockertest/v3"
	"go.uber.org/zap"
)

// var (
// 	DockerDBConn *dockerDBConn
// )

// type dockerDBConn struct {
// 	Conn *sqlx.DB
// 	Log *zap.SugaredLogger
// }

// func TestMain(m *testing.M) {
// 	pool, resource := initDB()
// 	code := m.Run()
// 	err := closeDB(pool, resource)
// 	if err !=
// 	os.Exit(code)
// }

const (
	Success = "\u2713"
	Failed  = "\u2717"
)

func New(t *testing.T) (*zap.SugaredLogger, *sqlx.DB, func()) {
	code := 0
	defer func() {
		os.Exit(code)
	}()

	var err error
	log, err := logger.New("TEST")
	if err != nil {
		t.Fatalf("logger error: %s", err)
	}

	pgURL := &url.URL{
		Scheme: "postgres",
		User:   url.UserPassword("postgres", "secret"),
		Path:   "vigia",
	}
	q := pgURL.Query()
	q.Add("sslmode", "disable")
	pgURL.RawQuery = q.Encode()

	pool, err := dockertest.NewPool("")
	if err != nil {
		t.Fatalf("could not connect to docker: %v", err)
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
		t.Fatalf("could not start postgres container: %v", err)
	}
	defer func() {
		err = pool.Purge(resource)
		if err != nil {
			t.Errorf("could not purge resource: %v", err)
		}
	}()

	pgURL.Host = resource.Container.NetworkSettings.IPAddress

	pool.MaxWait = 10 * time.Second
	var db *sqlx.DB
	err = pool.Retry(func() error {
		db, err = sqlx.Open("postgres", pgURL.String())
		if err != nil {
			return err
		}
		return db.Ping()
	})
	if err != nil {
		t.Fatalf("could not connect to postgres server: %v", err)
	}

	teardown := func() {
		if err := pool.Purge(resource); err != nil {
			t.Errorf("could not purge the resources: %v", err)
		}
	}

	return log, db, teardown
}
