package code

import "strings"

type Named struct {
	Name        string
	Description string
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