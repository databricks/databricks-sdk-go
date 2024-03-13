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

var regexpAlphanum = regexp.MustCompile(`^[0-9A-Za-z_-]+$`)

func matchSemVer(s string) error {
	if isSemVer(s) {
		return nil
	}
	return fmt.Errorf("invalid semver string: %s", s)
}

func matchAlphanum(s string) error {
	if isAlphanum(s) {
		return nil
	}
	return fmt.Errorf("invalid alphanumeric string: %s", s)
}

var allowedValueChars = `0-9A-Za-z_\-\.`

func isValidValue(s string) bool {
	return regexp.MustCompile(`^[` + allowedValueChars + `]+$`).MatchString(s)
}

func isAlphanum(s string) bool {
	return regexpAlphanum.MatchString(s)
}

func isSemVer(s string) bool {
	return regexpSemVer.MatchString(s)
}
