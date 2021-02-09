// Pre-process the macrons data, to make it more compact.
package main

import (
	"compress/gzip"
	"io"
	"os"

	"collat.io/macronizer-cli/compact"
)

const (
	macronsFile   = "./assets/macrons.txt"
	lemmasFile    = "./assets/packed_lemmas.txt.gz"
	morphTagsFile = "./assets/packed_morphtags.txt.gz"
	entriesFile   = "./assets/packed_entries.txt.gz"
)

type dataFile struct {
	f *os.File
	w io.WriteCloser
}

func (df *dataFile) Write(p []byte) (n int, err error) {
	return df.w.Write(p)
}

func (df *dataFile) Close() error {
	if err := df.w.Close(); err != nil {
		return err
	}
	if err := df.f.Close(); err != nil {
		return err
	}
	return nil
}

func createDataFile(path string) *dataFile {
	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	w := gzip.NewWriter(f)
	return &dataFile{f, w}
}

func main() {
	src, err := os.Open(macronsFile)
	if err != nil {
		panic(err)
	}
	defer src.Close()

	dstLemmas := createDataFile(lemmasFile)
	defer dstLemmas.Close()

	dstMorphTags := createDataFile(morphTagsFile)
	defer dstMorphTags.Close()

	dstEntries := createDataFile(entriesFile)
	defer dstEntries.Close()

	err = compact.Pack(dstLemmas, dstMorphTags, dstEntries, src)
	if err != nil {
		panic(err)
	}
}
