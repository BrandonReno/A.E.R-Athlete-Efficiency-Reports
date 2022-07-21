package config

import (
	"net/http"

	"github.com/caarlos0/env"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type (
	Config struct {
		ServerPort      int    `env:"SERVER_PORT" envDefault:"9090"`
		ServerDebugMode bool   `env:"SERVER_DEBUG_MODE" envDefault:"false"`
		DBURL           string `env:"DATABASE_URL" envDefault:"postgres://postgres:postgres@postgres/postgres?sslmode=disable"`
		DBDialect       string `env:"DATABASE_DIALECT" envDefault:"postgres"`
		DBIdleConns     int    `env:"DATABASE_IDLE_CONNS" envDefault:"5"`
		DBDebugMode     bool   `env:"DATABASE_DEBUG_MODE" envDefault:"false"`
	}
)

func NewConfig() (*Config, error) {
	cfg := &Config{}
	if err := LoadConfig(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}

func LoadConfig(cfgs ...interface{}) error {
	for _, c := range cfgs {
		err := env.Parse(c)
		if err != nil {
			return err
		}
	}
	return nil
}

func NewDatabaseClient(cfg *Config) (*gorm.DB, error) {
	opt := gorm.Config{
		PrepareStmt: true,
	}
	db, err := gorm.Open(postgres.Open(cfg.DBURL), &opt)
	if err != nil {
		return nil, err
	}
	inner, err := db.DB()
	if err != nil {
		return nil, err
	}
	inner.SetMaxIdleConns(cfg.DBIdleConns)
	return db, nil
}

func NewHTTPClient() *http.Client {
	return &http.Client{}
}
