package pretty

import "gopkg.in/cucumber/gherkin-go.v3"

type state int

const (
	passed state = iota
	undefined
	skipped
	failed
	ambiguous
)

type step struct {
	id               string
	summary, details string
	step             *gherkin.Step
	state            state
	defID            string
}
