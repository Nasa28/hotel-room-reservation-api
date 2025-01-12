build: 
	@go build -o bin/hotel_room_reservation cmd/main.go

test: 
	@go test -v ./...

run:build
	@./bin/hotel_room_reservation

migration:
	@migrate create -ext sql -dir cmd/migrate/migrations $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	@go run cmd/migrate/main.go -action up

migrate-down:
	@go run cmd/migrate/main.go -action down