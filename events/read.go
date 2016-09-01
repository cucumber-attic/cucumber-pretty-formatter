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
		return event, json.Unmarshal(data, &event)
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
		return event, event.parseLocation()
	// START test case (scenario, outline example)
	case "TestCaseStarted":
		var event TestCaseStarted
		if err = json.Unmarshal(data, &event); err != nil {
			return nil, err
		}
		return event, event.parseLocation()
	// START step
	case "StepDefinitionFound":
		var event StepDefinitionFound
		if err = json.Unmarshal(data, &event); err != nil {
			return nil, err
		}
		return event, event.parseLocation()
	case "TestStepStarted":
		var event TestStepStarted
		if err = json.Unmarshal(data, &event); err != nil {
			return nil, err
		}
		return event, event.parseLocation()
	case "TestStepFinished":
		var event TestStepFinished
		if err = json.Unmarshal(data, &event); err != nil {
			return nil, err
		}
		return event, event.parseLocation()
	// END step
	case "TestCaseFinished":
		var event TestCaseFinished
		if err = json.Unmarshal(data, &event); err != nil {
			return nil, err
		}
		return event, event.parseLocation()
	// END test case
	case "TestRunFinished":
		var event TestRunFinished
		return event, json.Unmarshal(data, &event)
	// END feature
	case "TestAttachment":
		var event TestAttachment
		return event, json.Unmarshal(data, &event)
	default:
		return nil, fmt.Errorf(`undetermined event type: "%s" parsed from input: %s`, typ, string(data))
	}
}

func Type(data []byte) (string, error) {
	typ := struct{ Event string }{}
	return typ.Event, json.Unmarshal(data, &typ)
}
