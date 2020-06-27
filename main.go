package main

import (
	"fmt"

	"github.com/c-bata/go-prompt"
)

func executor(input string) {
	fmt.Println(input)
}

func completer(doc prompt.Document) []prompt.Suggest {
	return []prompt.Suggest{}
}

func main() {
	p := prompt.New(executor, completer)
	p.Run()
}
