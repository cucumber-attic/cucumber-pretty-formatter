package events

import "github.com/cucumber/cucumber-pretty-formatter/gherkin"

type TestRunStarted struct {
	ProtocolVersion string          `json:"version"`
	Timestamp       UnixTimestampMS `json:"timestamp"`
	RunID           string          `json:"run_id"`
}

type TestSource struct {
	identifier
	Source  string           `json:"source"`
	Feature *gherkin.Feature `json:"-"`
}

type StepDefinitionFound struct {
	identifier
	Suite     string   `json:"suite"`
	DefID     string   `json:"definition_id"`
	Arguments [][2]int `json:"arguments"`
}

type TestCaseStarted struct {
	identifier
	Timestamp UnixTimestampMS `json:"timestamp"`
	Suite     string          `json:"suite"`
}

type TestStepStarted struct {
	identifier
	Timestamp UnixTimestampMS `json:"timestamp"`
	Suite     string          `json:"suite"`
}

type TestStepFinished struct {
	identifier
	Timestamp UnixTimestampMS `json:"timestamp"`
	Suite     string          `json:"suite"`
	Status    string          `json:"status"`
	Summary   string          `json:"summary"`
	Details   string          `json:"details"`
}

type TestCaseFinished struct {
	identifier
	Timestamp UnixTimestampMS `json:"timestamp"`
	Suite     string          `json:"suite"`
	Status    string          `json:"status"`
}

type TestRunFinished struct {
	Timestamp UnixTimestampMS `json:"timestamp"`
	Suite     string          `json:"suite"`
	Status    string          `json:"status"`
	Memory    string          `json:"memory"`
	RunID     string          `json:"run_id"`
}

type TestAttachment struct {
	identifier
	Timestamp UnixTimestampMS `json:"timestamp"`
	Suite     string          `json:"suite"`
	MimeType  string          `json:"mimeType"`
	Data      []byte          `json:"data"`
	Encoding  string          `json:"encoding"`
}
