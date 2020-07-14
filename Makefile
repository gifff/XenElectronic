all: test
	go build -o out/xenelectronic-server cmd/xenelectronic-server/main.go
	./out/xenelectronic-server --port 9000

run:
	go run cmd/xenelectronic-server/main.go --port 9000

test:
	set -o pipefail && \
	go test -failfast -count 1 -timeout 30s -race -covermode=atomic -cover -coverprofile=cover.out ./... && \
	go tool cover -html=cover.out -o cover.html
