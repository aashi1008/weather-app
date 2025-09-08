# Basic Makefile for Weather API (Updated for Docker Compose v2)

APP_NAME=weather-app

build:
	go build -o bin/$(APP_NAME) ./cmd/main.go

run:
	go run ./cmd/main.go

test:
	go test ./tests/...

# Updated Docker Compose commands (removed dash)
up:
	docker compose up -d

down:
	docker compose down

build-docker:
	docker compose build

test-docker:
	docker exec $(APP_NAME) go test ./tests/...

clean:
	rm -rf bin/
	docker compose down --rmi all -v

help:
	@echo "Available commands:"
	@echo "  build       - Build the application"
	@echo "  run         - Run locally"
	@echo "  test        - Run tests"
	@echo "  up          - Start Docker services"
	@echo "  down        - Stop Docker services"
	@echo "  build-docker - Build Docker images"
	@echo "  test-docker - Run tests in Docker"
	@echo "  clean       - Clean up everything"

.PHONY: build run test up down build-docker test-docker clean help
