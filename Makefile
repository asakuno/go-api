# Import .env file
ifneq (,$(wildcard ./.env))
		include .env
		export $(shell sed 's/=.*//' .env)
endif

# Variables
CONTAINER_NAME=${APP_NAME}-app
MYSQL_CONTAINER_NAME=${APP_NAME}-db

# Commands
init:
	docker compose up -d --build

up:
	docker compose up -d

down:
	docker compose down

logs:
	docker compose logs -f

log-app:
	docker compose logs -f app

# MySQL commands
container-mysql:
	docker exec -it ${MYSQL_CONTAINER_NAME} /bin/sh

create-db:
	docker exec -it ${MYSQL_CONTAINER_NAME} /bin/sh -c "mysql -u${DB_USER} -p${DB_PASS} -e 'CREATE DATABASE IF NOT EXISTS ${DB_NAME}'"

# Docker commands
container-go:
	docker exec -it ${CONTAINER_NAME} /bin/sh

migrate:
	docker exec -it ${CONTAINER_NAME} /bin/sh -c "go run main.go --migrate"

seed:
	docker exec -it ${CONTAINER_NAME} /bin/sh -c "go run main.go --seed"

migrate-seed:
	docker exec -it ${CONTAINER_NAME} /bin/sh -c "go run main.go --migrate --seed"

go-tidy:
	docker exec -it ${CONTAINER_NAME} /bin/sh -c "go mod tidy"

ps:
	docker compose ps
