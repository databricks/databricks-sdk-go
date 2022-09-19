package code

import "strings"

var reservedWords = []string{
	"break", "default", "func", "interface", "select", "case", "defer", "go",
	"map", "struct", "chan", "else", "goto", "switch", "const", "fallthrough",
	"if", "range", "type", "continue", "for", "import", "return", "var",
	"append", "bool", "byte", "iota", "len", "make", "new",
}

type Named struct {
	Name        string
	Description string
}

func (n *Named) IsNameReserved() bool {
	for _, v := range reservedWords {
		if n.CamelName() == v {
			return true
		}
	}
	return false
}

// TODO: Add tests, document accepted inputs/outputs.
func (n *Named) PascalName() string {
	return strings.ReplaceAll(
		strings.Title(
			strings.ReplaceAll(
				strings.ReplaceAll(
					strings.ReplaceAll(n.Name, "$", ""),
					"-", " ",
				),
				"_", " ")),
		" ", "")
}

func (n *Named) CamelName() string {
	cc := n.PascalName()
	return strings.ToLower(cc[0:1]) + cc[1:]
}

func (n *Named) HasComment() bool {
	return n.Description != ""
}

func (n *Named) Comment(prefix string, maxLen int) string {
	if n.Description == "" {
		return ""
	}
	var lines []string
	currentLine := strings.Builder{}
	for _, v := range whitespace.Split(n.Description, -1) {
		if len(prefix)+currentLine.Len()+len(v)+1 > maxLen {
			lines = append(lines, currentLine.String())
			currentLine.Reset()
		}
		if currentLine.Len() > 0 {
			currentLine.WriteRune(' ')
		}
		currentLine.WriteString(v)
	}
	if currentLine.Len() > 0 {
		lines = append(lines, currentLine.String())
		currentLine.Reset()
	}
	return strings.TrimLeft(prefix, "\t ") + strings.Join(lines, "\n"+prefix)
}
