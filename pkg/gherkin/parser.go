package gherkin

import (
	"github.com/rentool/rentool/pkg/gherkin/matcher"
	"github.com/rentool/rentool/pkg/gherkin/token"
	"log"
	"os"
)

type Parser struct {
	lexer *Lexer
}

func NewParser(filePath string) *Parser {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("failed to open file %v: %v", filePath, err)
	}

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

	return &Parser{
		lexer: NewLexer(file, keywordsMatcher),
	}
}

func (p *Parser) Parse() (*Feature, error) {
	var feature Feature

	return &feature, nil
}

func (p *Parser) parseExpression() {

}
