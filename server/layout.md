community-funds/
│── cmd/
│   │── server/               # Main entry point for the application
│   │   ├── main.go           # Bootstraps the app with Uber Fx
│── internal/
│   │── config/               # Configuration management
│   │   ├── config.go         # Load environment variables & configs
│   │── server/               # Gin server setup
│   │   ├── server.go         # Initializes Gin and registers routes
│   │── routes/               # API routes
│   │   ├── routes.go         # Route definitions
│   │── handlers/             # HTTP handlers (controllers)
│   │   ├── funds.go          # Handlers for fund management
│   │   ├── users.go          # Handlers for user operations
│   │── services/             # Business logic layer
│   │   ├── funds_service.go  # Business logic for fund management
│   │   ├── users_service.go  # Business logic for users
│   │── repositories/         # Data access layer
│   │   ├── funds_repo.go     # CRUD operations for funds
│   │   ├── users_repo.go     # CRUD operations for users
│   │── models/               # Database models
│   │   ├── fund.go           # Fund model
│   │   ├── user.go           # User model
│   │── db/                   # Database connection setup
│   │   ├── database.go       # Database connection logic
│   │── middlewares/          # Middleware functions (auth, logging, etc.)
│   │   ├── auth.go           # Authentication middleware
│   │   ├── logging.go        # Logging middleware
│── pkg/                      # Reusable utility packages
│   │── logger/               # Logging setup
│   │   ├── logger.go         # Logger configuration
│   │── utils/                # Helper functions
│   │   ├── hash.go           # Password hashing
│   │   ├── response.go       # Standard API response helpers
│── api/                      # API documentation (Swagger, Postman)
│   │── swagger.yaml          # Swagger/OpenAPI spec
│── test/                     # Test files
│   │── integration/          # Integration tests
│   │── unit/                 # Unit tests
│── configs/                  # Configuration files
│   │── config.yaml           # Application configs
│── migrations/               # Database migrations
│── Dockerfile                # Docker container setup
│── go.mod                    # Go module dependencies
│── go.sum                    # Go module checksums
│── Makefile                   # Task automation (run, build, test)
│── README.md                  # Project documentation
