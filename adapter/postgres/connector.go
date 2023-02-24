package postgres

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/pgx" //driver pgx used to run migrations
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/spf13/viper"
)

// PoolInterface is a wrapper to PgxPool to create test mocks
type PoolInterface interface {
	Close()
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	QueryFunc(
		ctx context.Context,
		sql string,
		args []interface{},
		scans []interface{},
		f func(pgx.QueryFuncRow) error,
	) (pgconn.CommandTag, error)
	SendBatch(ctx context.Context, b *pgx.Batch) pgx.BatchResults
	Begin(ctx context.Context) (pgx.Tx, error)
	BeginFunc(ctx context.Context, f func(pgx.Tx) error) error
	BeginTxFunc(ctx context.Context, txOptions pgx.TxOptions, f func(pgx.Tx) error) error
}

// GetConnection return connection poll from postgres drive PGX
func GetConnection(ctx context.Context) *pgxpool.Pool {
	dbUrl := viper.GetString("database.url")
	conn, err := pgxpool.Connect(ctx, "postgres"+dbUrl)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect ot database: %v", err)
		os.Exit(1)
	}

	return conn
}

// RunMigrations function that runs all migrations on path database/migrations
func RunMigrations() {
	dburl := viper.GetString("database.url")
	m, err := migrate.New("file://database/migrations", "pgx"+dburl)
	if err != nil {
		log.Println(err)
	}

	if err := m.Up(); err != nil {
		log.Println(err)
	}
}
