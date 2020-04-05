package matcher

type RuneTrieMatcher struct {
	value    interface{}
	children map[rune]*RuneTrieMatcher
}

func NewRuneTrieMatcher() Matcher {
	return &RuneTrieMatcher{}
}

func (m *RuneTrieMatcher) Get(key string) interface{} {
	node := m
	for _, r := range key {
		node = node.children[r]
		if node == nil {
			return nil
		}
	}
	return node.value
}

func (m *RuneTrieMatcher) Put(key string, value interface{}) {
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
