package lexer

import (
	"bufio"
	"github.com/rentool/rentool/pkg/gherkin/matcher"
	"github.com/rentool/rentool/pkg/gherkin/token"
	"io"
	"strings"
)

// Lexer structure.
type Lexer struct {
	keywordsMatcher matcher.ReadOnlyMatcher
	reader          *bufio.Reader
	lineProcessed   bool
	currentRune     rune
	line            string
	trimmedLine     string
	lineNo          int
	isEOF           bool
	error           error
	currentToken    token.Token
}

// New creates new Lexer structure.
func New(reader io.Reader, keywordsMatcher matcher.ReadOnlyMatcher) *Lexer {
	lexer := &Lexer{
		keywordsMatcher: keywordsMatcher,
		reader:          bufio.NewReader(reader),
		lineProcessed:   false,
		lineNo:          -1,
	}

	return lexer
}

// NextToken reads next sequence of runes and returns matching token.
func (l *Lexer) NextToken() token.Token {
	result := l.scanEOF() ||
		l.scanKeyword()

	if !result {
		if nil != l.error {
			panic(l.error)
		}

		panic("Unexpected error occurred")
	}

	return l.currentToken
}

func (l *Lexer) scanEOF() bool {
	if l.isEOF {
		l.currentToken = token.Token{
			Type:    token.Eof,
			Literal: "",
		}

		return true
	}

	return false
}

func (l *Lexer) scanKeyword() bool {

	return false
}

func (l *Lexer) scanComment() bool {
	return false
}

// readLine reads next line and saves line data to lexer's state.
func (l *Lexer) readLine() error {
	l.lineNo++
	l.lineProcessed = false

	line, err := l.reader.ReadString('\n')
	if err != nil {
		l.line = ""
		l.trimmedLine = ""
		return err
	}
	l.line = line
	l.trimmedLine = strings.Trim(line, " ")

	return nil
}
