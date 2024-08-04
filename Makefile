run: build
	./bin/main test.txt

build:
	go build -ldflags="-s -w" -o bin/main cmd/main.go

test:
	go test -v ./...

test-race:
	go test -v ./... --race
