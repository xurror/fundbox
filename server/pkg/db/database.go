package db

import (
	"community-funds/config"
	"community-funds/pkg/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabase(cfg *config.Config) *gorm.DB {
	db, err := gorm.Open(postgres.Open(cfg.DatabaseDSN)) // &gorm.Config{
	// 	Logger: logger.New(log, logger.Config{
	// 		SlowThreshold:             time.Second,   // Slow SQL threshold
	// 		LogLevel:                  logger.Silent, // Log level
	// 		IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
	// 		ParameterizedQueries:      true,          // Don't include params in the SQL log
	// 		Colorful:                  true,          // false,         // Disable color
	// 	}),
	// },

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	err = RunMigrations(db)
	if err != nil {
		log.Panicf("Failed to migrate database: %v", err)
	}

	return db
}

func RunMigrations(db *gorm.DB) error {
	err := db.AutoMigrate(
		&models.User{},
		&models.Contribution{},
		&models.Fund{},
	)

	return err
}
