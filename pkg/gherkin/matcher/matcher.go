package matcher

type Matcher interface {
	Get(key string) interface{}
	Put(key string, value interface{})
	FindFirstAtStartOfText(text string) (value interface{}, endOffset int)
}

type ReadOnlyMatcher interface {
	Get(key string) interface{}
}
