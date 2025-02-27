# Define variables
APP_NAME = community-funds
DOCKER_IMAGE = $(APP_NAME)
DOCKER_TAG = latest
SERVER_BINARY = $(APP_NAME)
BUILD_DIR = .build

# Load environment variables from .env
ifneq (,$(wildcard .env))
	include .env
	export
endif

# Help command
.PHONY: help
help:
	@echo "Makefile for $(APP_NAME)"
	@echo ""
	@echo "Available commands:"
	@echo "  make build       	- Build the Go binary"
	@echo "  make swagger       - Build swagger docs"
	@echo "  make run         	- Run the application locally"
	@echo "  make docker-build 	- Build the Docker image"
	@echo "  make docker-run  	- Run the Docker container"
	@echo "  make test        	- Run unit tests"
	@echo "  make deploy      	- Deploy to $(DEPLOY_PLATFORM)"

# Build the Go binary
.PHONY: build
build:
	rm -rf $(BUILD_DIR) && mkdir -p $(BUILD_DIR)
	go build -ldflags '-s -w' -o $(BUILD_DIR)/$(SERVER_BINARY) ./cmd/server/main.go
	@echo "✅ Build complete: $(BUILD_DIR)/$(SERVER_BINARY)"

# Build the Go binary
.PHONY: debug-build
debug-build:
	rm -rf $(BUILD_DIR) && mkdir -p $(BUILD_DIR)
	CGO_ENABLED=0 go build -gcflags=all="-N -l" -o $(BUILD_DIR)/$(SERVER_BINARY) ./cmd/server/main.go
	@echo "✅ Debug build complete: $(BUILD_DIR)/$(SERVER_BINARY)"

# Build swagger docs
.PHONY: swagger
swagger:
	swag init --output ./docs --g ./cmd/server/main.go
	@echo "✅ Swagger docs generation complete: $(BUILD_DIR)/$(SERVER_BINARY)"

# Run the application locally
.PHONY: run
run:
	source .env && PROJECT_ROOT_DIR=$(shell pwd) go run ./cmd/server/main.go

# Run unit tests
.PHONY: test
test:
	go test ./...

# Build the Docker image
.PHONY: docker-build
docker-build:
	docker build -t $(DOCKER_IMAGE):$(DOCKER_TAG) .

# Run the Docker container
.PHONY: docker-run
docker-run:
	docker run -p 8080:8080 --env-file .env $(DOCKER_IMAGE):$(DOCKER_TAG)
