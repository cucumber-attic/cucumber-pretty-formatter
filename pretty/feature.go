package pretty

import (
	"io"

	"gopkg.in/cucumber/gherkin-go.v3"
)

type feature struct {
	ast *gherkin.Feature // gherkin AST tree of feature
	at  int              // feature pretty printed until this line
}

func (ft *feature) flush(w io.Writer) error {
	return nil
}

func (p *pretty) hasFeature(s, id string) bool {
	st, hasSuite := p.suites[s]
	if !hasSuite {
		return false
	}

	_, hasFeature := st.features[id]
	if !hasFeature {
		return false
	}
	return true
}

func (p *pretty) feature(s, id string) *feature {
	st, hasSuite := p.suites[s]
	if !hasSuite {
		p.suites[s] = &suite{
			features: make(map[string]*feature),
		}
		st = p.suites[s]
	}

	ft, hasFeature := st.features[id]
	if !hasFeature {
		st.features[id] = &feature{}
		ft = st.features[id]
	}
	return ft
}
