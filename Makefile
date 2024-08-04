run: build
	./bin/eyhash

run-file: build
	./bin/eyhash test.txt -f

run-folder: build
	./bin/eyhash tests -d

run-unknown: build
	./bin/eyhash test.txt -u

build:
	go build -ldflags="-s -w" -o bin/eyhash cmd/main.go

test:
	go test -v ./...

test-race:
	go test -v ./... --race
