package testutils

import (
	"context"
	"fmt"
	"testing"

	dbModule "community-funds/pkg/db"

	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	postgresContainer "github.com/testcontainers/testcontainers-go/modules/postgres"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// TestDatabase holds test container information
type TestDatabase struct {
	Container *postgresContainer.PostgresContainer
	DB        *gorm.DB
}

// SetupTestDB starts a test container for PostgreSQL
func SetupTestDB(t *testing.T) *TestDatabase {
	ctx := context.Background()

	container, err := postgresContainer.Run(ctx,
		"postgres:16-alpine",
		postgresContainer.WithDatabase("test_db"),
		postgresContainer.WithUsername("test_user"),
		postgresContainer.WithPassword("test_password"),
		postgresContainer.BasicWaitStrategies(),
	)
	testcontainers.CleanupContainer(t, container)
	require.NoError(t, err, fmt.Sprintf("Failed to start PostgreSQL container: %v", err))

	// Get the connection string
	connStr, err := container.ConnectionString(ctx, "sslmode=disable")
	require.NoError(t, err, fmt.Sprintf("Failed to get connection string: %v", err))

	// Connect to database
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	require.NoError(t, err, fmt.Sprintf("Failed to connect to test database: %v", err))

	// Run migrations
	err = dbModule.RunMigrations(db)
	require.NoError(t, err, fmt.Sprintf("Failed to run migrations: %v", err))

	return &TestDatabase{Container: container, DB: db}
}
