package bank

import (
	"bufio"
	"errors"
	"io"
	"strings"
)

// WordForm represents a unique word form.
// It includes the fully-marked form and the part of speech tag.
type WordForm struct {
	Accented string
	MorphTag string
}

// FormBank tracks word forms.
type FormBank struct {
	root node
}

func New(r io.Reader) (*FormBank, error) {
	bank := &FormBank{}
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		l := scanner.Text()
		if l == "" {
			continue
		}
		cols := strings.Split(l, "\t")
		if len(cols) != 4 {
			return nil, errors.New("could not parse input")
		}
		bank.addForm(cols[0], WordForm{
			Accented: cols[3],
			MorphTag: cols[1],
		})
	}

	return bank, nil
}

func (fb *FormBank) String() string {
	return fb.root.String()
}

func (fb *FormBank) addForm(lookup string, form WordForm) {
	n := &fb.root
	for _, c := range lookup {
		i := c - 'a'
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
		i := c - 'a'
		if n.nodes[i] == nil {
			return nil
		}
		n = n.nodes[i]
	}
	return n
}

func (fb *FormBank) Lookup(s string) []WordForm {
	n := fb.findNode(s)
	if n == nil {
		return []WordForm{}
	}

	forms := n.ExactForms()
	return forms
}

func (fb *FormBank) LookupPartial(s string) []WordForm {
	n := fb.findNode(s)
	if n == nil {
		return []WordForm{}
	}

	forms := n.Forms()
	return forms
}
