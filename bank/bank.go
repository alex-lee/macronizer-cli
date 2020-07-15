package bank

import (
	"fmt"
	"strings"

	mzcli "collat.io/macronizer-cli"
)

// FormBank tracks word forms.
type FormBank struct {
	root *node
}

func New() *FormBank {
	return &FormBank{
		root: newNode('\x00'),
	}
}

func (fb *FormBank) String() string {
	return fb.root.String()
}

func (fb *FormBank) AddForm(lookup string, form mzcli.Form) {
	n := fb.root
	for _, c := range strings.ToLower(lookup) {
		i := byte(c - 'a')
		if i >= byte(alphabetSize) {
			// TODO better reporting
			fmt.Printf("%s - %s\n", lookup, string(c))
			return
		}
		n = n.find(i, true)
	}
	n.forms = append(n.forms, form)
}

func (fb *FormBank) findNode(lookup string) *node {
	n := fb.root
	for _, c := range strings.ToLower(lookup) {
		i := byte(c - 'a')
		if i >= byte(alphabetSize) {
			return nil
		}
		n = n.find(i, false)
		if n == nil {
			return nil
		}
	}
	return n
}

func (fb *FormBank) Find(s string) []mzcli.Form {
	n := fb.findNode(s)
	if n == nil {
		return []mzcli.Form{}
	}

	forms := n.ExactForms()
	return forms
}

func (fb *FormBank) FindPartial(s string) []mzcli.Form {
	n := fb.findNode(s)
	if n == nil {
		return []mzcli.Form{}
	}

	forms := n.Forms(10)
	return forms
}
