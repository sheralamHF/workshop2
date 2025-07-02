.PHONY: build run docker-build docker-run docker-up docker-down clean test

# Build the application
build:
	go build -o api ./cmd/api

# Run the application locally
run:
	go run ./cmd/api

# Build the Docker image
docker-build:
	docker build -t workshop2-api .

# Run the Docker container
docker-run:
	docker run -p 8080:8080 workshop2-api

# Start services with Docker Compose
docker-up:
	docker-compose up -d

# Stop services with Docker Compose
docker-down:
	docker-compose down

# Clean build artifacts
clean:
	rm -f api

# Run tests
test:
	go test -v ./... 