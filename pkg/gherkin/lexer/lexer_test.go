package lexer

import (
	"fmt"
	"github.com/rentool/rentool/pkg/gherkin/matcher"
	"github.com/rentool/rentool/pkg/gherkin/token"
	"github.com/spf13/viper"
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

// TestFixtures runs all fixture files from ./fixtures directory.
func TestFixtures(t *testing.T) {
	// Given
	config := viper.New()
	config.SetConfigFile("fixtures/complete.yaml")
	err := config.ReadInConfig()
	if err != nil {
		t.Fatal(err)
	}

	var featureContent strings.Builder
	var expectedTokens []string
	for line, rawTokens := range config.AllSettings() {
		featureContent.WriteString(line)

		tokensData, ok := rawTokens.([]interface{})
		if !ok {
			t.Fatal("cannot map tokensData to []interface{}")
		}

		for _, rawTok := range tokensData {
			tokenData, ok := rawTok.(map[interface{}]interface{})
			if !ok {
				t.Fatal("cannot map token to map[string]string")
			}

			typ, ok := tokenData["type"].(string)
			if !ok {
				t.Fatal("cannot convert token type to string")
			}

			literal, ok := tokenData["literal"].(string)
			if !ok {
				t.Fatal("cannot convert token literal to string")
			}

			expectedTokens = append(expectedTokens, fmt.Sprintf("%v: %v", typ, literal))
		}
	}

	// When
	lexer := createLexer(featureContent.String())
	var actualTokens []string
	for t := lexer.NextToken(); t.Type != token.Eof; t = lexer.NextToken() {
		actualTokens = append(actualTokens, fmt.Sprintf("%v: %v", t.Type.String(), t.Literal))
	}

	// Then
	assert.Equal(t, expectedTokens, actualTokens)
}

func Test_givenEmptyLine_whenNextToken_thenReturnEOF(t *testing.T) {
	// Given
	lexer := createLexer("")

	// When
	var tokens []token.Token
	for t := lexer.NextToken(); t.Type != token.Eof; t = lexer.NextToken() {
		tokens = append(tokens, t)
	}

	// Then
	assert.Equal(t, []token.Token{
		{
			Type:    token.Text,
			Literal: "",
		},
		{
			Type:    token.Eos,
			Literal: "",
		},
	}, tokens)
}

// createLexer creates Lexer service with given feature content and default keywords matcher.
func createLexer(input string) *Lexer {
	var keywordsMatcher = matcher.NewRuneTrieMatcher()
	for name, value := range englishKeywords {
		keywordsMatcher.Put(name, value)
	}

	return New(strings.NewReader(input), keywordsMatcher)
}

func createToken(typ token.Type, literal string) token.Token {
	return token.Token{
		Type:    typ,
		Literal: literal,
	}
}

func createTokenType(str string) token.Type {
	switch str {
	case token.Text.String():
		return token.Text
	case token.Feature.String():
		return token.Feature
	case token.Eof.String():
		return token.Eof
	case token.Eos.String():
		return token.Eos
	}

	panic(fmt.Sprintf("Cannot create token.Type from string: %v", str))
}
