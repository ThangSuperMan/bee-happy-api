build:
	@go build -o bin/bee-happy-api cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/bee-happy-api
