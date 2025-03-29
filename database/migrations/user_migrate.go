package migrations

import (
	"github.com/asakuno/go-api/entities"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(
		&entities.User{},
	); err != nil {
		return err
	}

	return nil
}
