migrate_up:
	migrate -path db/migrations -database "postgresql://root:password@localhost:5432/ticket-assignment?sslmode=disable" -verbose up
migrate_down:
	migrate -path db/migrations -database "postgresql://root:password@localhost:5432/ticket-assignment?sslmode=disable" -verbose down
sqlc:
	sqlc generate
tests:
	go test -v -cover ./...
gomock:
	mockgen -package mockdb --destination db/mocks/tx.go github.com/ShadrackAdwera/ticket-assignment/db/sqlc TxStore
start:
	go run main.go

# .PHONY: migrate_up migrate_down sqlc tests gomock start
	