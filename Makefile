DB_URL=postgresql://root:secret@localhost:5432/go_cms?sslmode=disable

test:
	go test -v ./...

cover:
	@go test -coverprofile=coverage.out ./... && go tool cover -html=coverage.out

coverage:
	@go test -cover ./...

migrateup:
	migrate -path db/migration -database "${DB_URL}" -verbose up

migratedown:
	migrate -path db/migration -database "${DB_URL}" -verbose down

sqlc:
	sqlc generate

.PHONY: test cover coverage migrateup migratedown sqlc
