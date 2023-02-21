D_URL=postgresql://root:password@localhost:5432/ticket-assignment?sslmode=disable
migrate_up:
	migrate -path db/migrations -database "${D_URL}" -verbose up
migrate_down:
	migrate -path db/migrations -database "${D_URL}" -verbose down
migrate_up_one:
	migrate -path db/migrations -database "${D_URL}" -verbose up 1
migrate_down_one:
	migrate -path db/migrations -database "${D_URL}" -verbose down 1
sqlc:
	sqlc generate
tests:
	go test -v -cover ./...
gomock:
	mockgen -package mockdb --destination db/mocks/tx.go github.com/ShadrackAdwera/ticket-assignment/db/sqlc TxStore
start:
	go run main.go

.PHONY: migrate_up migrate_down sqlc tests gomock start
	