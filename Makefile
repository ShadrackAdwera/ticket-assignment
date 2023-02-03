migrate_up:
	migrate -path db/migrations -database "postgresql://adwera:mdcclxxvi@localhost:5432/ticket-assignment?sslmode=disable" -verbose up
migrate_down:
	migrate -path db/migrations -database "postgresql://adwera:mdcclxxvi@localhost:5432/ticket-assignment?sslmode=disable" -verbose down
migrate_up_one:
	migrate -path db/migrations -database "postgresql://adwera:mdcclxxvi@localhost:5432/ticket-assignment?sslmode=disable" -verbose up 1
migrate_down_one:
	migrate -path db/migrations -database "postgresql://adwera:mdcclxxvi@localhost:5432/ticket-assignment?sslmode=disable" -verbose down 1
sqlc:
	sqlc generate
tests:
	go test -v -cover ./...
gomock:
	mockgen -package mockdb --destination db/mocks/tx.go github.com/ShadrackAdwera/ticket-assignment/db/sqlc TxStore
start:
	go run main.go

.PHONY: migrate_up migrate_down sqlc tests gomock start
	