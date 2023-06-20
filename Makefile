include .env.example
#include .env

test:
	go test -v ./...

cover:
	@go test -coverprofile=coverage.out ./... && go tool cover -html=coverage.out

coverage:
	@go test -cover ./...

migrateup:
	migrate -path db/migration -database "${DB_SOURCE}" -verbose up

migratedown:
	migrate -path db/migration -database "${DB_SOURCE}" -verbose down

sqlc:
	sqlc generate

.PHONY: test cover coverage migrateup migratedown sqlc
