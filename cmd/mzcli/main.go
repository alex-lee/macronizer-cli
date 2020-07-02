package main

import (
	"compress/gzip"

	"github.com/c-bata/go-prompt"

	mzcli "collat.io/macronizer-cli"
	"collat.io/macronizer-cli/bank"
	"collat.io/macronizer-cli/query"
)

func completer(doc prompt.Document) []prompt.Suggest {
	return []prompt.Suggest{}
}

func loadFormBank() *bank.FormBank {
	f, err := mzcli.MacronsData()
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
	h := query.NewHandler(b)
	p := prompt.New(h.Handle, completer)
	p.Run()
}
