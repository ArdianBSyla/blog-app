package helper

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/personal/blog-app/config"
	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewGormDB(cfg *config.Config) (*gorm.DB, error) {
	gormLogLevel := logger.Silent

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,  // Slow SQL threshold
			LogLevel:                  gormLogLevel, // Log level
			IgnoreRecordNotFoundError: true,         // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,        // Disable color
		},
	)

	postgresDbConnection, err := NewDBConnection(cfg)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to open Postgres database connection")
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: postgresDbConnection,
	}), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return nil, errors.WithMessage(err, "failed to open GORM database connection")
	}

	return gormDB, nil
}

// NewDBConnection returns a connection to Postgres database
func NewDBConnection(conf *config.Config) (*sql.DB, error) {
	return sql.Open("postgres", fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		conf.DB_User,
		conf.DB_Password,
		conf.DB_Host,
		conf.DB_Port,
		conf.DB_Name,
	))
}
