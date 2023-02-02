package useragent

import (
	"fmt"
	"regexp"
)

var regexpSemVer = regexp.MustCompile(`^\d+\.\d+\.\d+$`)
var regexpAlphanum = regexp.MustCompile(`^[0-9A-Za-z_-]+$`)

func matchSemVer(s string) error {
	if regexpSemVer.MatchString(s) {
		return nil
	}
	return fmt.Errorf("invalid semver string: %s", s)
}

func matchAlphanum(s string) error {
	if regexpAlphanum.MatchString(s) {
		return nil
	}
	return fmt.Errorf("invalid alphanumeric string: %s", s)
}

func matchAlphanumOrSemVer(s string) error {
	if regexpAlphanum.MatchString(s) || regexpSemVer.MatchString(s) {
		return nil
	}
	return fmt.Errorf("invalid alphanumeric or semver string: %s", s)
}
