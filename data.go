package mzcli

//go:generate binclude -gzip

import (
	"io"

	"github.com/lu4p/binclude"
)

var assetPath = binclude.Include("./assets")

// MacronsData returns a reader providing the (gzipped) macrons data.
func MacronsData() (io.ReadCloser, error) {
	return BinFS.Open(assetPath + "/macrons.txt")
}
