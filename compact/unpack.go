package compact

import (
	"bufio"
	"io"
	"strconv"
	"strings"

	mzcli "collat.io/macronizer-cli"
)

type ParseError struct {
	Message string
	Err     error
}

func (e *ParseError) Error() string {
	return "could not parse packed data: " + e.Message
}

func (e *ParseError) Unwrap() error {
	return e.Err
}

type PackedEntry struct {
	Bare string
	Form mzcli.Form
}

func Unpack(src io.Reader) ([]PackedEntry, error) {
	var packedEntries []PackedEntry
	lemmas := make(map[int]string)
	morphTags := make(map[int]string)

	scanner := bufio.NewScanner(src)

	// Parse lemmas.
	if ok := scanner.Scan(); !ok {
		return nil, &ParseError{"empty data", nil}
	}
	if l := scanner.Text(); l != packedDividerLemmas {
		return nil, &ParseError{"lemma section not found", nil}
	}
	for scanner.Scan() {
		l := scanner.Text()
		if l == packedDividerMorphTags {
			break
		}

		index, lemma, err := parseTableEntry(l)
		if err != nil {
			return nil, &ParseError{"invalid lemma", err}
		}
		lemmas[index] = lemma
	}

	// Parse morphTags.
	for scanner.Scan() {
		l := scanner.Text()
		if l == packedDividerEntries {
			break
		}

		index, morphTag, err := parseTableEntry(l)
		if err != nil {
			return nil, &ParseError{"invalid morphTag", err}
		}
		morphTags[index] = morphTag
	}

	// Parse entries.
	for scanner.Scan() {
		l := scanner.Text()

		entry, err := parseFormEntry(l)
		if err != nil {
			return nil, &ParseError{"invalid entry", err}
		}

		morphTag, ok := morphTags[entry.MorphTag]
		if !ok {
			return nil, &ParseError{"missing morph tag", nil}
		}
		lemma, ok := lemmas[entry.Lemma]
		if !ok {
			return nil, &ParseError{"missing lemma", nil}
		}

		form := mzcli.Form{
			Accented: entry.Accented,
			Lemma:    lemma,
			MorphTag: morphTag,
		}

		packedEntries = append(packedEntries, PackedEntry{entry.Bare, form})
	}

	return packedEntries, nil
}

func copyString(s string) string {
	if len(s) == 0 {
		return ""
	}
	return s[0:1] + s[1:]
}

func parseTableEntry(l string) (index int, value string, err error) {
	cols := strings.Split(l, "\t")
	if len(cols) != 2 {
		err = &ParseError{"invalid row format", nil}
		return
	}

	index, err = strconv.Atoi(cols[0])
	if err != nil {
		err = &ParseError{"invalid index", err}
		return
	}

	value = copyString(cols[1])
	return
}

func parseFormEntry(l string) (entry, error) {
	cols := strings.Split(l, "\t")
	if len(cols) != 4 {
		return entry{}, &ParseError{"invalid row format", nil}
	}

	morphTagIndex, err := strconv.Atoi(cols[1])
	if err != nil {
		return entry{}, &ParseError{"invalid morphTag index", err}
	}
	lemmaIndex, err := strconv.Atoi(cols[2])
	if err != nil {
		return entry{}, &ParseError{"invalid lemma index", err}
	}

	accented := copyString(cols[3])
	bare := copyString(cols[0])
	if bare == "" {
		bare = stripAccents(accented)
	}

	return entry{
		Bare:     bare,
		MorphTag: morphTagIndex,
		Lemma:    lemmaIndex,
		Accented: accented,
	}, nil
}
