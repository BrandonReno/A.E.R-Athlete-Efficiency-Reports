package config

import (
	"fmt"
	"net/http"

	"github.com/caarlos0/env"
	"github.com/jinzhu/gorm"
)

type (
	Config struct {
		ServerPort      int    `env:"SERVER_PORT" envDefault:"5050"`
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
	db, err := gorm.Open(cfg.DBDialect, cfg.DBURL)
	if err != nil {
		fmt.Println("error")
		return nil, err
	}
	db.DB().SetMaxIdleConns(cfg.DBIdleConns)
	db.LogMode(cfg.DBDebugMode)
	return db, nil
}

func NewHTTPClient() *http.Client {
	return &http.Client{}
}
