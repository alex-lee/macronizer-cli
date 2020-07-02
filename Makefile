.PHONY: build test

build: binclude.go
	go build -o mzcli ./cmd/mzcli

test:
	go test ./...

binclude.go: assets/macrons.txt
	go generate
