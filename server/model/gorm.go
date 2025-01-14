package model

import (
	"errors"
	"fmt"
	"getting-to-go/config"
	_type "getting-to-go/type"
	"net/http"

	log "github.com/sirupsen/logrus"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

func (entity *Persistable) BeforeCreate(tx *gorm.DB) (err error) {
	entity.Id = uuid.New()
	return
}

func NewDB(c *config.AppConfig) *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.Database.Host, c.Database.Port, c.Database.User, c.Database.Password, c.Database.Name, c.Database.Sslmode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic(fmt.Sprintf("Failed to connect to database: %v", err))
	}

	runMigrations(db)

	return db
}

func runMigrations(db *gorm.DB) {
	err := db.AutoMigrate(&User{}, &Fund{}, &Contribution{}, &Currency{}, &Amount{})
	if err != nil {
		log.Panic(fmt.Sprintf("Failed to migrate database: %v", err))
	}
}

func HandleError(err error) error {
	if err == nil {
		return nil
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return _type.NewAppError(http.StatusNotFound, "Record not found")
	}
	return _type.NewAppError(http.StatusInternalServerError, err.Error())
}
