postgres:
	podman run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=welsoncrelson -d postgres

createdb:
	podman exec -it postgres createdb --username=root --owner=root simple_bank

dropdb:
	podman exec -it postgres dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://root:welsoncrelson@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:welsoncrelson@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test