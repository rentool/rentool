package matcher

type Matcher interface {
	Get(key string) interface{}
	GetByText(text string) (value interface{}, rest string)
	Put(key string, value interface{})
	FindFirstAtStartOfText(text string) (value interface{}, endOffset int)
}

type ReadOnlyMatcher interface {
	Get(key string) interface{}
	GetByText(text string) (value interface{}, rest string)
	FindFirstAtStartOfText(text string) (value interface{}, endOffset int)
}
