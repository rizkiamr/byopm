package sqlite

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type Config struct{}

func NewDB(config Config) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("byopm.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}