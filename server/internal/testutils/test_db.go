package testutils

import (
	"context"
	"testing"
	"time"

	dbModule "community-funds/internal/db"

	"github.com/testcontainers/testcontainers-go"
	postgresContainer "github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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
		// postgresContainer.WithInitScripts(filepath.Join("testdata", "init-user-db.sh")),
		// postgresContainer.WithConfigFile(filepath.Join("testdata", "my-postgres.conf")),
		postgresContainer.WithDatabase("test_db"),
		postgresContainer.WithUsername("test_user"),
		postgresContainer.WithPassword("test_password"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(5*time.Second)),
	)

	if err != nil {
		t.Fatalf("Failed to start PostgreSQL container: %v", err)
	}

	// Get the connection string
	connStr, err := container.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		t.Fatalf("Failed to get connection string: %v", err)
	}

	// Connect to database
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to test database: %v", err)
	}

	// Run migrations
	dbModule.RunMigrations(db)

	return &TestDatabase{Container: container, DB: db}
}

// CleanupTestDB stops the test container
func CleanupTestDB(t *testing.T, testDB *TestDatabase) {
	ctx := context.Background()
	if err := testDB.Container.Terminate(ctx); err != nil {
		t.Fatalf("Failed to terminate PostgreSQL container: %v", err)
	}
}
