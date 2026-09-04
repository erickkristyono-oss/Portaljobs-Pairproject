run:
	go run ./cmd/api

build:
	go build -o bin/server ./cmd/api

test:
	go test ./... -v

tidy:
	go mod tidy

up:
	docker compose up --build

down:
	docker compose down

.PHONY: run build test tidy up down
