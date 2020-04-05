package matcher

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGet(t *testing.T) {
	var matcher = NewRuneTrieMatcher()
	matcher.Put("key1", "value1")
	matcher.Put("key2", "value2")

	assert.Equal(t, "value1", matcher.Get("key1"))
	assert.Equal(t, "value2", matcher.Get("key2"))
	assert.Equal(t, nil, matcher.Get("invalid"))
}
