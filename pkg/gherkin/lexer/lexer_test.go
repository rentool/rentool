package lexer

import (
	"bufio"
	"bytes"
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
		} else if t.Type == token.Tag {
			for _, tag := range t.Values {
				actualTokens.WriteString(" @" + tag)
			}
		} else if "" != t.Value {
			actualTokens.WriteString(" " + t.Value)
		}
		actualTokens.WriteString("\n")
	}

	// Assert
	assert.Equal(t, expectedTokens, actualTokens.String())
}

func TestLexer_readLine(t *testing.T) {
	t.Run("success readline", func(t *testing.T) {
		var buffer bytes.Buffer
		line := "Hello \n"
		buffer.Write([]byte(line + " world"))

		lexer := &Lexer{
			keywordsMatcher: matcher.NewRuneTrieMatcher(),
			reader:          bufio.NewReader(&buffer),
			lineNo:          -1,
		}

		lexer.readLine()

		if line != lexer.line {
			t.Fatal("lexer.line should be equals \"Hello\\n\" ")
		}
		if "Hello" != lexer.trimmedLine {
			t.Fatal("lexer.line should be equals \"Hello\"")
		}
	})

	t.Run("io.EOF error", func(t *testing.T) {
		data := []string{"", "Hello World!", "Hello \t World!"}

		for _, datum := range data {
			var buffer bytes.Buffer
			buffer.Write([]byte(datum))

			lexer := &Lexer{
				keywordsMatcher: matcher.NewRuneTrieMatcher(),
				reader:          bufio.NewReader(&buffer),
				lineNo:          -1,
			}

			lexer.readLine()

			if false == lexer.isEOF {
				t.Fatal("lexer.isEOF should be true")
			}
		}
	})
}

func TestLexer_scanEOF(t *testing.T) {
	t.Run("scanEOF returned true", func(t *testing.T) {
		lexer := &Lexer{
			isEOF: true,
		}

		if false == lexer.scanEOF() {
			t.Fatal("lexer.scanEOF should return true")
		}
	})
	t.Run("scanEOF returned true", func(t *testing.T) {
		lexer := &Lexer{}

		if true == lexer.scanEOF() {
			t.Fatal("lexer.scanEOF should return false")
		}
	})
}

func TestLexer_scanDocString(t *testing.T) {
	t.Run("line has not slash", func(t *testing.T) {
		lexer := &Lexer{}

		if true == lexer.scanDocString() {
			t.Fatal("lexer.scanDocString should return false")
		}
	})

	t.Run("line has slash", func(t *testing.T) {
		var buffer bytes.Buffer
		buffer.Write([]byte("Hello \n world"))

		lexer := &Lexer{
			line:            "test \"\"\"",
			keywordsMatcher: matcher.NewRuneTrieMatcher(),
			reader:          bufio.NewReader(&buffer),
			lineNo:          -1,
		}

		if false == lexer.scanDocString() {
			t.Fatal("lexer.scanDocString should return true")
		}
	})
}

func TestLexer_scanDocStringContent(t *testing.T) {
	t.Run("line inDocString is false", func(t *testing.T) {
		lexer := &Lexer{}

		if true == lexer.scanDocStringContent() {
			t.Fatal("lexer.scanDocString should return false")
		}
	})

	t.Run("line inDocString is true", func(t *testing.T) {
		var buffer bytes.Buffer
		buffer.Write([]byte("Hello \n world"))

		lexer := &Lexer{
			inDocString: true,
			reader:      bufio.NewReader(&buffer),
		}

		if false == lexer.scanDocStringContent() {
			t.Fatal("lexer.scanDocString should return true")
		}
	})
}

func TestLexer_scanKeyword(t *testing.T) {
	t.Run("Panic: not found token.Type", func(t *testing.T) {
		var buffer bytes.Buffer
		line := "Hello \n"
		buffer.Write([]byte(line + " world"))

		matcherObj := matcher.NewRuneTrieMatcher()
		matcherObj.Put("Hello", "world")

		lexer := &Lexer{
			keywordsMatcher: matcherObj,
			reader:          bufio.NewReader(&buffer),
			lineNo:          -1,
			trimmedLine:     "Hello",
		}

		assert.PanicsWithValue(t,
			"Keywords matcher returned invalid value: world. "+
				"It must be instance of token.Type",
			func() {
				lexer.scanKeyword()
			},
		)
	})

	t.Run("scanKeyword should return true", func(t *testing.T) {
		var buffer bytes.Buffer
		line := "Hello \n"
		buffer.Write([]byte(line + " world"))

		matcherObj := matcher.NewRuneTrieMatcher()
		_type, _ := token.StringToType("Scenario")
		matcherObj.Put("Scenario:", _type)

		lexer := &Lexer{
			keywordsMatcher: matcherObj,
			reader:          bufio.NewReader(&buffer),
			lineNo:          -1,
			trimmedLine:     "Scenario:",
		}

		if false == lexer.scanKeyword() {
			t.Fatal("lexer.scanKeyword should return true")
		}
	})

	t.Run("scanKeyword should return false", func(t *testing.T) {

		lexer := &Lexer{
			keywordsMatcher: matcher.NewRuneTrieMatcher(),
			lineNo:          -1,
		}

		if true == lexer.scanKeyword() {
			t.Fatal("lexer.scanKeyword should return false")
		}
	})
}

func TestLexer_scanTag(t *testing.T) {
	t.Run("trimmedLine == 0", func(t *testing.T) {
		lexer := &Lexer{}

		if true == lexer.scanTag() {
			t.Fatal("scanTag should return false")
		}
	})
	t.Run("trimmedLine[0] != '@'", func(t *testing.T) {
		lexer := &Lexer{
			trimmedLine: "",
		}

		if true == lexer.scanTag() {
			t.Fatal("scanTag should return false")
		}
	})

	t.Run("trimmedLine[0] != '@'", func(t *testing.T) {
		var buffer bytes.Buffer
		line := "Hello \n"
		buffer.Write([]byte(line + " world"))

		lexer := &Lexer{
			trimmedLine: "@",
			reader:      bufio.NewReader(&buffer),
		}

		if false == lexer.scanTag() {
			t.Fatal("scanTag should return true")
		}
	})
}
