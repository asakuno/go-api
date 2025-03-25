package tests

import (
	"fmt"
	"os"
	"testing"

	"github.com/asakuno/go-api/config"
	"github.com/asakuno/go-api/constants"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetUpDatabaseConnection() *gorm.DB {
	var (
		dbUser, dbPass, dbHost, dbName, dbPort string
		getenv                                 = os.Getenv
		godotenv                               = godotenv.Load
	)

	if getenv("APP_ENV") != constants.ENUM_RUN_PRODUCTION {
		err := godotenv("../.env")
		if err != nil {
			panic("Error loading .env file: " + err.Error())
		}
	}

	dbUser = getenv("DB_USER")
	dbPass = getenv("DB_PASS")
	dbHost = getenv("DB_HOST")
	dbName = "test_go_api"
	dbPort = getenv("DB_PORT")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: config.SetupLogger(),
	})
	if err != nil {
		panic(err)
	}

	return db
}

func Test_DBConnection(t *testing.T) {
	db := SetUpDatabaseConnection()
	assert.NoError(t, db.Error, "Expected no error during database connection")
	assert.NotNil(t, db, "Expected a non-nil database connection")
}
