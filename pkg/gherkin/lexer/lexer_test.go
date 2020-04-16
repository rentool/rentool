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
	var expectedContent strings.Builder
	var actualContent strings.Builder

	featureRawData := config.Get("feature")
	featureData, ok := featureRawData.(map[string]interface{})
	if !ok {
		t.Fatalf("cannot map feature data to map[string]interface{}: %+v", featureRawData)
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

			expectedContent.WriteString(fmt.Sprintf("%v: %v\n", typ, literal))
		}
	}

	rawKeywords := config.Get("keywords")
	keywords, ok := rawKeywords.(map[string]interface{})
	if !ok {
		t.Fatalf("cannot map keywords to map[string]interface{}: %+v", rawKeywords)
	}

	var keywordsMatcher = matcher.NewRuneTrieMatcher()
	for keyword, typeRawStr := range keywords {
		typeStr, ok := typeRawStr.(string)
		if !ok {
			t.Fatalf("cannot map token type to string: %+v", typeRawStr)
		}

		tokType, err := token.StringToTokenType(typeStr)
		if err != nil {
			t.Fatalf("cannot create token type from string %+v: %v", typeStr, err)
		}

		keywordsMatcher.Put(keyword, tokType)
	}

	// When
	lexer := New(strings.NewReader(featureContent.String()), keywordsMatcher)
	for t := lexer.NextToken(); t.Type != token.Eof; t = lexer.NextToken() {
		actualContent.WriteString(fmt.Sprintf("%v: %v\n", t.Type.String(), t.Literal))
	}

	// Then
	assert.Equal(t, expectedContent.String(), actualContent.String())
}
