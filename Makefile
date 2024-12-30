.PHONY: build test run migration migrate-up migrate-down docker-build docker-up docker-down

build:
	@go build -o bin/backend cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/backend

migration:
	@migrate create -ext sql -dir migrate/migrations $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	@go run migrate/main.go up

migrate-down:
	@go run migrate/main.go down

# Docker commands
docker-build:
	docker-compose build

docker-up:
	docker-compose up -d

docker-down:
	docker-compose down

# Run migrations in Docker
docker-migrate-up:
	docker-compose exec app ./backend migrate up

docker-migrate-down:
	docker-compose exec app ./backend migrate down

# View logs
docker-logs:
	docker-compose logs -f

# Clean up Docker volumes
docker-clean:
	docker-compose down -v