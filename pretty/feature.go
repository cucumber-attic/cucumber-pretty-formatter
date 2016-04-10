package pretty

import (
	"fmt"

	"github.com/cucumber/cucumber-pretty-formatter/events"

	"gopkg.in/cucumber/gherkin-go.v3"
)

type feature struct {
	ast        *gherkin.Feature    // gherkin AST tree of feature
	background *gherkin.Background // if not finished printing background, it won't be nil
	steps      map[string]*step    // all feature steps

	print *printer
}

// flush is always called on test case (step) execution result
// this function determines what is already printed and what it
// is able to print to output
func (ft *feature) flush(s *step) error {
	if s.printed {
		// was already printed (background, outline)
		// @TODO: may wish to print an error if second background iteration fails
		return nil
	}

	if s.isFirstInContainer() {
		if err := ft.print.container(s); err != nil {
			return err
		}
	}

	if err := ft.print.step(s); err != nil {
		return err
	}

	s.printed = true

	// scenarios and outlines
	return nil
}

func (ft *feature) step(id events.Identifier) (*step, error) {
	if s, available := ft.steps[id.ID]; available {
		return s, nil
	}

	gs, gn := ft.findStepAST(id.Line)
	if nil == gs {
		return nil, fmt.Errorf("step was not found in gherkin AST at line: \"%d\" as expected", id.Line)
	}

	ft.steps[id.ID] = &step{
		id:        id,
		step:      gs,
		container: gn,
	}
	return ft.steps[id.ID], nil
}

// findStep finds step in gherkin AST
// return step node and its container node
func (ft *feature) findStepAST(line int) (*gherkin.Step, interface{}) {
	if ft.ast.Background != nil {
		for _, s := range ft.ast.Background.Steps {
			if s.Location.Line == line {
				return s, ft.ast.Background
			}
		}
	}
	for _, s := range ft.ast.ScenarioDefinitions {
		if so, ok := s.(*gherkin.ScenarioOutline); ok {
			for _, stp := range so.Steps {
				if stp.Location.Line == line {
					return stp, so
				}
			}
		}

		if sc, ok := s.(*gherkin.Scenario); ok {
			for _, stp := range sc.Steps {
				if stp.Location.Line == line {
					return stp, sc
				}
			}
		}
	}
	return nil, nil
}
