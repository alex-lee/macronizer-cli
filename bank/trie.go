package bank

import (
	"fmt"
	"strings"

	mzcli "collat.io/macronizer-cli"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"
const alphabetSize = len(alphabet)

type node struct {
	nodes map[int8]*node
	forms []mzcli.Form
}

func newNode() *node {
	return &node{
		nodes: make(map[int8]*node),
	}
}

func (n *node) String() string {
	var lines []string

	for _, f := range n.forms {
		row := fmt.Sprintf("%s\t%s", f.Accented, f.MorphTag)
		lines = append(lines, row)
	}
	for _, child := range n.nodes {
		lines = append(lines, child.String())
	}

	return strings.Join(lines, "\n")
}

func (n *node) ExactForms() []mzcli.Form {
	return n.forms
}

func (n *node) Forms(limit int) []mzcli.Form {
	forms := make([]mzcli.Form, len(n.forms))
	copy(forms, n.forms)
	if len(forms) > limit {
		return forms
	}

	for _, child := range n.nodes {
		childForms := child.Forms(limit - len(forms))
		forms = append(forms, childForms...)
	}

	return forms
}
