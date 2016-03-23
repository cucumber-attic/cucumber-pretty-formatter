package events

type TestCaseStarted struct {
	id
}

type TestCasePassed struct {
	id
}

type TestCaseFailed struct {
	id
}
