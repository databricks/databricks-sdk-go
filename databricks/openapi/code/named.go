package code

import (
	"strings"
	"unicode"
)

var reservedWords = []string{
	"break", "default", "func", "interface", "select", "case", "defer", "go",
	"map", "struct", "chan", "else", "goto", "switch", "const", "fallthrough",
	"if", "range", "type", "continue", "for", "import", "return", "var",
	"append", "bool", "byte", "iota", "len", "make", "new", "package",
}

// Named holds common methods for identifying and describing things
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

// emulate positive lookaheads from JVM regex:
// (?<=[a-z])(?=[A-Z])|(?<=[A-Z])(?=[A-Z][a-z])|([-_\s])
// and convert all words to lower case
func (n *Named) splitASCII() (w []string) {
	var current []rune
	nameLen := len(n.Name)
	var last, this, next, lookahead bool
	// we do assume here that all named entities are strictly ASCII
	for i := 0; i < nameLen; i++ {
		r := rune(n.Name[i])
		if r == '$' {
			// we're naming language literals, $ is usually not allowed
			continue
		}
		this = unicode.IsUpper(r)
		r = unicode.ToLower(r)
		next = false
		lookahead = i+1 < nameLen
		if lookahead {
			next = unicode.IsUpper(rune(n.Name[i+1]))
		}
		split, before, after := false, false, true
		if last && this && !next && lookahead {
			// (?<=[A-Z])(?=[A-Z][a-z])
			split = true
			before = false
			after = true
		}
		if !this && next {
			// (?<=[a-z])(?=[A-Z])
			split = true
			before = true
			after = false
		}
		if r == '-' || r == '_' || r == ' ' {
			// ([-_\s])
			split = true
			before = false
			after = false
		}
		if before {
			current = append(current, r)
		}
		if split {
			w = append(w, string(current))
			current = []rune{}
		}
		if after {
			current = append(current, r)
		}
		last = this
	}
	if len(current) > 0 {
		w = append(w, string(current))
	}
	return w
}

// PascalName creates NamesLikesThis
func (n *Named) PascalName() string {
	var sb strings.Builder
	for _, w := range n.splitASCII() {
		sb.WriteString(strings.Title(w))
	}
	return sb.String()
}

// CamelName creates namesLikesThis
func (n *Named) CamelName() string {
	cc := n.PascalName()
	return strings.ToLower(cc[0:1]) + cc[1:]
}

// SnakeName creates names_like_this
func (n *Named) SnakeName() string {
	return strings.Join(n.splitASCII(), "_")
}

// ConstantName creates NAMES_LIKE_THIS
func (n *Named) ConstantName() string {
	return strings.ToUpper(n.SnakeName())
}

// KebabName creates names-like-this
func (n *Named) KebabName() string {
	return strings.Join(n.splitASCII(), "-")
}

func (n *Named) HasComment() bool {
	return n.Description != ""
}

// Comment formats description into language-specific comment multi-line strings
func (n *Named) Comment(prefix string, maxLen int) string {
	if n.Description == "" {
		return ""
	}
	trimmed := strings.TrimSpace(n.Description)
	description := strings.ReplaceAll(trimmed, "\n\n", " __BLANK__ ")
	var lines []string
	currentLine := strings.Builder{}
	for _, v := range whitespace.Split(description, -1) {
		if v == "__BLANK__" {
			lines = append(lines, currentLine.String())
			lines = append(lines, "")
			currentLine.Reset()
			continue
		}
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
