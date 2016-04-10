package pretty

import (
	"io"

	"github.com/cucumber/cucumber-pretty-formatter"
	"github.com/cucumber/cucumber-pretty-formatter/events"
)

// register pretty formatter
func init() {
	formatter.Register("pretty", formatter.Initializer(func(output io.Writer) formatter.Formatter {
		return &pretty{
			output: output,
			suites: make(map[string]*suite),
		}
	}))
}

// synchronous as first implementation
type pretty struct {
	output io.Writer

	suites map[string]*suite
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
