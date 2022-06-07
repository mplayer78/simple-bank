postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=matt -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=matt --owner=matt simple-bank

dropdb:
	docker exec -it postgres12 dropdb simple-bank

migrateup:
	migrate -path db/migration -database "postgresql://matt:secret@localhost:5432/simple-bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://matt:secret@localhost:5432/simple-bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: postgres createdb dropdb migrateup migratedown generate