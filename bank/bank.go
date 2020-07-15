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
		root: newNode(),
	}
}

func (fb *FormBank) String() string {
	return fb.root.String()
}

func (fb *FormBank) AddForm(lookup string, form mzcli.Form) {
	n := fb.root
	for _, c := range strings.ToLower(lookup) {
		i := int8(c - 'a')
		if i < 0 || i >= int8(alphabetSize) {
			// TODO better reporting
			fmt.Printf("%s - %s\n", lookup, string(c))
			return
		}
		if _, ok := n.nodes[i]; !ok {
			n.nodes[i] = newNode()
		}
		n = n.nodes[i]
	}
	n.forms = append(n.forms, form)
}

func (fb *FormBank) findNode(lookup string) *node {
	n := fb.root
	for _, c := range lookup {
		i := int8(c - 'a')
		if i < 0 || i >= int8(alphabetSize) {
			return nil
		}
		if _, ok := n.nodes[i]; !ok {
			return nil
		}
		n = n.nodes[i]
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
