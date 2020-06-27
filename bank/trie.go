package bank

import (
	"fmt"
	"strings"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"
const alphabetSize = len(alphabet)

type node struct {
	nodes [alphabetSize]*node
	forms []WordForm
}

func (n *node) String() string {
	var lines []string

	for _, f := range n.forms {
		row := fmt.Sprintf("%s\t%s", f.Accented, f.MorphTag)
		lines = append(lines, row)
	}
	for _, child := range n.nodes {
		if child != nil {
			lines = append(lines, child.String())
		}
	}

	return strings.Join(lines, "\n")
}

func (n *node) ExactForms() []WordForm {
	return n.forms
}

func (n *node) Forms() []WordForm {
	forms := make([]WordForm, len(n.forms))
	copy(forms, n.forms)

	for _, child := range n.nodes {
		if child != nil {
			forms = append(forms, child.Forms()...)
		}
	}

	return forms
}
