package bank

import (
	"fmt"
	"strings"

	mzcli "collat.io/macronizer-cli"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"
const alphabetSize = len(alphabet)

type node struct {
	children []*node
	forms    []mzcli.Form
	char     byte
}

func newNode(char byte) *node {
	return &node{char: char}
}

func (n *node) find(c byte, create bool) *node {
	for _, child := range n.children {
		if child.char == c {
			return child
		}
	}
	if create {
		child := newNode(c)
		n.children = append(n.children, child)
		return child
	}
	return nil
}

func (n *node) String() string {
	var lines []string

	for _, f := range n.forms {
		row := fmt.Sprintf("%s\t%s", f.Accented, f.MorphTag)
		lines = append(lines, row)
	}
	for _, child := range n.children {
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

	for _, child := range n.children {
		childForms := child.Forms(limit - len(forms))
		forms = append(forms, childForms...)
	}

	return forms
}
