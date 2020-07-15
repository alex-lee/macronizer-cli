package compact

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strings"
)

// entry is a single row in the macrons file.
type entry struct {
	Bare     string
	MorphTag int
	Lemma    int
	Accented string
}

// payload is the structure that gets serialized and saved.
type payload struct {
	Entries   []entry
	MorphTags map[int]string
	Lemmas    map[int]string
}

func Pack(dstLemmas io.Writer, dstMorphTags io.Writer, dstEntries io.Writer, src io.Reader) error {
	morphTags := newLookupTable()
	lemmas := newLookupTable()
	var entries []entry

	// Parse entries and build the lemma and morphTag tables.
	scanner := bufio.NewScanner(src)
	for scanner.Scan() {
		l := scanner.Text()
		if l == "" || strings.HasPrefix(l, "#") {
			continue
		}

		cols := strings.Split(l, "\t")
		if len(cols) != 4 {
			return errors.New("could not parse input")
		}

		// See if the bare form can be derived from the accented form.
		// If so, just put an empty string for the bare form.
		bare := cols[0]
		normalized := strings.ToLower(stripAccents(cols[3]))
		if normalized == bare {
			bare = ""
		}

		// Save references to morphology tag and lemma.
		morphTagIndex := morphTags.register(cols[1])
		lemmaIndex := lemmas.register(cols[2])

		e := entry{
			Bare:     bare,
			MorphTag: morphTagIndex,
			Lemma:    lemmaIndex,
			Accented: cols[3],
		}
		entries = append(entries, e)
	}

	// Write the tables
	lemmas.write(bufio.NewWriter(dstLemmas))
	morphTags.write(bufio.NewWriter(dstMorphTags))

	// Write the entries.
	w := bufio.NewWriter(dstEntries)
	for _, e := range entries {
		line := fmt.Sprintf("%s\t%d\t%d\t%s\n", e.Bare, e.MorphTag, e.Lemma, e.Accented)
		w.WriteString(line)
	}
	w.Flush()

	fmt.Printf("Recorded %d lemmas\n", lemmas.size())
	fmt.Printf("Recorded %d morph tags\n", morphTags.size())
	fmt.Printf("Recorded %d entries\n", len(entries))

	return nil
}
