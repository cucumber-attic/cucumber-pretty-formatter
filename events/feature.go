package events

import "github.com/cucumber/cucumber-pretty-formatter/gherkin"

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
