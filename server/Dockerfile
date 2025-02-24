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
RUN go build -o community-funds ./cmd/server/main.go

# ---- DEPLOYMENT STAGE ----
FROM alpine:latest

# Install CA certificates
RUN apk --no-cache add ca-certificates

# Set working directory
WORKDIR /root/

# Copy built binary from builder stage
COPY --from=builder /app/community-funds .

# Expose API port
EXPOSE 8080

# Run the application
CMD ["./community-funds"]
