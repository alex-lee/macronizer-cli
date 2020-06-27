package bank_test

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"

	"collat.io/macronizer-cli/bank"
)

const macrons = `
a	e--------	a	a_
aaron	n-s---mn-	Aaron	A^a^ro_n
a	r--------	ab	a_
ab	r--------	ab	ab
abs	r--------	ab	abs
aps	r--------	ab	aps
aba	n-s---mb-	Aba	Aba_
aba	n-s---mn-	Aba	aba
aba	n-s---mv-	Aba	aba
abas	n-p---ma-	Aba	Aba_s
`

func extractAccented(forms []bank.WordForm) []string {
	accented := make([]string, 0, len(forms))
	for _, f := range forms {
		accented = append(accented, f.Accented)
	}
	return accented
}

func TestNew(t *testing.T) {
	_, err := bank.New(strings.NewReader(macrons))
	if err != nil {
		t.Errorf("failed to create: %v", err)
	}
}

func TestLookup(t *testing.T) {
	fb, _ := bank.New(strings.NewReader(macrons))

	tests := []struct {
		query    string
		expected []string
	}{
		{"aaron", []string{"A^a^ro_n"}},
		{"ab", []string{"ab"}},
	}

	for _, test := range tests {
		forms := fb.Lookup(test.query)
		accented := extractAccented(forms)

		if diff := cmp.Diff(test.expected, accented); diff != "" {
			t.Errorf("Wrong lookup results:\n%s", diff)
		}
	}
}

func TestLookupPartial(t *testing.T) {
	fb, _ := bank.New(strings.NewReader(macrons))

	results := fb.LookupPartial("ab")
	if diff := cmp.Diff([]string{"ab"}, results); diff != "" {
		t.Errorf("Lookup failure:\n%s", diff)
	}

	t.Log(fb.String())
}
