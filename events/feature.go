package events

import "gopkg.in/cucumber/gherkin-go.v3"

type TestRunStarted struct {
	ProtocolVersion string `json:"protocol_version"`
}

type TestRunFinished struct {
	Time   string
	Memory string
}

type FeatureSourceRead struct {
	Identifier
	Source  string
	Feature *gherkin.Feature
}
