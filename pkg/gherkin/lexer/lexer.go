package lexer

import (
	"bufio"
	"github.com/rentool/rentool/pkg/gherkin/matcher"
	"github.com/rentool/rentool/pkg/gherkin/token"
	"io"
)

// Lexer structure.
type Lexer struct {
	keywordsMatcher matcher.ReadOnlyMatcher
	reader          *bufio.Reader
	position        int
	readPosition    int
	rune            rune
	line            int
	linePos         int
}

// New creates new Lexer structure.
func New(reader io.Reader, keywordsMatcher matcher.ReadOnlyMatcher) *Lexer {
	lexer := &Lexer{
		keywordsMatcher: keywordsMatcher,
		reader:          bufio.NewReader(reader),
	}

	return lexer
}

// NextToken reads next sequence of runes and returns matching token.
func (l *Lexer) NextToken() token.Token {

	return token.Token{
		Type:    token.ILLEGAL,
		Literal: "",
	}
	//
	//var tok token.Token
	//
	//switch l.rune {
	//case rune(0):
	//	tok.Type = token.EOF
	//default:
	//	if unicode.IsLetter(l.rune) {
	//		literal := l.readLiteral()
	//		foundToken := token.LookupKeyword(literal)
	//		_ = foundToken
	//
	//		return tok
	//	} else {
	//		tok = newToken(token.ILLEGAL, l.rune)
	//	}
	//}
	//
	//l.readRune()
	//return tok
}

//
//// readLiteral continuously reads runes until it reaches non-letter rune.
//func (l *Lexer) readLiteral() string {
//	var sb strings.Builder
//	for unicode.IsLetter(l.rune) {
//		sb.WriteRune(l.rune)
//		l.readRune()
//	}
//	return sb.String()
//}
//
//// readRune reads next character as a rune and sets it to the current state. Also updates current line and cursor
//// positions.
//func (l *Lexer) readRune() {
//	if rn, _, err := l.reader.ReadRune(); err != nil {
//		if err == io.EOF {
//			l.rune = rune(0)
//		} else {
//			fmt.Printf("ReadRune() error: %v", err)
//			os.Exit(1)
//		}
//	} else {
//		l.rune = rn
//		if l.rune == '\n' {
//			l.line++
//			l.linePos = 0
//		}
//		l.linePos++
//	}
//}
//
//// newToken creates new token.
//func newToken(tokenType token.Type, rn rune) token.Token {
//	return token.Token{Type: tokenType, Literal: string(rn)}
//}
