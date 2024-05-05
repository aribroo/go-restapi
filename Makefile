build: 
	@go build -o bin/golang-ecommerce cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/golang-ecommerce

migration: 
	@migrate create -ext sql -dir cmd/migrate/migrations $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	@go run cmd/migrate/main.go up

migrate-dow:
	@go run cmd/migrate/main.go downf