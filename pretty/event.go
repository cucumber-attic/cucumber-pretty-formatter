package pretty

import (
	"fmt"

	"github.com/cucumber/cucumber-pretty-formatter/events"
)

func (p *pretty) Event(e events.Event) error {
	switch t := e.(type) {
	case *events.FeatureSourceRead:
		if p.hasFeature(t.Suite, t.ID) {
			return fmt.Errorf(`unexpected feature reread event, feature: "%s", was already read for suite: "%s"`, t.ID, t.Suite)
		}
		ft := p.feature(t.Suite, t.ID)
		ft.ast = t.Feature
		if err := ft.flush(p.output); err != nil {
			return fmt.Errorf("failed to write part of feature: %s", err)
		}
	}
	return nil
}
