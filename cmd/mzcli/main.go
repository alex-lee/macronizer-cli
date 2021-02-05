package main

import (
	"bytes"
	"flag"
	"os"
	"runtime"
	"runtime/pprof"

	"github.com/c-bata/go-prompt"

	"collat.io/macronizer-cli/assets"
	"collat.io/macronizer-cli/bank"
	"collat.io/macronizer-cli/compact"
	"collat.io/macronizer-cli/query"
)

func completer(doc prompt.Document) []prompt.Suggest {
	return []prompt.Suggest{}
}

func startProfile() func() {
	cpuProf, err := os.Create("cpu.prof")
	if err != nil {
		panic(err)
	}

	if err := pprof.StartCPUProfile(cpuProf); err != nil {
		panic(err)
	}

	return func() {
		memProf, err := os.Create("mem.prof")
		if err != nil {
			panic(err)
		}

		runtime.GC()
		if err := pprof.WriteHeapProfile(memProf); err != nil {
			panic(err)
		}

		memProf.Close()

		pprof.StopCPUProfile()
		cpuProf.Close()
	}
}

func loadFormBank(profile bool) *bank.FormBank {
	if profile {
		cleanup := startProfile()
		defer cleanup()
	}

	lemmasData := bytes.NewReader(assets.LemmasData)
	morphTagsData := bytes.NewReader(assets.MorphTagsData)
	entriesData := bytes.NewReader(assets.EntriesData)

	b := bank.New()
	entriesChan, err := compact.Unpack(lemmasData, morphTagsData, entriesData)
	if err != nil {
		panic(err)
	}

	for pe := range entriesChan {
		b.AddForm(pe.Bare, pe.Form)
	}

	return b
}

func main() {
	profile := flag.Bool("profile", false, "Profile the data load.")
	flag.Parse()

	b := loadFormBank(*profile)
	h := query.NewHandler(b)
	p := prompt.New(h.Handle, completer)
	p.Run()
}
