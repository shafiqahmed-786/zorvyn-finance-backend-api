run:
	go run cmd/server/main.go

seed:
	go run cmd/seed/main.go

test:
	go test ./...

docs:
	swag init -g cmd/server/main.go