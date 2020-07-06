package bank_test

import (
	"bufio"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"

	mzcli "collat.io/macronizer-cli"
	"collat.io/macronizer-cli/bank"
	"collat.io/macronizer-cli/compact"
)

const macrons = `
a	e--------	a	a_
a	r--------	ab	a_
ab	r--------	ab	ab
abs	r--------	ab	abs
aps	r--------	ab	aps
ac	c--------	atque	ac
ad	r--------	ad	ad
addo	v1spia---	addo	addo_
advenio	v1spia---	advenio	adve^nio_
`

func testForms(t *testing.T) []compact.PackedEntry {
	var entries []compact.PackedEntry

	scanner := bufio.NewScanner(strings.NewReader(macrons))
	for scanner.Scan() {
		l := scanner.Text()
		if l == "" {
			continue
		}
		cols := strings.Split(l, "\t")
		if len(cols) != 4 {
			t.Fatalf("macrons test data is invalid")
		}
		entries = append(entries, compact.PackedEntry{
			Bare: cols[0],
			Form: mzcli.Form{
				Accented: cols[3],
				Lemma:    cols[2],
				MorphTag: cols[1],
			},
		})
	}

	return entries
}

func testFormBank(t *testing.T) bank.FormBank {
	b := bank.FormBank{}
	for _, pe := range testForms(t) {
		b.AddForm(pe.Bare, pe.Form)
	}
	return b
}

func extractAccented(forms []mzcli.Form) []string {
	accented := make([]string, 0, len(forms))
	for _, f := range forms {
		accented = append(accented, f.Accented)
	}
	return accented
}

func TestLookup(t *testing.T) {
	fb := testFormBank(t)

	tests := []struct {
		query    string
		expected []string
	}{
		{"a", []string{"a_", "a_"}},
		{"ab", []string{"ab"}},
	}

	for _, test := range tests {
		forms := fb.Find(test.query)
		accented := extractAccented(forms)

		if diff := cmp.Diff(test.expected, accented); diff != "" {
			t.Errorf("Wrong lookup results:\n%s", diff)
		}
	}
}

func TestLookupPartial(t *testing.T) {
	fb := testFormBank(t)

	tests := []struct {
		query    string
		expected []string
	}{
		{"ab", []string{"ab", "abs"}},
		{"ad", []string{"ad", "addo_", "adve^nio_"}},
	}

	for _, test := range tests {
		forms := fb.FindPartial(test.query)
		accented := extractAccented(forms)

		if diff := cmp.Diff(test.expected, accented); diff != "" {
			t.Errorf("Wrong partial lookup results:\n%s", diff)
		}
	}
}
