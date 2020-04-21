//go:generate stringer -type=Type

package token

import (
	"errors"
	"fmt"
)

// Type custom type of uint8
type Type uint8

// Token
type Token struct {
	Type    Type
	Value   string
	Keyword string
	Values  []string
}

// Make initialize new Token object
func Make(typ Type, value string, keyword string) Token {
	return Token{
		Type:    typ,
		Value:   value,
		Keyword: keyword,
	}
}

// String string representation of Token
func (t Token) String() string {
	return fmt.Sprintf("Token(%v, %v)", t.Type, t.Keyword)
}

// Listing Type const
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
	Eos
	DocString
	TableRow
	Tag
)

// TypeNotFound type not found error
var TypeNotFound = errors.New("TokenNotFound")

// StringToType looks for token type for given string. For example, for string "Scenario" it
// returns token.Scenario type. If token type not found it returns Illegal type and error.
func StringToType(str string) (Type, error) {
	switch str {
	case Illegal.String():
		return Illegal, nil
	case Eof.String():
		return Eof, nil
	case Feature.String():
		return Feature, nil
	case Background.String():
		return Background, nil
	case Scenario.String():
		return Scenario, nil
	case Outline.String():
		return Outline, nil
	case Given.String():
		return Given, nil
	case When.String():
		return When, nil
	case Then.String():
		return Then, nil
	case And.String():
		return And, nil
	case But.String():
		return But, nil
	case Examples.String():
		return Examples, nil
	case Comment.String():
		return Comment, nil
	case Text.String():
		return Text, nil
	case Eos.String():
		return Eos, nil
	}

	return Illegal, TypeNotFound
}
