# Makefile for macronizer-cli

binary := mzcli
macrons_source := assets/macrons.txt
data_files := lemmas morphtags entries
packed_sources := $(patsubst %,assets/packed_%.txt.gz,$(data_files))

# Build rules

.PHONY: build
build: packed-data
	go build -o $(binary) ./cmd/mzcli

.PHONY: packed-data
packed-data: $(packed_sources)

$(packed_sources): $(macrons_source)
	go run ./cmd/dataprep

.PHONY: clean
clean:
	-rm $(packed_sources)
	-rm $(binary)

.PHONY: test
test:
	go test ./...

# Release rules

.PHONY: dist-test
dist-test:
	goreleaser --snapshot --skip-publish --rm-dist
