package repository

import (
	"io"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct{
	URL string
}

func NewDBClient(cfg *Config) (*gorm.DB, io.Closer, error){
	db, err := gorm.Open(postgres.Open(cfg.URL))
	if err != nil{
		return nil, nil, err
	}
	inside, err := db.DB()
	if err != nil{
		return nil, nil, err
	}
	return db, inside, nil
}