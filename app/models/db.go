package models

import (
	"errors"
	"fmt"
	"getting-to-go/app/utils"
	"net/http"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	_ "github.com/lib/pq"
)

var db *gorm.DB

// Connect initializes the database connection
func Connect(host, port, user, password, dbname string) error {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(fmt.Sprintf("Failed to connect to database: %v", err))
	}

	return nil
}

// Run Migration
func RunMigrations() {
	db.AutoMigrate(&User{})
}

func HandleError(err error) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return utils.NewError(http.StatusNotFound, "Record not found")
	}
	return utils.NewError(http.StatusInternalServerError, err.Error())
}
