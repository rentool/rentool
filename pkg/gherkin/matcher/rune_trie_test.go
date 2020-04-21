package matcher

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRuneTrieMatcher_Get(t *testing.T) {
	var matcher = NewRuneTrieMatcher()
	matcher.Put("key1", "value1")
	matcher.Put("key2", "value2")
	matcher.Put(" key3 ", "value3")

	assert.Equal(t, "value1", matcher.Get("key1"))
	assert.Equal(t, "value2", matcher.Get("key2"))
	assert.Equal(t, "value3", matcher.Get(" key3 "))
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

	t.Run("must return nil value and full text", func(t *testing.T) {
		val, rest := matcher.GetByText("ololol")
		assert.Nil(t, val)
		assert.Equal(t, "ololol", rest)
	})

	t.Run("must return value and empty string", func(t *testing.T) {
		val, rest := matcher.GetByText("")
		assert.Nil(t, val)
		assert.Equal(t, "", rest)
	})
}

func TestRuneTrieMatcher_FindFirstAtStartOfText(t *testing.T) {
	var matcher = NewRuneTrieMatcher()
	matcher.Put("SingleWord", 0)
	matcher.Put("Two words", 1)
	matcher.Put("caseInsenSEtivE", 2)
	matcher.Put("Two another words", 3)

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

	t.Run("must be nil by part of key", func(t *testing.T) {
		val, offset := matcher.FindFirstAtStartOfText("Two")
		assert.Nil(t, val)
		assert.Equal(t, 0, offset)
	})

	t.Run("must be case insensitive", func(t *testing.T) {
		val, offset := matcher.FindFirstAtStartOfText("CASEinsensetive")
		assert.Equal(t, 2, val)
		assert.Equal(t, 15, offset)
	})

	t.Run("must be nil and zero offset", func(t *testing.T) {
		val, offset := matcher.FindFirstAtStartOfText("")
		assert.Nil(t, val)
		assert.Equal(t, 0, offset)
	})

	t.Run("must be case insensitive without rest of text", func(t *testing.T) {
		val, offset := matcher.FindFirstAtStartOfText("caseInsenSEtivE sdfsdfsdf")
		assert.Equal(t, 2, val)
		assert.Equal(t, 15, offset)
	})
}

func Test_prepareKey(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "key == key",
			args: args{
				key: "key",
			},
			want: "key",
		},
		{
			name: "Keys to lower case",
			args: args{
				key: "KEY",
			},
			want: "key",
		},
		{
			name: "Clearing whitespace at the ends of keys",
			args: args{
				key: " key ",
			},
			want: "key",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := prepareKey(tt.args.key); got != tt.want {
				t.Errorf("prepareKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
