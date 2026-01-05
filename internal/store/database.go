package store

import (
	"context"
	"database/sql"
	"fmt"
	"io/fs"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

const connectionString = "host=localhost port=5432 user=postgres password=postgres dbname=fe_go sslmode=disable"

func Open() (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig(connectionString)
	if err != nil {
		return nil, fmt.Errorf("db: parse config: %w", err)
	}

	db, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return nil, fmt.Errorf("db: open: %w", err)
	}

	if err := db.Ping(context.Background()); err != nil {
		return nil, fmt.Errorf("db: ping: %w", err)
	}

	fmt.Println("Connected to Database...")

	return db, nil
}

func MigrateFS(migrationsFS fs.FS, dir string) error {
	// Create a *sql.DB for goose migrations using the same connection string
	config, err := pgx.ParseConfig(connectionString)
	if err != nil {
		return fmt.Errorf("migrate: parse config: %w", err)
	}

	var sqlDb = stdlib.OpenDB(*config)
	defer sqlDb.Close()

	goose.SetBaseFS(migrationsFS)
	defer func() {
		goose.SetBaseFS(nil)
	}()

	err = goose.SetDialect("postgres")
	if err != nil {
		return fmt.Errorf("migrate: set dialect: %w", err)
	}

	err = goose.Up(sqlDb, dir)
	if err != nil {
		return fmt.Errorf("goose up: %w", err)
	}

	return nil
}

// Migrate runs migrations on the provided database connection using the migrations FS
func Migrate(db *sql.DB, migrationsFS fs.FS, dir string) error {
	goose.SetBaseFS(migrationsFS)
	defer func() {
		goose.SetBaseFS(nil)
	}()

	err := goose.SetDialect("postgres")
	if err != nil {
		return fmt.Errorf("migrate: set dialect: %w", err)
	}

	err = goose.Up(db, dir)
	if err != nil {
		return fmt.Errorf("goose up: %w", err)
	}

	return nil
}
