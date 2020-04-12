package matcher

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRuneTrieMatcher_Get(t *testing.T) {
	var matcher = NewRuneTrieMatcher()
	matcher.Put("key1", "value1")
	matcher.Put("key2", "value2")

	assert.Equal(t, "value1", matcher.Get("key1"))
	assert.Equal(t, "value2", matcher.Get("key2"))
	assert.Equal(t, nil, matcher.Get("invalid"))
}

func TestRuneTrieMatcher_GetByText(t *testing.T) {
	var matcher = NewRuneTrieMatcher()
	matcher.Put("Scenario", 0)
	matcher.Put("Scenario:", 1)

	t.Run("must return value and rest text", func(t *testing.T) {
		val, rest := matcher.GetByText("Scenario: some text")
		assert.Equal(t, 1, val)
		assert.Equal(t, " some text", rest)
	})
}

func TestRuneTrieMatcher_FindFirstAtStartOfText(t *testing.T) {
	var matcher = NewRuneTrieMatcher()
	matcher.Put("SingleWord", 0)
	matcher.Put("Two words", 1)
	matcher.Put("caseInsenSEtivE", 2)

	t.Run("must match single word", func(t *testing.T) {
		val, offset := matcher.FindFirstAtStartOfText("SingleWord some other text")
		assert.Equal(t, 0, val)
		assert.Equal(t, 10, offset)
	})

	t.Run("must match two words", func(t *testing.T) {
		val, offset := matcher.FindFirstAtStartOfText("Two words some other text")
		assert.Equal(t, 1, val)
		assert.Equal(t, 9, offset)
	})

	t.Run("must be case insensitive", func(t *testing.T) {
		val, offset := matcher.FindFirstAtStartOfText("CASEinsensetive")
		assert.Equal(t, 2, val)
		assert.Equal(t, 15, offset)
	})
}
