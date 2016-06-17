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
		return nil, err
	}
	switch typ {
	// START feature
	case "TestRunStarted":
		event := &TestRunStarted{}
		return event, json.Unmarshal(data, event)
	case "FeatureSourceRead":
		event := &FeatureSourceRead{}
		err = json.Unmarshal(data, event)
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
		event := &TestCaseStarted{}
		if err = json.Unmarshal(data, event); err != nil {
			return nil, err
		}
		return event, event.parseLocation()
	// START step
	case "StepDefinitionFound":
		event := &StepDefinitionFound{}
		if err = json.Unmarshal(data, event); err != nil {
			return nil, err
		}
		return event, event.parseLocation()
	case "TestStepStarted":
		event := &TestStepStarted{}
		if err = json.Unmarshal(data, event); err != nil {
			return nil, err
		}
		return event, event.parseLocation()
	case "TestStepPassed":
		event := &TestStepPassed{}
		if err = json.Unmarshal(data, event); err != nil {
			return nil, err
		}
		return event, event.parseLocation()
	case "TestStepFailed":
		event := &TestStepFailed{}
		if err = json.Unmarshal(data, event); err != nil {
			return nil, err
		}
		return event, event.parseLocation()
	case "TestStepSkipped":
		event := &TestStepSkipped{}
		if err = json.Unmarshal(data, event); err != nil {
			return nil, err
		}
		return event, event.parseLocation()
	case "TestStepUndefined":
		event := &TestStepUndefined{}
		if err = json.Unmarshal(data, event); err != nil {
			return nil, err
		}
		return event, event.parseLocation()
	case "TestStepAmbiguous":
		event := &TestStepAmbiguous{}
		if err = json.Unmarshal(data, event); err != nil {
			return nil, err
		}
		return event, event.parseLocation()
	// END step
	case "TestCasePassed":
		event := &TestCasePassed{}
		if err = json.Unmarshal(data, event); err != nil {
			return nil, err
		}
		return event, event.parseLocation()
	case "TestCaseFailed":
		event := &TestCaseFailed{}
		if err = json.Unmarshal(data, event); err != nil {
			return nil, err
		}
		return event, event.parseLocation()
	// END test case
	case "TestRunFinished":
		event := &TestRunFinished{}
		return event, json.Unmarshal(data, event)
	// END feature
	default:
		return nil, fmt.Errorf(`undetermined event type: "%s" parsed from input: %s`, typ, string(data))
	}
}

func Type(data []byte) (string, error) {
	typ := struct{ Event string }{}
	return typ.Event, json.Unmarshal(data, &typ)
}
