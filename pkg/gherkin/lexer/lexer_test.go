package lexer

import (
	"github.com/rentool/rentool/pkg/gherkin/matcher"
	"github.com/rentool/rentool/pkg/gherkin/token"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestFixtures(t *testing.T) {
	// Arrange
	keywordsMatcher := matcher.NewRuneTrieMatcher()
	keywordsMatcher.Put("Feature", token.Feature)
	keywordsMatcher.Put("Background", token.Background)
	keywordsMatcher.Put("Scenario", token.Scenario)
	keywordsMatcher.Put("Scenario Outline", token.Outline)
	keywordsMatcher.Put("Given", token.Given)
	keywordsMatcher.Put("When", token.When)
	keywordsMatcher.Put("Then", token.Then)
	keywordsMatcher.Put("And", token.And)
	keywordsMatcher.Put("But", token.But)
	keywordsMatcher.Put("Examples", token.Examples)

	filePath := "../fixtures/complete.feature"
	file, err := os.Open(filePath)
	if err != nil {
		t.Fatalf("failed to open file %v: %v", filePath, err)
	}

	var actualTokens strings.Builder
	expectedTokensFilePath := filePath + "-tokens.txt"
	expectedTokensData, err := ioutil.ReadFile(expectedTokensFilePath)
	if err != nil {
		t.Fatalf("failed to read file %v: %v", expectedTokensFilePath, err)
	}
	expectedTokens := string(expectedTokensData)

	// Act
	lexer := New(file, keywordsMatcher)
	for t := lexer.NextToken(); t.Type != token.Eof; t = lexer.NextToken() {
		actualTokens.WriteString(t.Type.String() + ":")
		if "" != t.Keyword {
			actualTokens.WriteString(" " + t.Keyword)
		}
		if t.Type == token.TableRow {
			actualTokens.WriteString(" | ")
			for i, column := range t.Values {
				actualTokens.WriteString(column)

				if i != len(t.Values)-1 {
					actualTokens.WriteString(" |")
				}
			}
			actualTokens.WriteString(" |")
		} else if "" != t.Value {
			actualTokens.WriteString(" " + t.Value)
		}
		actualTokens.WriteString("\n")
	}

	// Assert
	assert.Equal(t, expectedTokens, actualTokens.String())
}
