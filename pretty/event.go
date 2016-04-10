package pretty

import (
	"fmt"

	"github.com/cucumber/cucumber-pretty-formatter/events"
)

func (p *pretty) Event(e events.Event) error {
	switch t := e.(type) {
	case *events.FeatureSourceRead:
		if p.hasFeature(t.Identifier) {
			return fmt.Errorf(`unexpected feature reread event, feature: "%s", was already read for suite: "%s"`, t.ID, t.Suite)
		}
		ft := p.feature(t.Identifier)
		ft.ast = t.Feature
		ft.background = ft.ast.Background
		if err := ft.print.header(ft.ast); err != nil {
			return err
		}
	case *events.StepDefinitionFound:
		ft := p.feature(t.Identifier)
		st, err := ft.step(t.Identifier)
		if err != nil {
			return err
		}
		st.defID = t.DefID
	case *events.TestStepFailed:
		ft := p.feature(t.Identifier)
		st, err := ft.step(t.Identifier)
		if err != nil {
			return err
		}
		st.summary = t.Error
		st.details = t.Trace
		st.state = failed
		if err := ft.flush(st); err != nil {
			return fmt.Errorf("failed to write part of feature: %s", err)
		}
	case *events.TestStepPassed:
		ft := p.feature(t.Identifier)
		st, err := ft.step(t.Identifier)
		if err != nil {
			return err
		}
		st.state = passed
		if err := ft.flush(st); err != nil {
			return fmt.Errorf("failed to write part of feature: %s", err)
		}
	case *events.TestStepSkipped:
		ft := p.feature(t.Identifier)
		st, err := ft.step(t.Identifier)
		if err != nil {
			return err
		}
		st.state = skipped
		if err := ft.flush(st); err != nil {
			return fmt.Errorf("failed to write part of feature: %s", err)
		}
	case *events.TestStepUndefined:
		ft := p.feature(t.Identifier)
		st, err := ft.step(t.Identifier)
		if err != nil {
			return err
		}
		st.state = undefined
		st.summary = t.Todo
		st.details = t.Snippet
		if err := ft.flush(st); err != nil {
			return fmt.Errorf("failed to write part of feature: %s", err)
		}
	}
	return nil
}
