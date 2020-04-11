package lexer

import (
	"github.com/rentool/rentool/pkg/gherkin/matcher"
	"github.com/rentool/rentool/pkg/gherkin/token"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

var englishKeywords = map[string]token.Type{
	"Feature":    token.Feature,
	"Background": token.Background,
	"Scenario":   token.Scenario,
	"Given":      token.Given,
	"When":       token.When,
	"Then":       token.Then,
	"And":        token.And,
	"But":        token.Background,
	"Examples":   token.Examples,
}

func Test_givenEmptyLine_whenNextToken_thenReturnEOF(t *testing.T) {
	// Given
	lexer := createLexer("")

	// When
	result := lexer.NextToken()

	// Then
	assert.Equal(t, token.Eof, result.Type)
}

// createLexer creates Lexer service with given feature content and default keywords matcher.
func createLexer(input string) *Lexer {
	var keywordsMatcher = matcher.NewRuneTrieMatcher()
	for name, value := range englishKeywords {
		keywordsMatcher.Put(name, value)
	}

	return New(strings.NewReader(input), keywordsMatcher)
}
