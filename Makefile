# Makefile for macronizer-cli

binary := mzcli
macrons_source := assets/macrons.txt
packed_sources := assets/packed_lemmas.txt assets/packed_morphtags.txt assets/packed_entries.txt

# Build rules

.PHONY: build
build: $(packed_sources)
	go build -o $(binary) ./cmd/mzcli

.PHONY: clean
clean:
	-rm $(packed_sources)
	-rm $(binary)

.PHONY: test
test:
	go test ./...

$(packed_sources): $(macrons_source)
	go run ./cmd/dataprep $(packed_sources)

# Release rules

dist-test:
	goreleaser --snapshot --skip-publish --rm-dist
