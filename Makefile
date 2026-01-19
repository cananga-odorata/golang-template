.PHONY: help build run test migrate-up migrate-down lint docker-up docker-down clean

# Default target
help:
	@echo "Golang Modular Monolith Template - Available Commands"
	@echo ""
	@echo "Development:"
	@echo "  make run          - Run the application"
	@echo "  make build        - Build the application binary"
	@echo "  make test         - Run tests with coverage"
	@echo "  make lint         - Run linter"
	@echo ""
	@echo "Database:"
	@echo "  make migrate-up   - Run database migrations"
	@echo "  make migrate-down - Rollback database migrations"
	@echo ""
	@echo "Docker:"
	@echo "  make docker-up    - Start Docker containers"
	@echo "  make docker-down  - Stop Docker containers"
	@echo "  make docker-build - Build Docker image"
	@echo ""
	@echo "Cleanup:"
	@echo "  make clean        - Remove build artifacts"

# Build the application
build:
	@echo "Building..."
	go build -o bin/api cmd/api/main.go
	@echo "Build complete: bin/api"

# Run the application
run:
	go run cmd/api/main.go

# Run tests with coverage
test:
	go test -v -race -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report: coverage.html"

# Run tests without coverage (faster)
test-fast:
	go test -v ./...

# Run database migrations up
migrate-up:
	@if [ -z "$(DATABASE_URL)" ]; then \
		echo "DATABASE_URL is not set. Example:"; \
		echo "  export DATABASE_URL='postgres://dev:dev@localhost:5432/affiliate_db?sslmode=disable'"; \
		exit 1; \
	fi
	migrate -path migrations -database "$(DATABASE_URL)" up

# Run database migrations down
migrate-down:
	@if [ -z "$(DATABASE_URL)" ]; then \
		echo "DATABASE_URL is not set"; \
		exit 1; \
	fi
	migrate -path migrations -database "$(DATABASE_URL)" down

# Create a new migration
migrate-create:
	@if [ -z "$(name)" ]; then \
		echo "Usage: make migrate-create name=migration_name"; \
		exit 1; \
	fi
	migrate create -ext sql -dir migrations -seq $(name)

# Run linter
lint:
	golangci-lint run

# Format code
fmt:
	go fmt ./...

# Tidy dependencies
tidy:
	go mod tidy

# Start Docker containers
docker-up:
	docker-compose up -d

# Stop Docker containers
docker-down:
	docker-compose down

# Build Docker image
docker-build:
	docker build -t go-app:latest .

# Clean build artifacts
clean:
	rm -rf bin/
	rm -f coverage.out coverage.html

# Install development dependencies
install-tools:
	go install github.com/golang-migrate/migrate/v4/cmd/migrate@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
