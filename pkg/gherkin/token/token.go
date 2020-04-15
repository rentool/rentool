//go:generate stringer -type=Type

package token

import (
	"errors"
	"fmt"
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
	Eos
)

var TypeNotFound = errors.New("TokenNotFound")

// stringToTokenType looks for token type for given string. For example, for string "Scenario" it
// returns token.Scenario type. If token type not found it returns Illegal type and error.
func StringToTokenType(str string) (Type, error) {
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
