package compact

import (
	"bufio"
	"fmt"
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

func Unpack(
	lemmasData io.Reader,
	morphTagsData io.Reader,
	entriesData io.Reader,
) (<-chan PackedEntry, error) {
	lemmas, err := loadLemmas(lemmasData)
	if err != nil {
		return nil, err
	}

	morphTags, err := loadMorphTags(morphTagsData)
	if err != nil {
		return nil, err
	}

	entriesChan := make(chan PackedEntry)

	// Parse entries incrementally.
	go func() {
		var l string
		scanner := bufio.NewScanner(entriesData)

		for scanner.Scan() {
			l = scanner.Text()

			entry, err := parseFormEntry(l)
			if err != nil {
				fmt.Println(&ParseError{"invalid entry", err}) // TODO clean up
				return
			}

			form := mzcli.Form{
				Accented: entry.Accented,
				Lemma:    lemmas[entry.Lemma],
				MorphTag: morphTags[entry.MorphTag],
			}
			entriesChan <- PackedEntry{entry.Bare, form}
		}

		fmt.Println("Done loading entries.")
		close(entriesChan)
	}()

	return entriesChan, nil
}

func loadLemmas(src io.Reader) ([]string, error) {
	var lemmas []string

	scanner := bufio.NewScanner(src)

	if ok := scanner.Scan(); !ok {
		return nil, &ParseError{"empty lemmas data", nil}
	}
	l := scanner.Text()
	numLemmas, err := strconv.Atoi(l)
	if err != nil {
		return nil, &ParseError{"could not determine lemma count", err}
	}
	lemmas = make([]string, numLemmas)

	for scanner.Scan() {
		l = scanner.Text()
		index, lemma, err := parseTableEntry(l)
		if err != nil {
			return nil, &ParseError{"invalid lemma", err}
		}
		lemmas[index] = lemma
	}
	return lemmas, nil
}

func loadMorphTags(src io.Reader) ([]string, error) {
	var morphTags []string

	scanner := bufio.NewScanner(src)

	if ok := scanner.Scan(); !ok {
		return nil, &ParseError{"empty lemmas data", nil}
	}
	l := scanner.Text()
	numMorphTags, err := strconv.Atoi(l)
	if err != nil {
		return nil, &ParseError{"could not determine lemma count", err}
	}
	morphTags = make([]string, numMorphTags)

	for scanner.Scan() {
		l = scanner.Text()
		index, morphTag, err := parseTableEntry(l)
		if err != nil {
			return nil, &ParseError{"invalid morphTag", err}
		}
		morphTags[index] = morphTag
	}
	return morphTags, nil
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
