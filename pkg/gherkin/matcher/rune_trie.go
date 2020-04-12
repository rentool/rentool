package matcher

import "strings"

type RuneTrieMatcher struct {
	value    interface{}
	children map[rune]*RuneTrieMatcher
}

func NewRuneTrieMatcher() Matcher {
	return &RuneTrieMatcher{}
}

func (m *RuneTrieMatcher) Get(key string) interface{} {
	key = strings.ToLower(key)
	node := m
	for _, r := range key {
		node = node.children[r]
		if node == nil {
			return nil
		}
	}
	return node.value
}

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

func (m *RuneTrieMatcher) Put(key string, value interface{}) {
	key = strings.ToLower(strings.Trim(key, " "))
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
