build:
	docker-compose build

up:
	docker-compose up

down:
	docker-compose down

rebuild:
	docker-compose down
	docker-compose build
	docker-compose up

test:
	go test ./...

doc:
	swag init -g cmd/main.go -o swagger