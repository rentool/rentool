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

	featureRawData := config.Get("feature")
	featureData, ok := featureRawData.(map[string]interface{})
	if !ok {
		t.Fatalf("Cannot map feature data to map[string]interface{}: %+v", featureRawData)
	}

	for line, rawTokens := range featureData {
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

	rawKeywords := config.Get("keywords")
	keywords, ok := rawKeywords.(map[string]interface{})
	if !ok {
		t.Fatalf("Cannot map keywords to map[string]interface{}: %+v", rawKeywords)
	}

	for keyword, tokType := range keywords {

	}

	var keywordsMatcher = matcher.NewRuneTrieMatcher()
	for name, value := range englishKeywords {
		keywordsMatcher.Put(name, value)
	}

	return New(strings.NewReader(input), keywordsMatcher)

	// When
	lexer := createLexer(featureContent.String())
	var actualTokens []string
	for t := lexer.NextToken(); t.Type != token.Eof; t = lexer.NextToken() {
		actualTokens = append(actualTokens, fmt.Sprintf("%v: %v", t.Type.String(), t.Literal))
	}

	// Then
	assert.Equal(t, expectedTokens, actualTokens)
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
