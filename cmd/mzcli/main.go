package main

import (
	"compress/gzip"
	"flag"
	"os"
	"runtime"
	"runtime/pprof"

	"github.com/c-bata/go-prompt"

	"collat.io/macronizer-cli/bank"
	"collat.io/macronizer-cli/compact"
	"collat.io/macronizer-cli/query"
)

func completer(doc prompt.Document) []prompt.Suggest {
	return []prompt.Suggest{}
}

func loadFormBank(profile bool) *bank.FormBank {
	if profile {
		cpuProf, err := os.Create("cpu.prof")
		if err != nil {
			panic(err)
		}
		defer cpuProf.Close()

		if err := pprof.StartCPUProfile(cpuProf); err != nil {
			panic(err)
		}
		defer pprof.StopCPUProfile()
	}

	f, err := macronsData()
	if err != nil {
		panic(err)
	}
	defer f.Close()

	r, err := gzip.NewReader(f)
	if err != nil {
		panic(err)
	}
	defer r.Close()

	b := &bank.FormBank{}
	packedEntries, err := compact.Unpack(r)
	if err != nil {
		panic(err)
	}

	for _, pe := range packedEntries {
		b.AddForm(pe.Bare, pe.Form)
	}

	if profile {
		memProf, err := os.Create("mem.prof")
		if err != nil {
			panic(err)
		}
		defer memProf.Close()

		runtime.GC()
		if err := pprof.WriteHeapProfile(memProf); err != nil {
			panic(err)
		}
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
