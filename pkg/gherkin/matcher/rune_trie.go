package matcher

import "strings"

// RuneTrieMatcher Prefix tree data structure
// https://ru.wikipedia.org/wiki/%D0%9F%D1%80%D0%B5%D1%84%D0%B8%D0%BA%D1%81%D0%BD%D0%BE%D0%B5_%D0%B4%D0%B5%D1%80%D0%B5%D0%B2%D0%BE
type RuneTrieMatcher struct {
	value    interface{}
	children map[rune]*RuneTrieMatcher
}

// NewRuneTrieMatcher initialize pointer to new RuneTrieMatcher object
// Return Matcher interface
func NewRuneTrieMatcher() Matcher {
	return &RuneTrieMatcher{}
}

// Get looks for a match in the prefix tree.
// Returns the found value or nil
func (m *RuneTrieMatcher) Get(key string) interface{} {
	key = prepareKey(key)
	node := m
	for _, r := range key {
		node = node.children[r]
		if node == nil {
			return nil
		}
	}
	return node.value
}

// GetByText get the value and the rest of the text from the tree by text.
// Will return the found value or nil and the rest of the transmitted text
// that does not match the stored key.
func (m *RuneTrieMatcher) GetByText(text string) (interface{}, string) {
	node := m
	text = strings.ToLower(text)
	var foundValue interface{}
	foundEndOffset := 0
	for i, r := range text {
		node = node.children[r]
		if nil == node {
			return foundValue, text[foundEndOffset:]
		}

		if nil != node.value {
			foundValue = node.value
			foundEndOffset = i + 1
		}
	}

	return foundValue, text[foundEndOffset:]
}

// FindFirstAtStartOfText get the value and the offset of the text from the tree by text.
// Will return the found value or nil and the offset of the transmitted text
// that does match the stored key.
func (m *RuneTrieMatcher) FindFirstAtStartOfText(text string) (interface{}, int) {
	node := m
	text = strings.ToLower(text)
	var foundValue interface{}
	foundEndOffset := 0
	for i, r := range text {
		node = node.children[r]
		if nil == node {
			return foundValue, foundEndOffset
		}

		if nil != node.value {
			foundValue = node.value
			foundEndOffset = i + 1
		}
	}

	return foundValue, foundEndOffset
}

// Put puts the value in the key row in the tree if the tree does not have a key,
// or changes the value if it has
func (m *RuneTrieMatcher) Put(key string, value interface{}) {
	key = prepareKey(key)
	node := m
	for _, r := range key {
		child, _ := node.children[r]
		if child == nil {
			if node.children == nil {
				node.children = map[rune]*RuneTrieMatcher{}
			}
			child = new(RuneTrieMatcher)
			node.children[r] = child
		}
		node = child
	}

	node.value = value
}

func prepareKey(key string) string {
	return strings.ToLower(strings.Trim(key, " "))
}
