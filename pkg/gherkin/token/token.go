//go:generate stringer -type=Type

package token

import (
	"fmt"
	"strings"
)

type Type uint8

type Token struct {
	Type    Type
	Literal string
}

func (t Token) String() string {
	return fmt.Sprintf("Token(%v, %v)", t.Type, t.Literal)
}

const (
	Illegal Type = iota
	Eof
	Feature
	Background
	Scenario
	Outline
	Given
	When
	Then
	And
	But
	Examples
	Comment
	Text
	KeywordValue
)

// keywords are words or sentences from which every statement should start from.
var keywords = map[string]Type{
	"feature":          Feature,
	"background":       Background,
	"scenario":         Scenario,
	"scenario outline": Outline,
	"given":            Given,
	"when":             When,
	"then":             Then,
	"and":              And,
	"but":              But,
	"examples":         Examples,
}

// LookupKeyword returns corresponding to the literal token.
func LookupKeyword(literal string) Type {
	if tok, ok := keywords[strings.ToLower(literal)]; ok {
		return tok
	}
	return Illegal
}
