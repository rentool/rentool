package lexer

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/rentool/rentool/pkg/gherkin/matcher"
	"github.com/rentool/rentool/pkg/gherkin/token"
	"io"
	"strings"
)

// Lexer structure.
type Lexer struct {
	keywordsMatcher matcher.ReadOnlyMatcher
	reader          *bufio.Reader
	currentRune     rune
	line            string
	trimmedLine     string
	lineNo          int
	isEOF           bool
	error           error
	tokens          []token.Token
}

// New creates new Lexer structure.
func New(reader io.Reader, keywordsMatcher matcher.ReadOnlyMatcher) *Lexer {
	lexer := &Lexer{
		keywordsMatcher: keywordsMatcher,
		reader:          bufio.NewReader(reader),
		lineNo:          -1,
	}

	lexer.readLine()

	return lexer
}

// NextToken reads next sequence of runes and returns matching token.
func (l *Lexer) NextToken() token.Token {
	if len(l.tokens) > 0 {
		return l.dequeueToken()
	}

	result := l.scanEOF() ||
		l.scanKeyword() ||
		l.scanComment() ||
		l.scanText()

	if !result {
		if nil != l.error {
			panic(l.error)
		}

		panic(fmt.Errorf("unexpected error occurred for line: %v", l.line))
	}

	return l.dequeueToken()
}

func (l *Lexer) scanEOF() bool {
	if l.isEOF {
		l.queueToken(token.Eof, "")

		return true
	}

	return false
}

func (l *Lexer) scanKeyword() bool {
	tokenValue, restLine := l.keywordsMatcher.GetByText(l.trimmedLine)
	if nil == tokenValue {
		return false
	}

	typ, ok := tokenValue.(token.Type)
	if !ok {
		panic(fmt.Sprintf("Keywords matcher returned invalid value: %+v. It must be instance of token.Type", tokenValue))
	}
	l.queueToken(typ, typ.String())

	l.queueToken(token.Text, restLine)

	l.consumeLine()

	return true
}

func (l *Lexer) scanComment() bool {
	return false
}

func (l *Lexer) scanText() bool {
	l.queueToken(token.Text, l.line)

	l.consumeLine()

	return true
}

func (l *Lexer) consumeLine() {
	l.queueToken(token.Eos, "")

	l.readLine()
}

func (l *Lexer) dequeueToken() token.Token {
	t := l.tokens[0]
	l.tokens = l.tokens[1:]

	return t
}

func (l *Lexer) queueToken(typ token.Type, literal string) {
	l.tokens = append(l.tokens, token.Token{
		Type:    typ,
		Literal: strings.Trim(strings.Trim(literal, "\n"), " "),
	})
}

// readLine reads next line and saves line data to lexer's state.
func (l *Lexer) readLine() {
	l.lineNo++

	line, err := l.reader.ReadString('\n')
	if err != nil {
		if errors.Is(err, io.EOF) {
			l.isEOF = true
		} else {
			l.line = ""
			l.trimmedLine = ""
			l.error = err
		}

		return
	}

	l.line = line
	l.trimmedLine = strings.Trim(line, " ")
}
