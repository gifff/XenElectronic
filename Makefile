all: cover build run

build:
	go build -o out/xenelectronic-server cmd/xenelectronic-server/main.go

run:
	./out/xenelectronic-server --port 9000

cover: test
	go tool cover -html=cover.out -o cover.html

test:
	set -o pipefail && \
	go test -failfast -count 1 -timeout 30s -race -covermode=atomic \
	-cover -coverprofile=cover.out ./... | tee /dev/stderr | tr '%' ' ' | \
	awk '{ if ($$5 ~ /[\\d\.]+/) { total += $$5; count++ } } END { print "coverage: " total/count "%" }'

run-frontend:
	API_BASE_URL=http://localhost:9000 npm start --prefix web

setup:
	cp etc/git-pre-push .git/hooks/pre-push
	go mod tidy -v
	npm install --prefix web
