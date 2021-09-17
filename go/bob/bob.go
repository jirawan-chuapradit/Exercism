// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package main

import (
	"strings"
	"unicode"
)

// Remark is a convenience type for determining the type of content made by the conversational partner.
type Remark struct {
	content string
}

func (r Remark) isSilent() bool {
	return r.content == ""
}

func (r Remark) isQuestion() bool {
	return strings.HasSuffix(r.content, "?")
}

func (r Remark) isShouted() bool {
	containsLetters := strings.IndexFunc(r.content, unicode.IsLetter) > -1
	isUpperCase := strings.ToUpper(r.content) == r.content
	return containsLetters && isUpperCase
}

func (r Remark) isAnnoyed() bool {
	return r.isShouted() && r.isQuestion()
}

func newRemark(remark string) Remark {
	trimmedRemark := strings.TrimSpace(remark)
	return Remark{trimmedRemark}
}

// Hey parses a content made by the conversational partner and returns Bob's response.
func Hey(remark string) string {
	var r = newRemark(remark)

	switch {
	case r.isSilent():
		return "Fine. Be that way!"
	case r.isAnnoyed():
		return "Calm down, I know what I'm doing!"
	case r.isQuestion():
		return "Sure."
	case r.isShouted():
		return "Whoa, chill out!"
	default:
		return "Whatever."
	}
}