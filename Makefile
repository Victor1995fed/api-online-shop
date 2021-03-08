include .env
export

.PHONY: build
build:
	go build -v ./cmd/apiserver
.PHONY: test
test:
	go test -v -race -timeout 30s ./...
.PHONY: migrate-up 
migrate-up:
	./migrate -path migrations -database "postgres://pgsql?dbname=${POSTGRES_DB}&user=${POSTGRES_USER}&password=${POSTGRES_PASSWORD}&sslmode=disable" up
.PHONY: migrate-down

migrate-down:
	./migrate -path migrations -database "postgres://pgsql?dbname=${POSTGRES_DB}&user=${POSTGRES_USER}&password=${POSTGRES_PASSWORD}&sslmode=disable" down

.PHONY: migrate-create
migrate-create:
				./migrate create -ext sql -dir migrations $(MN)

.PHONY: install
install:
	curl -L  $(MIGRATION_URL_DOWNLOAD)  \
    | tar xvz ; \
    mv migrate.linux-amd64 migrate
.DEFAULT_GOAL := build