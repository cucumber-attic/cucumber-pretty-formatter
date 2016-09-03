package events

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/cucumber/cucumber-pretty-formatter/gherkin"
)

type Event interface{}

func Read(data []byte) (Event, error) {
	typ, err := Type(data)
	if err != nil {
		return nil, fmt.Errorf("parse event type from %s: %v", string(data), err)
	}
	switch typ {
	// START feature
	case "TestRunStarted":
		var event TestRunStarted
		if err := json.Unmarshal(data, &event); err != nil {
			return nil, err
		}
		return event, nil
	case "TestSource":
		var event TestSource
		err = json.Unmarshal(data, &event)
		if err != nil {
			return nil, err
		}

		event.Feature, err = gherkin.ParseFeature(strings.NewReader(event.Source))
		if err != nil {
			return nil, err
		}

		if err := event.parseLocation(); err != nil {
			return nil, err
		}
		return event, nil
	// START test case (scenario, outline example)
	case "TestCaseStarted":
		var event TestCaseStarted
		if err = json.Unmarshal(data, &event); err != nil {
			return nil, err
		}

		if err := event.parseLocation(); err != nil {
			return nil, err
		}
		return event, nil
	// START step
	case "StepDefinitionFound":
		var event StepDefinitionFound
		if err = json.Unmarshal(data, &event); err != nil {
			return nil, err
		}

		if err := event.parseLocation(); err != nil {
			return nil, err
		}
		return event, nil
	case "TestStepStarted":
		var event TestStepStarted
		if err = json.Unmarshal(data, &event); err != nil {
			return nil, err
		}
		if err := event.parseLocation(); err != nil {
			return nil, err
		}
		return event, nil
	case "TestStepFinished":
		var event TestStepFinished
		if err = json.Unmarshal(data, &event); err != nil {
			return nil, err
		}
		if err := event.parseLocation(); err != nil {
			return nil, err
		}
		return event, nil
	// END step
	case "TestCaseFinished":
		var event TestCaseFinished
		if err = json.Unmarshal(data, &event); err != nil {
			return nil, err
		}
		if err := event.parseLocation(); err != nil {
			return nil, err
		}
		return event, nil
	// END test case
	case "TestRunFinished":
		var event TestRunFinished
		if err := json.Unmarshal(data, &event); err != nil {
			return nil, err
		}
		return event, nil
	// END feature
	case "TestAttachment":
		var event TestAttachment
		if err := json.Unmarshal(data, &event); err != nil {
			return nil, err
		}
		return event, nil
	default:
		return nil, fmt.Errorf(`undetermined event type: "%s" parsed from input: %s`, typ, string(data))
	}
}

func Type(data []byte) (string, error) {
	typ := struct{ Event string }{}
	return typ.Event, json.Unmarshal(data, &typ)
}
