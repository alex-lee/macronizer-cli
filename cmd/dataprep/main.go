// Pre-process the macrons data, to make it more compact.
package main

import (
	"os"

	"collat.io/macronizer-cli/compact"
)

func main() {
	src, err := os.Open("./assets/macrons.txt")
	if err != nil {
		panic(err)
	}
	defer src.Close()

	dst, err := os.Create(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer dst.Close()

	err = compact.Pack(dst, src)
	if err != nil {
		panic(err)
	}
}
