package gherkin

type Feature struct {
	Text        string
	Description string
	Background  *Background
	Scenarios   []*Scenario
	Tags        []string
}

type Background struct {
	Text        string
	Description string
	Steps       []*Step
	Tags        []string
}

type Scenario struct {
	Text        string
	Description string
	Steps       []*Step
	Tags        []string
}

type Step struct {
	Text string
	Tags []string
}
