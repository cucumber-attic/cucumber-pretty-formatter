package events

import "github.com/cucumber/cucumber-pretty-formatter/gherkin"

type TestRunStarted struct {
	Version   string          `json:"version"`
	Timestamp UnixTimestampMS `json:"timestamp"`
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
	Timestamp UnixTimestampMS `json:"timestamp"`
}

type TestStepStarted struct {
	identifier
	Timestamp UnixTimestampMS `json:"timestamp"`
}

type TestStepFinished struct {
	identifier
	Timestamp UnixTimestampMS `json:"timestamp"`
	Status    string          `json:"status"`
	Summary   string          `json:"summary"`
	Details   string          `json:"details"`
}

type TestCaseFinished struct {
	identifier
	Timestamp UnixTimestampMS `json:"timestamp"`
	Status    string          `json:"status"`
}

type TestRunFinished struct {
	Timestamp UnixTimestampMS `json:"timestamp"`
	Status    string          `json:"status"`
	Memory    string          `json:"memory"`
}

type TestAttachment struct {
	identifier
	Timestamp UnixTimestampMS `json:"timestamp"`
	MimeType  string          `json:"mimeType"`
	Data      []byte          `json:"data"`
	Encoding  string          `json:"encoding"`
}
