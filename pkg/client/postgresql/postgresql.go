package postgresql

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"time"
)

type Client interface {
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Begin(ctx context.Context) (pgx.Tx, error)
}

type TmpDBConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
}

func NewClient(ctx context.Context, retries int, delay time.Duration, cfg TmpDBConfig) (*pgxpool.Pool, error) {
	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database)
	var err error

	for i := 0; i < retries; i++ {
		db, err := pgxpool.New(ctx, connString)
		if err == nil {

			err = db.Ping(ctx)
			if err == nil {

				return db, nil
			}

		}
		log.Printf("Failed to connect to DB, retrying in %v... (%d/%d)", delay, i+1, retries)
		time.Sleep(delay)
	}

	return nil, fmt.Errorf("could not connect to DB after %d attempts: %w", retries, err)
}
