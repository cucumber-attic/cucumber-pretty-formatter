package events

type StepDefinitionFound struct {
	id
	DefID string `json:"definition_id"`
}

type TestStepStarted struct {
	id
}

type TestStepPassed struct {
	id
}

type TestStepSkipped struct {
	id
}

type TestStepAmbiguous struct {
	id
	// @TODO: define details
}

type TestStepUndefined struct {
	id
	Todo    string `json:"todo"`
	Snippet string `json:"snippet"`
}

type TestStepFailed struct {
	id
	Error string `json:"error"`
	Trace string `json:"trace"`
}
