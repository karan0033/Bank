postgres:
	docker run --name=bank -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=karan123 -d postgres:16-alpine

createdb:
	docker exec -it bank createdb --username=root --owner=root bank

dropdb:
	docker exec -it bank dropdb bank

migrateup:
		migrate -path db/migration -database "postgres://root:karan123@localhost:5432/bank?sslmode=disable" up

migratedown:
		migrate -path db/migration -database "postgres://root:karan123@localhost:5432/bank?sslmode=disable" down

migrateup1:
		migrate -path db/migration -database "postgres://root:karan123@localhost:5432/bank?sslmode=disable" up 1

migratedown1:
		migrate -path db/migration -database "postgres://root:karan123@localhost:5432/bank?sslmode=disable" down 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/karan0033/bank/db/sqlc Store

.PHONY: postgres createdb dropdb sqlc test mock migratedown1 migrateup1
