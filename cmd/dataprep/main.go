// Pre-process the macrons data, to make it more compact.
package main

import (
	"os"

	"collat.io/macronizer-cli/compact"
)

const (
	macronsFile   = "./assets/macrons.txt"
	lemmasFile    = "./assets/packed_lemmas.txt"
	morphTagsFile = "./assets/packed_morphtags.txt"
	entriesFile   = "./assets/packed_entries.txt"
)

func main() {
	src, err := os.Open(macronsFile)
	if err != nil {
		panic(err)
	}
	defer src.Close()

	dstLemmas, err := os.Create(lemmasFile)
	if err != nil {
		panic(err)
	}
	defer dstLemmas.Close()

	dstMorphTags, err := os.Create(morphTagsFile)
	if err != nil {
		panic(err)
	}
	defer dstMorphTags.Close()

	dstEntries, err := os.Create(entriesFile)
	if err != nil {
		panic(err)
	}
	defer dstEntries.Close()

	err = compact.Pack(dstLemmas, dstMorphTags, dstEntries, src)
	if err != nil {
		panic(err)
	}
}
