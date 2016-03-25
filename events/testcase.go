package events

type TestCaseStarted struct {
	Identifier
}

type TestCasePassed struct {
	Identifier
}

type TestCaseFailed struct {
	Identifier
}
