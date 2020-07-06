package mzcli

// Form represents a unique word form.
// It includes the fully-marked form and the part of speech tag.
type Form struct {
	Accented string
	Lemma    string
	MorphTag string
}
