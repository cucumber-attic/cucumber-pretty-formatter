package events

type StepDefinitionFound struct {
	Identifier
	DefID string `json:"definition_id"`
}

type TestStepStarted struct {
	Identifier
}

type TestStepPassed struct {
	Identifier
}

type TestStepSkipped struct {
	Identifier
}

type TestStepAmbiguous struct {
	Identifier
	// @TODO: define details
}

type TestStepUndefined struct {
	Identifier
	Todo    string `json:"todo"`
	Snippet string `json:"snippet"`
}

type TestStepFailed struct {
	Identifier
	Error string `json:"error"`
	Trace string `json:"trace"`
}
