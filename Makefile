DB_URL=postgresql://root:secret@localhost:5432/go_cms?sslmode=disable

test:
	go test -v ./...

migrateup:
	migrate -path db/migration -database "${DB_URL}" -verbose up

migratedown:
	migrate -path db/migration -database "${DB_URL}" -verbose down

.PHONY: test migrateup migratedown
