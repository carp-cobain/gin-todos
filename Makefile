.PHONY: all
all: fmt test build

.PHONY: fmt
fmt:
	@gofmt -w ./**/*.go

.PHONY: build
build:
	go build

.PHONY: test
test:
	go test -v ./database/repo

.PHONY: clean
clean:
	@go clean

.PHONY: run
run:
	@go run main.go

# Example litestream commands below
# Assumes local database name is 'todos.db'

.PHONY: replicate
replicate:
	@litestream replicate todos.db s3://gin-todos.localhost:9000/db

.PHONY: restore
restore:
	@rm -f todos.db todos.db-shm todos.db-wal
	@litestream restore -o todos.db s3://gin-todos.localhost:9000/db
