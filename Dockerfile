FROM golang:1.23-alpine

# Set the working directory
WORKDIR /app



# Install necessary packages
RUN apk add --no-cache git bash gcc musl-dev make

# Install system dependencies: PostgreSQL client (psql) and make
# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the application
RUN go build -o bin/backend cmd/main.go

# Expose the application port
EXPOSE 8080

# Copy entrypoint script
COPY entrypoint.sh /app/entrypoint.sh
RUN chmod +x /app/entrypoint.sh

# Set entrypoint
ENTRYPOINT ["/app/entrypoint.sh"]


