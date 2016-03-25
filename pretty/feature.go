package pretty

import (
	"fmt"

	"github.com/cucumber/cucumber-pretty-formatter/events"

	"gopkg.in/cucumber/gherkin-go.v3"
)

type feature struct {
	ast        *gherkin.Feature    // gherkin AST tree of feature
	background *gherkin.Background // if not finished printing background, it won't be nil
	at         int                 // feature pretty printed until this line
	steps      map[string]*step    // all feature steps

	print *printer
}

// flush is always called on test case (step) execution result
// this function determines what is already printed and what it
// is able to print to output
func (ft *feature) flush(s *step) error {
	// feature header
	if ft.at < ft.ast.Location.Line {
		if err := ft.print.header(ft.ast); err != nil {
			return err
		}
		ft.at = ft.ast.Location.Line
	}

	// background
	if ft.background != nil {
		bg := ft.background
		if ft.at < bg.Location.Line {
			ft.print.background(bg)
			ft.at = bg.Location.Line
		}

		for _, bgStep := range bg.Steps {
			if bgStep.Location.Line == s.step.Location.Line {
				ft.print.backgroundStep(s, bg)
				ft.at = s.step.Location.Line
				break
			}
		}
	}
	return nil
}

func (ft *feature) step(id events.Identifier) (*step, error) {
	if s, available := ft.steps[id.ID]; available {
		return s, nil
	}

	gs := ft.findStepAST(id.Line)
	if nil == gs {
		return nil, fmt.Errorf("step was not found in gherkin AST at line: \"%d\" as expected", id.Line)
	}

	ft.steps[id.ID] = &step{
		id:   id.ID,
		step: gs,
	}
	return ft.steps[id.ID], nil
}

// findStep finds step in gherkin AST
func (ft *feature) findStepAST(line int) *gherkin.Step {
	if ft.ast.Background != nil {
		for _, s := range ft.ast.Background.Steps {
			if s.Location.Line == line {
				return s
			}
		}
	}
	for _, s := range ft.ast.ScenarioDefinitions {
		if so, ok := s.(*gherkin.ScenarioOutline); ok {
			for _, stp := range so.Steps {
				if stp.Location.Line == line {
					return stp
				}
			}
		}

		if sc, ok := s.(*gherkin.Scenario); ok {
			for _, stp := range sc.Steps {
				if stp.Location.Line == line {
					return stp
				}
			}
		}
	}
	return nil
}

func (p *pretty) hasFeature(id events.Identifier) bool {
	st, hasSuite := p.suites[id.Suite]
	if !hasSuite {
		return false
	}

	_, hasFeature := st.features[id.ID]
	if !hasFeature {
		return false
	}
	return true
}

func (p *pretty) feature(id events.Identifier) *feature {
	st, hasSuite := p.suites[id.Suite]
	if !hasSuite {
		p.suites[id.Suite] = &suite{
			features: make(map[string]*feature),
		}
		st = p.suites[id.Suite]
	}

	ft, hasFeature := st.features[id.ID]
	if !hasFeature {
		st.features[id.ID] = &feature{
			steps: make(map[string]*step),
			print: &printer{output: p.output},
		}
		ft = st.features[id.ID]
	}
	return ft
}
