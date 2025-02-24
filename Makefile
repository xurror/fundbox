# Define variables
APP_NAME = community-funds
DOCKER_IMAGE = your-dockerhub-username/$(APP_NAME)
DOCKER_TAG = latest
SERVER_BINARY = $(APP_NAME)
SERVER_DIR = server
BUILD_DIR = $(SERVER_DIR)/build
DEPLOY_PLATFORM = render # Change to 'render' if deploying to Render

# Load environment variables from .env
ifneq (,$(wildcard $(SERVER_DIR)/.env))
	include $(SERVER_DIR)/.env
	export
endif

# Help command
.PHONY: help
help:
	@echo "Makefile for $(APP_NAME)"
	@echo ""
	@echo "Available commands:"
	@echo "  make build       - Build the Go binary"
	@echo "  make run         - Run the application locally"
	@echo "  make docker-build - Build the Docker image"
	@echo "  make docker-run  - Run the Docker container"
	@echo "  make test        - Run unit tests"
	@echo "  make deploy      - Deploy to $(DEPLOY_PLATFORM)"

# Build the Go binary
.PHONY: build
build:
	cd $(SERVER_DIR) && go build -ldflags '-s -w' -o $(BUILD_DIR)/$(SERVER_BINARY) ./cmd/server/main.go
	@echo "✅ Build complete: $(BUILD_DIR)/$(SERVER_BINARY)"

# Run the application locally
.PHONY: run
run:
	cd $(SERVER_DIR) && go run ./cmd/server/main.go

# Run unit tests
.PHONY: test
test:
	cd $(SERVER_DIR) && go test ./...

# Build the Docker image
.PHONY: docker-build
docker-build:
	cd $(SERVER_DIR) && docker build -t $(DOCKER_IMAGE):$(DOCKER_TAG) .

# Run the Docker container
.PHONY: docker-run
docker-run:
	docker run -p 8080:8080 --env-file $(SERVER_DIR)/.env $(DOCKER_IMAGE):$(DOCKER_TAG)

# Deploy to cloud provider
.PHONY: deploy
deploy:
ifeq ($(DEPLOY_PLATFORM), koyeb)
	@echo "Deploying to Koyeb..."
	koyeb service deploy --name $(APP_NAME) --docker $(DOCKER_IMAGE):$(DOCKER_TAG)
else ifeq ($(DEPLOY_PLATFORM), render)
	@echo "Deploying to Render..."
	git push render main
else
	@echo "❌ Unknown deployment platform: $(DEPLOY_PLATFORM)"
endif
