package pretty

import (
	"io"

	"github.com/cucumber/cucumber-pretty-formatter"
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
