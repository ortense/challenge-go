build:
	go build -o ./bin/challange ./cmd/launch/main.go

test:
	go test -v ./...

run:
	go run cmd/launch/main.go

run-inmemory:
	go run cmd/memory/main.go

start: build
	./bin/challange