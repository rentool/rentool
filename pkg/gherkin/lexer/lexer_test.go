package lexer

import (
	"bufio"
	"fmt"
	"strings"
	"testing"
)

func TestFeature(t *testing.T) {
	input := `Feature Test Feature
	
	Привет как дела

	`

	var aga = bufio.NewReader(strings.NewReader(input))
	for {
		line, err := aga.ReadString('\n')

		fmt.Printf(" > Read %d characters: %v\n", len(line), line)
		if err != nil {
			break
		}
	}

	lexer := New(strings.NewReader(input))

	var token = lexer.NextToken()

	t.Fatalf("Got: %v", token)

	//tests := []struct{
	//	expectedType token.TokenType
	//	expectedLiteral string
	//}{
	//	{token.FEATURE, "Feature"},
	//	{token.EOF, ""},
	//}
	//
	//for i, test := range tests {
	//	tok := lexer.NextToken()
	//
	//	if tok.Type != test.expectedType {
	//		t.Fatalf("tests[%d] - wrong token type. expected=%v, got=%v", i, test.expectedType, tok.Type)
	//	}
	//}
}
