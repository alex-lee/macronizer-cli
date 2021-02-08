# Makefile for macronizer-cli

binary := mzcli
macrons_source := assets/macrons.txt
data_files := lemmas morphtags entries
packed_sources := $(patsubst %,assets/packed_%.txt.gz,$(data_files))

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
	go run ./cmd/dataprep

# Release rules

dist-test:
	goreleaser --snapshot --skip-publish --rm-dist
