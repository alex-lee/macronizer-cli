package bank

import (
	"fmt"
	"strings"

	mzcli "collat.io/macronizer-cli"
)

// FormBank tracks word forms.
type FormBank struct {
	root node
}

func (fb *FormBank) String() string {
	return fb.root.String()
}

func (fb *FormBank) AddForm(lookup string, form mzcli.Form) {
	n := &fb.root
	for _, c := range strings.ToLower(lookup) {
		i := int(c - 'a')
		if i < 0 || i >= alphabetSize {
			fmt.Printf("%s - %s\n", lookup, string(c))
			return
		}
		if n.nodes[i] == nil {
			n.nodes[i] = new(node)
		}
		n = n.nodes[i]
	}
	n.forms = append(n.forms, form)
}

func (fb *FormBank) findNode(lookup string) *node {
	n := &fb.root
	for _, c := range lookup {
		i := int(c - 'a')
		if i < 0 || i >= alphabetSize {
			return nil
		}
		if n.nodes[i] == nil {
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
