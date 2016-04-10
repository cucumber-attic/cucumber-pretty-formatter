package pretty

import (
	"github.com/cucumber/cucumber-pretty-formatter/events"
	"gopkg.in/cucumber/gherkin-go.v3"
)

type state int

const (
	passed state = iota
	undefined
	skipped
	failed
	ambiguous
)

type step struct {
	id               events.Identifier
	summary, details string
	step             *gherkin.Step
	container        interface{} // background or scenario
	state            state
	defID            string
	printed          bool
}

func (s *step) maxLen() int {
	var longest int
	var steps []*gherkin.Step

	switch t := s.container.(type) {
	case *gherkin.Background:
		steps = t.Steps
		longest = len(t.Keyword + t.Name)
	case *gherkin.Scenario:
		steps = t.Steps
		longest = len(t.Keyword + t.Name)
	case *gherkin.ScenarioOutline:
		steps = t.Steps
		longest = len(t.Keyword + t.Name)
	}
	longest += 2 // has colon and space between keyword and name

	for _, stp := range steps {
		maybeLonger := len(stp.Keyword+stp.Text) + 1
		if longest < maybeLonger {
			longest = maybeLonger
		}
	}
	return longest
}

func (s *step) isFirstInContainer() bool {
	var steps []*gherkin.Step
	switch t := s.container.(type) {
	case *gherkin.Background:
		steps = t.Steps
	case *gherkin.Scenario:
		steps = t.Steps
	case *gherkin.ScenarioOutline:
		steps = t.Steps
	}
	return len(steps) > 0 && steps[0] == s.step
}
