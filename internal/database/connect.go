// Package database provides functionality for interacting with the Postgres
// database.
package database

import (
	"fmt"

	"github.com/NooFreeNames/MultiSocket/internal/database/models"
	"github.com/NooFreeNames/MultiSocket/internal/utils"
	"github.com/NooFreeNames/MultiSocket/pkg/logger"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

// initDB initializes the connection to the Postgres database and makes
// migration. If it was not possible to connect to the database logs errors and
// shuts down the application
func initDB() {
	var err error
	db, err = gorm.Open("postgres",
		fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
			utils.GetOrLog("DB_HOST"),
			utils.GetOrLog("DB_PORT"),
			utils.GetOrLog("DB_USER"),
			utils.GetOrLog("DB_NAME"),
			utils.GetOrLog("DB_PASSWORD"),
		),
	)
	if err != nil {
		logger.Fatal("Failed to connect to database: %s", err)
	}
	db.AutoMigrate(&models.User{}, &models.Channel{})
}

// GetDB returns a pointer to a gorm.DB object, which represents a connection
// to the database. If the connection has not yet been established, the initDB()
// function is called.
func GetDB() *gorm.DB {
	if db == nil {
		initDB()
	}
	return db
}
