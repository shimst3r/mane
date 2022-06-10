// Package mane implements currency locale formatting logic.
package mane

import (
	"fmt"
	"regexp"
)

// Locale represents the locale of a given curreny string.
type Locale string

var (
	rEU = regexp.MustCompile(`(\d),(\d{2})`)
)

// Convert convert a curreny string from one locale to another.
func Convert(input string, source, destination Locale) (string, error) {
	return "", nil
}

func fromEU(input string) (int64, error) {
	matches := rEU.FindStringSubmatch(input)
	if len(matches) == 0 {
		return 0, fmt.Errorf("error: no matches for input %s", input)
	}
	return 0, nil
}
