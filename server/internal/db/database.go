package db

import (
	"community-funds/internal/config"
	"community-funds/internal/models"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDatabase(cfg *config.Config, log *logrus.Logger) *gorm.DB {
	db, err := gorm.Open(postgres.Open(cfg.DatabaseDSN), &gorm.Config{
		Logger: logger.New(log, logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,          // Don't include params in the SQL log
			Colorful:                  true,          // false,         // Disable color
		}),
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	runMigrations(db, log)

	return db
}

func runMigrations(db *gorm.DB, log *logrus.Logger) {
	err := db.AutoMigrate(
		&models.User{},
		&models.Contribution{},
		&models.Fund{},
	)

	if err != nil {
		log.Panicf("Failed to migrate database: %v", err)
	}
}
