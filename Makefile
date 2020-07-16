# Makefile for macronizer-cli

binary := mzcli
binclude_source := cmd/mzcli/binclude.go
macrons_source := assets/macrons.txt
packed_sources := assets/packed_lemmas.txt assets/packed_morphtags.txt assets/packed_entries.txt

.PHONY: build clean test binclude

# Build rules

build: binclude
	go build -o $(binary) ./cmd/mzcli

clean:
	-rm $(packed_sources)
	-rm $(binclude_source)
	-rm $(binary)

test:
	go test ./...

binclude: $(binclude_source)

$(binclude_source): $(packed_sources)
	go generate ./cmd/mzcli

$(packed_sources): $(macrons_source)
	go run ./cmd/dataprep $(packed_sources)

# Release rules

dist-test:
	goreleaser --snapshot --skip-publish --rm-dist
