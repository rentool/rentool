package matcher

import (
	"github.com/rentool/rentool/pkg/gherkin/token"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSetGet(t *testing.T) {
	var matcher = NewRuneTrieMatcher()
	matcher.Put("key1", "value1")
	matcher.Put("key2", "value2")

	assert.Equal(t, "value1", matcher.Get("key1"))
	assert.Equal(t, "value2", matcher.Get("key2"))
	assert.Equal(t, nil, matcher.Get("invalid"))
}

func TestFindAtStartOfText(t *testing.T) {
	var matcher = NewRuneTrieMatcher()
	matcher.Put("Scenario", token.Scenario)
	matcher.Put("Scenario Outline", token.Outline)

	scenario, _ := matcher.FindFirstAtStartOfText("Scenario other text")
	assert.Equal(t, token.Scenario, scenario)

	scenarioOutline, _ := matcher.FindFirstAtStartOfText("Scenario Outline other text")
	assert.Equal(t, token.Outline, scenarioOutline)
}

func TestFindAtStartOfText_givenScenarioInOneCase_thenShouldFindValueByOtherCase(t *testing.T) {
	var matcher = NewRuneTrieMatcher()
	matcher.Put("scenARIo", token.Scenario)

	scenario, _ := matcher.FindFirstAtStartOfText("Scenario other text")
	assert.Equal(t, token.Scenario, scenario)
}
