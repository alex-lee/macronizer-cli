binclude_data := cmd/mzcli/binclude.go
macrons_source := assets/macrons.txt
packed_data := assets/packed_lemmas.txt assets/packed_morphtags.txt assets/packed_entries.txt
binary := mzcli

.PHONY: build clean test binclude

build: binclude
	go build -o $(binary) ./cmd/mzcli

clean:
	-rm $(packed_data)
	-rm $(binclude_data)
	-rm $(binary)

test:
	go test ./...

binclude: $(binclude_data)

$(binclude_data): $(packed_data)
	go generate ./cmd/mzcli

$(packed_data): $(macrons_source)
	go run ./cmd/dataprep $(packed_data)
