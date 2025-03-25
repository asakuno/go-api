package database

import (
	"github.com/asakuno/go-api/database/seeders"
	"gorm.io/gorm"
)

func Seed(db *gorm.DB) error {
	if err := seeders.ListUserSeeder(db); err != nil {
		return err
	}

	return nil
}