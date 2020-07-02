.PHONY: build test

build:
	go build -o mzcli .

test:
	go test ./...
