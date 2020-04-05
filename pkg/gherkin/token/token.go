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
	ILLEGAL Type = iota
	EOF
	FEATURE
	BACKGROUND
	SCENARIO
	GIVEN
	WHEN
	THEN
	AND
	BUT
	EXAMPLES
	COMMENT
	TEXT
)

// keywords are words or sentences from which every logical step should start from.
var keywords = map[string]Type{
	"feature":    FEATURE,
	"background": BACKGROUND,
	"scenario":   SCENARIO,
	"given":      GIVEN,
	"when":       WHEN,
	"then":       THEN,
	"and":        AND,
	"but":        BUT,
	"examples":   EXAMPLES,
}

// LookupKeyword returns corresponding to the literal token.
func LookupKeyword(literal string) Type {
	if tok, ok := keywords[strings.ToLower(literal)]; ok {
		return tok
	}
	return ILLEGAL
}
