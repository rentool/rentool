package matcher

// Matcher matcher interface
type Matcher interface {
	Put(key string, value interface{})
	ReadOnlyMatcher
}

// ReadOnlyMatcher read only matcher interface
type ReadOnlyMatcher interface {
	Get(key string) interface{}
	GetByText(text string) (value interface{}, rest string)
	FindFirstAtStartOfText(text string) (value interface{}, endOffset int)
}
