startdb: 
	docker compose -f ./db/dev-db/docker-compose.yml up -d 

stopdb: 
	docker compose -f ./db/dev-db/docker-compose.yml down -v 

createdb: 
	docker exec -it postgres createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://root:rootpass@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:rootpass@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test: 
	go test -v -cover ./...

.PHONY: startdevdb createdb stopdb dropdb migrateup migratedown sqlc