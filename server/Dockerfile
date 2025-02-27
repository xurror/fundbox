# ---- BUILD STAGE ----
FROM golang:1.24-alpine AS builder

# Install dependencies
RUN apk --no-cache add ca-certificates git

# Set working directory
WORKDIR /app

# Copy go modules & install dependencies
COPY go.mod go.sum ./
RUN go mod tidy

# Copy the rest of the application
COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o community-funds ./cmd/server/main.go

# ---- DEPLOYMENT STAGE ----
FROM alpine:latest

# Install CA certificates
RUN apk --no-cache add ca-certificates

ENV PROJECT_ROOT_DIR=/root/

# Set working directory
WORKDIR $PROJECT_ROOT_DIR

# Copy built binary from builder stage
COPY --from=builder /app/community-funds .
COPY --from=builder /app/docs docs

# Expose API port
EXPOSE 8080

# Run the application
CMD ["./community-funds"]
