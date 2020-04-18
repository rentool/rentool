package gherkin

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
	inDocString     bool
	docStringIdent  int
	error           error
	tokens          []token.Token
}

// NewLexer creates new Lexer structure.
func NewLexer(reader io.Reader, keywordsMatcher matcher.ReadOnlyMatcher) *Lexer {
	lexer := &Lexer{
		keywordsMatcher: keywordsMatcher,
		reader:          bufio.NewReader(reader),
		lineNo:          -1,
	}

	lexer.readLine()

	return lexer
}

func (l *Lexer) PredictToken() token.Token {
	if len(l.tokens) == 0 {
		l.tokens = append(l.tokens, l.NextToken())
	}

	return l.tokens[0]
}

// NextToken reads next sequence of runes and returns matching token.
func (l *Lexer) NextToken() token.Token {
	if len(l.tokens) > 0 {
		return l.dequeueToken()
	}

	result := l.scanEOF() ||
		l.scanDocString() ||
		l.scanDocStringContent() ||
		l.scanKeyword() ||
		l.scanTag() ||
		l.scanComment() ||
		l.scanTableRow() ||
		l.scanText()

	if !result {
		if nil != l.error {
			panic(l.error)
		}

		panic(fmt.Errorf("unexpected error occurred for line: %v", l.line))
	}

	return l.dequeueToken()
}

func (l *Lexer) scanTag() bool {
	if len(l.trimmedLine) == 0 || l.trimmedLine[0] != '@' {
		return false
	}

	var tags []string
	for _, tag := range strings.Split(l.trimmedLine[1:], "@") {
		tags = append(tags, strings.Trim(tag, " "))
	}

	l.tokens = append(l.tokens, token.Token{
		Type:   token.Tag,
		Values: tags,
	})
	l.readLine()

	return true
}

func (l *Lexer) scanTableRow() bool {
	if len(l.trimmedLine) == 0 || l.trimmedLine[0] != '|' || l.trimmedLine[len(l.trimmedLine)-1] != '|' {
		return false
	}

	row := l.trimmedLine[1 : len(l.trimmedLine)-1]
	var columns []string
	for _, column := range strings.Split(row, "|") {
		columns = append(columns, strings.Trim(column, " "))
	}

	l.tokens = append(l.tokens, token.Token{
		Type:   token.TableRow,
		Values: columns,
	})
	l.readLine()

	return true
}

func (l *Lexer) scanDocString() bool {
	i := strings.Index(l.line, "\"\"\"")
	if -1 == i {
		return false
	}

	l.inDocString = !l.inDocString
	l.docStringIdent = i
	l.queueToken(token.DocString, "", "")
	l.readLine()

	return true
}

func (l *Lexer) scanDocStringContent() bool {
	if !l.inDocString {
		return false
	}

	l.queueToken(token.Text, "", strings.Repeat(" ", l.docStringIdent)+strings.Trim(l.line, " "))
	l.readLine()

	return true
}

func (l *Lexer) scanEOF() bool {
	if l.isEOF {
		l.queueToken(token.Eof, "", "")

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
	l.queueToken(typ, typ.String(), strings.Trim(restLine, " "))

	l.consumeLine()

	return true
}

func (l *Lexer) scanComment() bool {
	if len(l.trimmedLine) > 0 && l.trimmedLine[0] == '#' {
		l.queueToken(token.Comment, "", strings.Trim(l.line, " "))
		l.consumeLine()

		return true
	}

	return false
}

func (l *Lexer) scanText() bool {
	l.queueToken(token.Text, "", strings.Trim(l.line, " "))

	l.consumeLine()

	return true
}

func (l *Lexer) consumeLine() {
	l.readLine()
}

func (l *Lexer) dequeueToken() token.Token {
	t := l.tokens[0]
	l.tokens = l.tokens[1:]

	return t
}

func (l *Lexer) queueToken(typ token.Type, keyword string, value string) {
	l.tokens = append(l.tokens, token.Token{
		Type:    typ,
		Keyword: strings.Trim(strings.TrimRight(keyword, "\n"), " "),
		Value:   strings.TrimRight(value, "\n"),
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
	l.trimmedLine = strings.Trim(strings.Trim(line, "\n"), " ")
}
