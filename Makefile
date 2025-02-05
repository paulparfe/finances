.PHONY: run stop migrate build test

run:
	docker-compose up -d --build
	sleep 5
	docker-compose exec app goose -dir migrations postgres "postgres://user:password@db:5432/finances?sslmode=disable" up
	echo "Server running on http://localhost:8080"

stop:
	docker-compose down

migrate:
	docker-compose exec app goose -dir migrations postgres "postgres://user:password@db:5432/finances?sslmode=disable" up

build:
	docker-compose build

test:
    go test -v ./...
