package main

//go:generate binclude -gzip

import (
	"compress/gzip"
	"fmt"
	"strings"

	"github.com/c-bata/go-prompt"
	"github.com/lu4p/binclude"

	"collat.io/macronizer-cli/bank"
)

var assetPath = binclude.Include("./assets")

func executor(b *bank.FormBank) func(string) {
	return func(input string) {
		fmt.Println(input)
		var matches []bank.WordForm
		if strings.HasSuffix(input, " ") {
			matches = b.Lookup(strings.TrimSpace(input))
		} else {
			matches = b.LookupPartial(strings.TrimSpace(input))
		}

		for _, m := range matches {
			fmt.Printf("%s (%s)\n", m.Accented, m.MorphTag)
		}
	}
}

func completer(doc prompt.Document) []prompt.Suggest {
	return []prompt.Suggest{}
}

func loadFormBank() *bank.FormBank {
	f, err := BinFS.Open(assetPath + "/macrons.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	r, err := gzip.NewReader(f)
	if err != nil {
		panic(err)
	}
	defer r.Close()

	b, err := bank.New(r)
	if err != nil {
		panic(err)
	}

	return b
}

func main() {
	b := loadFormBank()
	p := prompt.New(executor(b), completer)
	p.Run()
}
