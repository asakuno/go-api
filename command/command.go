package command

import (
	"log"
	"os"

	"github.com/asakuno/go-api/constants"
	"github.com/asakuno/go-api/database/migrations"
	"github.com/samber/do"
	"gorm.io/gorm"
)

func Commands(injector *do.Injector) bool {
	db := do.MustInvokeNamed[*gorm.DB](injector, constants.DB)

	migrate := false
	run := false

	for _, arg := range os.Args[1:] {
		if arg == "--migrate" {
			migrate = true
		}
		if arg == "--run" {
			run = true
		}
	}

	if migrate {
		if err := migrations.Migrate(db); err != nil {
			log.Fatalf("error migration: %v", err)
		}
		log.Println("migration success")
	}

	if run {
		return true
	}

	return false
}
