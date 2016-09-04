package events

import "github.com/cucumber/cucumber-pretty-formatter/gherkin"

type TestRunStarted struct {
	Version   string `json:"version"`
	Timestamp int64  `json:"timestamp"`
}

type TestSource struct {
	identifier
	Source  string           `json:"source"`
	Feature *gherkin.Feature `json:"-"`
}

type StepDefinitionFound struct {
	identifier
	Definition string   `json:"definition"`
	Arguments  [][2]int `json:"arguments"`
}

type TestCaseStarted struct {
	identifier
	Timestamp int64 `json:"timestamp"`
}

type TestStepStarted struct {
	identifier
	Timestamp int64 `json:"timestamp"`
}

type TestStepFinished struct {
	identifier
	Timestamp int64  `json:"timestamp"`
	Status    string `json:"status"`
	Summary   string `json:"summary"`
	Details   string `json:"details"`
}

type TestCaseFinished struct {
	identifier
	Timestamp int64  `json:"timestamp"`
	Status    string `json:"status"`
}

type TestRunFinished struct {
	Timestamp int64  `json:"timestamp"`
	Status    string `json:"status"`
	Memory    string `json:"memory"`
	Snippets  string `json:"snippets"`
}

type TestAttachment struct {
	identifier
	Timestamp int64  `json:"timestamp"`
	Mime      string `json:"mime"`
	Data      []byte `json:"data"`
	Encoding  string `json:"encoding"`
}
