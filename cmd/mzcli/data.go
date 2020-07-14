package main

//go:generate binclude -gzip

import (
	"compress/gzip"
	"io"

	"github.com/lu4p/binclude"
)

var (
	lemmasDataPath    = binclude.Include("../../assets/packed_lemmas.txt")
	morphTagsDataPath = binclude.Include("../../assets/packed_morphtags.txt")
	entriesDataPath   = binclude.Include("../../assets/packed_entries.txt")
)

// loadData returns readers for all the packed data.
// Data is uncompressed. Use the cleanup function to close the files when done.
func loadData() (
	lemmasData io.Reader,
	morphTagsData io.Reader,
	entriesData io.Reader,
	cleanup func(),
	err error,
) {
	lemmasData, closeLemmasData, err := loadPath(lemmasDataPath)
	if err != nil {
		return
	}

	morphTagsData, closeMorphTagsData, err := loadPath(morphTagsDataPath)
	if err != nil {
		return
	}

	entriesData, closeEntriesData, err := loadPath(entriesDataPath)
	if err != nil {
		return
	}

	cleanup = func() {
		closeLemmasData()
		closeMorphTagsData()
		closeEntriesData()
	}
	return
}

func loadPath(path string) (io.Reader, func(), error) {
	compressed, err := BinFS.Open(path)
	if err != nil {
		return nil, nil, err
	}
	f, err := gzip.NewReader(compressed)
	if err != nil {
		return nil, nil, err
	}

	return f, func() {
		f.Close()
		compressed.Close()
	}, nil
}
