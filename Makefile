macrons_source := assets/macrons.txt
macrons_data := assets/macrons_packed.txt
binclude_data := cmd/mzcli/binclude.go

.PHONY: build test

build: $(binclude_data)
	go build -o mzcli ./cmd/mzcli

clean:
	-rm $(macrons_data) $(binclude_data)

test:
	go test ./...

$(macrons_data): $(macrons_source)
	go run ./cmd/dataprep $(macrons_data)

$(binclude_data): $(macrons_data)
	go generate ./cmd/mzcli
