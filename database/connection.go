package database

import (
	"context"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/mborders/logmatic"
	"os"
)

var DB *pgxpool.Pool
var l = logmatic.NewLogger()

func Migrate(migrations, url string) error {
	m, err := migrate.New(migrations, url)
	if err != nil {
		return err
	}

	err = m.Up()

	if err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}

func Connect() {
	DB_HOST, _ := os.LookupEnv("DB_HOST")
	DB_USER, _ := os.LookupEnv("DB_USER")
	DB_PASS, _ := os.LookupEnv("DB_PASS")
	DB_NAME, _ := os.LookupEnv("DB_NAME")
	DB_PORT, _ := os.LookupEnv("DB_PORT")
	DB_SSL, _ := os.LookupEnv("DB_SSL")

	dbUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", DB_USER, DB_PASS, DB_HOST, DB_PORT, DB_NAME, DB_SSL)
	connConfig, err := pgxpool.ParseConfig(dbUrl)

	if err != nil {
		l.Error("%s", err)
	}

	connConfig.ConnConfig.PreferSimpleProtocol = true

	conn, err := pgxpool.ConnectConfig(context.Background(), connConfig)
	//conn, err := pgxpool.Connect(context.Background(), dbUrl)

	DB = conn
	if err != nil {
		l.Error("Unable to connect to database: %v\n, err", err)
		os.Exit(1)
	}

	err = Migrate("file://./migrations", dbUrl)

	if err != nil {
		l.Error("%s", err)
	}
}
