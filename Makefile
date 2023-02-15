.PHONY:
.SILENT:
.DEFAULT_GOAL := run

build:
	go mod download && go build -o ./.bin/app ./cmd/server/main.go

run: build
	 ./.bin/app

test:
	go test -coverprofile=coverage.out ./...
	make test.coverage

test.coverage:
	go tool cover -func=cover.out | grep "total"