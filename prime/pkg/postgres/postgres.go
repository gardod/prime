package postgres

import (
	"database/sql"
	"fmt"
	"time"

	migrate "github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

type Config struct {
	Host            string        `mapstructure:"host"`
	Port            int           `mapstructure:"port"`
	User            string        `mapstructure:"user"`
	Password        string        `mapstructure:"password"`
	DBName          string        `mapstructure:"db_name"`
	SSLMode         string        `mapstructure:"ssl_mode"`
	MaxOpenConns    int           `mapstructure:"max_open_conns"`
	MaxIdleConns    int           `mapstructure:"max_idle_conns"`
	ConnMaxLifetime time.Duration `mapstructure:"conn_max_lifetime"`
	MigrateSource   string        `mapstructure:"migrate_source"`
}

func New(config Config) *sql.DB {
	logrus.Info("Setting up database")

	db, err := sql.Open("postgres", getURL(config))
	if err != nil {
		logrus.WithError(err).Fatal("Unable to open database")
	}

	db.SetMaxOpenConns(config.MaxOpenConns)
	db.SetMaxIdleConns(config.MaxIdleConns)
	db.SetConnMaxLifetime(time.Minute * config.ConnMaxLifetime)

	return db
}

func Migrate(config Config) {
	logrus.Info("Running database migrations")

	m, err := migrate.New(config.MigrateSource, getURL(config))
	if err != nil {
		logrus.WithError(err).Fatal("Unable to migrate database")
	}
	defer m.Close()

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		logrus.WithError(err).Fatal("Unable to migrate database")
	}
}

func getURL(config Config) string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.DBName,
		config.SSLMode,
	)
}
