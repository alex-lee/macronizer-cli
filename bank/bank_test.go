package bank_test

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"

	"collat.io/macronizer-cli/bank"
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
		{"a", []string{"a_", "a_"}},
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

	tests := []struct {
		query    string
		expected []string
	}{
		{"ab", []string{"ab", "abs"}},
		{"ad", []string{"ad", "addo_", "adve^nio_"}},
	}

	for _, test := range tests {
		forms := fb.LookupPartial(test.query)
		accented := extractAccented(forms)

		if diff := cmp.Diff(test.expected, accented); diff != "" {
			t.Errorf("Wrong partial lookup results:\n%s", diff)
		}
	}
}
