FROM golang:1.22.2 AS build-stage
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o bee-happy-api cmd/main.go

FROM build-stage AS run-test-stage
RUN go test -v ./...

FROM scratch AS run-release-stage
WORKDIR /app

COPY --from=build-stage /app/bee-happy-api /app/bee-happy-api

EXPOSE 3000

CMD ["/app/bee-happy-api"]
