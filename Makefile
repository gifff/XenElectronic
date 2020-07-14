all: test build run

build:
	go build -o out/xenelectronic-server cmd/xenelectronic-server/main.go

run:
	./out/xenelectronic-server --port 9000

test:
	set -o pipefail && \
	go test -failfast -count 1 -timeout 30s -race -covermode=atomic -cover -coverprofile=cover.out ./... && \
	go tool cover -html=cover.out -o cover.html
