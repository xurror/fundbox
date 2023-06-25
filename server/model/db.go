package model

import (
	"errors"
	"fmt"
	"getting-to-go/config"
	_type "getting-to-go/type"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

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

func NewDB(c *config.AppConfig) *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.Database.Host, c.Database.Port, c.Database.User, c.Database.Password, c.Database.Name)

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
