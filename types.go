package mzcli

// Form represents a unique word form.
// It includes the fully-marked form and the part of speech tag.
type Form struct {
	Accented string
	Lemma    string
	MorphTag string
}

// FormFinder looks for word forms.
type FormFinder interface {
	Find(s string) []Form
	FindPartial(s string) []Form
}
