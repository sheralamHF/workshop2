# Backend Service

A simple Go backend service running on Docker.

## Requirements

- Docker
- Docker Compose
- Go 1.20+ (for local development)

## Running the Service

### Using Docker Compose (recommended)

Start the service:
```bash
docker-compose up -d
```

Stop the service:
```bash
docker-compose down
```

### Using Make commands

Start the service:
```bash
make docker-up
```

Stop the service:
```bash
make docker-down
```

Run locally without Docker:
```bash
make run
```

## API Endpoints

### Status Endpoint

- **URL**: `/status`
- **Method**: GET
- **Response**: `{"status": true}`
- **Status Code**: 200 OK

## Directory Structure

```
backend/
├── cmd/
│   └── api/            # Application entrypoints
│       └── main.go     # Main application file
├── Dockerfile          # Docker configuration
├── docker-compose.yml  # Docker Compose configuration
├── go.mod              # Go module definition
└── Makefile            # Build automation
``` 