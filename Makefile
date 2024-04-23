build:
	@go build -o bin/bee-happy-api cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/bee-happy-api

api-docs:
	@swag init -g ./cmd/main.go -o ./docs

migration:
	@migrate create -ext sql -dir cmd/migrate/migrations $(filter-out $@,$(MAKECMDGOALS))

migration-force:
	@go run cmd/migrate/main.go force $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	@go run cmd/migrate/main.go up

migrate-down:
	@go run cmd/migrate/main.go down
