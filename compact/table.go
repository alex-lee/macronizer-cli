package compact

import (
	"bufio"
	"fmt"
)

// lookupTable maps strings to ints, and vice versa.
type lookupTable struct {
	cur     int
	toInt   map[string]int
	fromInt map[int]string
}

func newLookupTable() *lookupTable {
	return &lookupTable{
		cur:     0,
		toInt:   make(map[string]int),
		fromInt: make(map[int]string),
	}
}

func (t *lookupTable) register(s string) int {
	if i, ok := t.toInt[s]; ok {
		return i
	}

	i := t.cur
	t.cur++

	t.toInt[s] = i
	t.fromInt[i] = s
	return i
}

func (t *lookupTable) size() int {
	return len(t.toInt)
}

func (t *lookupTable) write(w *bufio.Writer) {
	for k, v := range t.fromInt {
		w.WriteString(fmt.Sprintf("%d\t%s\n", k, v))
	}
}
