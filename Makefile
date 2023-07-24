# include example env values
include .env.example

# override env values if real .env file exists
ifneq (,$(wildcard ./.env))
    include .env
#    export
endif

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

# migratecreate
# migrate create -ext sql -dir db/migration -seq posts

.PHONY: test cover coverage migrateup migratedown sqlc
