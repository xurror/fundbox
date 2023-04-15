package models

import (
	"errors"
	"fmt"
	"getting-to-go/utils"
	"net/http"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

var db *gorm.DB

type Persistable struct {
	ID uuid.UUID `json:"id" gorm:"primary_key;type:uuid"`
}

type Auditable struct {
	Persistable
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

func (entity *Persistable) BeforeCreate(tx *gorm.DB) (err error) {
	entity.ID = uuid.New()
	return
}

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
	err := db.AutoMigrate(&User{}, &Fund{}, &Contributor{}, &Currency{}, &Amount{}, &Contribution{})
	if err != nil {
		panic(fmt.Sprintf("Failed to migrate database: %v", err))
	}
}

func HandleError(err error) error {
	if err == nil {
		return nil
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return utils.NewError(http.StatusNotFound, "Record not found")
	}
	return utils.NewError(http.StatusInternalServerError, err.Error())
}
