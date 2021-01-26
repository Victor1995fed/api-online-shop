.PHONY: build
build:
		go build -v ./cmd/apiserver
.PHONY: test
test:
	go test -v -race -timeout 30s ./...
.PHONY: install
install:
	curl -L https://github.com/golang-migrate/migrate/releases/download/v4.14.1/migrate.linux-amd64.tar.gz  \
    | tar xvz ; \
    mv migrate.linux-amd64 migrate
.DEFAULT_GOAL := build