package main

//go:generate binclude -gzip

import (
	"io"

	"github.com/lu4p/binclude"
)

var macronsDataPath = binclude.Include("../../assets/macrons_packed.txt")

// macronsData returns a reader providing the (gzipped) macrons data.
func macronsData() (io.ReadCloser, error) {
	return BinFS.Open(macronsDataPath)
}
