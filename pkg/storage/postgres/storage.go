package postgres

import (
	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type storage struct {
	db *gorm.DB
}

func NewStorage(dsn string) (*storage, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&Plan{}, &Task{}); err != nil {
		return nil, errors.Wrap(err, "auto_migration")
	}

	return &storage{db: db}, nil
}
