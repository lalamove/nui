package nstrings

import (
	"io"
	"strings"
)

var (
	_ Replacer = (*strings.Replacer)(nil)
	// ReplacerToUpper is a replacer which replaces lower cases by upper cases
	ReplacerToUpper = ReplacerFunc(func(s string) string {
		return strings.ToUpper(s)
	})
	// ReplacerToLower is a replacer which replaces upper cases to lower cases
	ReplacerToLower = ReplacerFunc(func(s string) string {
		return strings.ToLower(s)
	})
	// ReplacerTrimSpace is a replacer which removes spaces
	ReplacerTrimSpace = ReplacerFunc(func(s string) string {
		return strings.Trim(s, " ")
	})
)

// Replacer is an interface to characters of a string
type Replacer interface {
	Replace(s string) string
	WriteString(w io.Writer, s string) (n int, err error)
}

// ReplacerFunc is a function which implements the Replacer interface
type ReplacerFunc func(s string) string

// Replace runs f
func (f ReplacerFunc) Replace(s string) string {
	return f(s)
}

// WriteString writes to w the string s with replacement applied by f
func (f ReplacerFunc) WriteString(w io.Writer, s string) (n int, err error) {
	return w.Write([]byte(f(s)))
}
