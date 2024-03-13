package useragent

import (
	"fmt"
	"regexp"
)

// Regular expression copied from https://semver.org/.
const (
	semVerCore          = `(0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*)`
	semVerPrerelease    = `(?:-((?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?`
	semVerBuildmetadata = `(?:\+([0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?`
)

var regexpSemVer = regexp.MustCompile(`^` + semVerCore + semVerPrerelease + semVerBuildmetadata + `$`)

func isSemVer(s string) bool {
	return regexpSemVer.MatchString(s)
}

func matchSemVer(s string) error {
	if isSemVer(s) {
		return nil
	}
	return fmt.Errorf("invalid semver string: %s", s)
}

// Alphanumeric characters, hyphen, underscore, and period. This is the subset of
// characters that we allow in user agent keys and values. This is to ensure that
// downstream applications can correctly parse the full user agent header.
//
// NOTE: HTTP headers in general only work well with ASCII characters. see:
// https://stackoverflow.com/questions/4400678/what-character-encoding-should-i-use-for-a-http-header
var validChars = `0-9A-Za-z_\-\.`

func isValid(s string) bool {
	return regexp.MustCompile(`^[` + validChars + `]+$`).MatchString(s)
}

func matchValidChars(s string) error {
	if isValid(s) {
		return nil
	}
	return fmt.Errorf("invalid alphanumeric string: %s", s)
}
