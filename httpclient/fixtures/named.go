package fixtures

import (
	"strings"
	"unicode"
)

// Named holds common methods for identifying and describing things
type Named struct {
	Name        string
	Description string
}

// Return the value of cond evaluated at the nearest letter to index i in name.
// dir determines the direction of search: if true, search forwards, if false,
// search backwards.
func (n *Named) search(name string, cond func(rune) bool, dir bool, i int) bool {
	nameLen := len(name)
	incr := 1
	if !dir {
		incr = -1
	}
	for j := i; j >= 0 && j < nameLen; j += incr {
		if unicode.IsLetter(rune(name[j])) {
			return cond(rune(name[j]))
		}
	}
	return false
}

// Return the value of cond evaluated on the rune at index i in name. If that
// rune is not a letter, search in both directions for the nearest letter and
// return the result of cond on those letters.
func (n *Named) checkCondAtNearestLetters(name string, cond func(rune) bool, i int) bool {
	r := rune(name[i])

	if unicode.IsLetter(r) {
		return cond(r)
	}
	return n.search(name, cond, true, i) && n.search(name, cond, false, i)
}

// emulate positive lookaheads from JVM regex:
// (?<=[a-z])(?=[A-Z])|(?<=[A-Z])(?=[A-Z][a-z])|([-_\s])
// and convert all words to lower case

func (n *Named) splitASCII() (w []string) {
	var current []rune
	var name = n.Name
	nameLen := len(name)
	var isPrevUpper, isCurrentUpper, isNextLower, isNextUpper, isNotLastChar bool
	// we do assume here that all named entities are strictly ASCII
	for i := 0; i < nameLen; i++ {
		r := rune(name[i])
		if r == '$' {
			// we're naming language literals, $ is usually not allowed
			continue
		}
		// if the current rune is a digit, check the neighboring runes to
		// determine whether to treat this one as upper-case.
		isCurrentUpper = n.checkCondAtNearestLetters(name, unicode.IsUpper, i)
		r = unicode.ToLower(r)
		isNextLower = false
		isNextUpper = false
		isNotLastChar = i+1 < nameLen
		if isNotLastChar {
			isNextLower = n.checkCondAtNearestLetters(name, unicode.IsLower, i+1)
			isNextUpper = n.checkCondAtNearestLetters(name, unicode.IsUpper, i+1)
		}
		split, before, after := false, false, true
		// At the end of a string of capital letters (e.g. HTML[P]arser).
		if isPrevUpper && isCurrentUpper && isNextLower && isNotLastChar {
			// (?<=[A-Z])(?=[A-Z][a-z])
			split = true
			before = false
			after = true
		}
		// At the end of a camel- or pascal-case word (e.g. htm[l]Parser).
		if !isCurrentUpper && isNextUpper {
			// (?<=[a-z])(?=[A-Z])
			split = true
			before = true
			after = false
		}
		if !unicode.IsLetter(r) && !unicode.IsNumber(r) {
			// ([-_\s])
			split = true
			before = false
			after = false
		}
		if before {
			current = append(current, r)
		}
		if split && len(current) > 0 {
			w = append(w, string(current))
			current = []rune{}
		}
		if after {
			current = append(current, r)
		}
		isPrevUpper = isCurrentUpper
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
