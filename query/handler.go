package query

import (
	"fmt"
	"strings"

	mzcli "collat.io/macronizer-cli"
)

// FormFinder looks for word forms.
type FormFinder interface {
	Find(s string) []mzcli.Form
	FindPartial(s string) []mzcli.Form
}

// Handler interprets input and looks up form matches.
type Handler struct {
	finder FormFinder
}

// NewHandler returns a new handler based on the given form finder.
func NewHandler(finder FormFinder) *Handler {
	return &Handler{
		finder: finder,
	}
}

// Handle processes an input string and finds matching forms.
func (h *Handler) Handle(input string) {
	var forms []mzcli.Form

	s, exact := analyze(input)

	if exact {
		forms = h.finder.Find(s)
	} else {
		forms = h.finder.FindPartial(s)
	}

	// Figure out how much padding we need for clean alignment.
	padding := maxAccentedLength(forms) + 1
	format := fmt.Sprintf("%%-%ds %%s\n", padding)

	// Print out the results.
	for _, m := range forms {
		fmt.Printf(format, prettify(m.Accented), m.MorphTag)
	}
}

func analyze(input string) (s string, exact bool) {
	if strings.HasSuffix(input, " ") {
		exact = true
	}
	s = strings.TrimSpace(input)
	return
}

func maxAccentedLength(forms []mzcli.Form) int {
	max := 0
	for _, f := range forms {
		l := len(f.Accented)
		if l > max {
			max = l
		}
	}
	return max
}
