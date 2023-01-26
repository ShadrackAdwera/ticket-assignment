migrate_up:
	migrate -path db/migrations -database "postgresql://root:password@localhost:5432/ticket-assignment?sslmode=disable" -verbose up
migrate_down:
	migrate -path db/migrations -database "postgresql://root:password@localhost:5432/ticket-assignment?sslmode=disable" -verbose down
sqlc:
	sqlc generate
tests:
	go test -v -cover ./...
	