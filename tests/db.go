package tests

import (
	"fmt"
	"testing"

	"github.com/asakuno/go-api/config"
	"github.com/asakuno/go-api/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetUpDatabaseConnection() *gorm.DB {
	var (
		dbUser, dbPass, dbHost, dbName, dbPort string
	)
	dbUser = "root"
	dbPass = "password"
	dbHost = "mysql"
	dbName = "test_go_api"
	dbPort = "3306"

	dsnForCreation := fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPass, dbHost, dbPort)

	tempDB, err := gorm.Open(mysql.Open(dsnForCreation), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	tempDB.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s;", dbName))

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: config.SetupLogger(),
	})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entity.User{})

	return db
}

func Test_DBConnection(t *testing.T) {
	db := SetUpDatabaseConnection()
	assert.NoError(t, db.Error, "Expected no error during database connection")
	assert.NotNil(t, db, "Expected a non-nil database connection")
}
