# Use the official Golang image as base
FROM golang:1.20-alpine AS builder

# Set the working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod ./
# Try to download dependencies, but don't fail if no go.sum exists yet
RUN go mod download || true

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o api ./cmd/api

# Use a smaller image for the final container
FROM alpine:latest

# Add necessary packages
RUN apk --no-cache add ca-certificates

# Set the working directory
WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/api .

# Expose the port the app runs on
EXPOSE 8080

# Run the binary
CMD ["./api"] 