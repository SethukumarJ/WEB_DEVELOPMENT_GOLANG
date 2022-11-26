package storage

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host     string
	Port     int
	Password string
	User     string
	DBName   string
	SSLMode  string 

	}


	func newConnection(cfg Config) (*gorm.DB, error) {
		dsn := fmt.Sprintf("host=%s user=%s password dbname=%s port=%d sslmode=%s", 
		cfg.Host, cfg.User, cfg.Password, cfg.DBName, cfg.Port, cfg.SSLMode)

		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			return db,err
		}
		return db, nil

	}