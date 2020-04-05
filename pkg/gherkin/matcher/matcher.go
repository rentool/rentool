package matcher

type Matcher interface {
	Get(key string) interface{}
	Put(key string, value interface{})
}

type ReadOnlyMatcher interface {
	Get(key string) interface{}
}
