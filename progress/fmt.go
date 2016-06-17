package progress

import (
	"io"

	"github.com/cucumber/cucumber-pretty-formatter"
	"github.com/cucumber/cucumber-pretty-formatter/events"
)

const formaterDescription = "progress formatter"

// register progress formatter
// when this package is imported
func init() {
	formatter.Register("progress", formaterDescription, func(output io.Writer) formatter.Formatter {
		return &format{
			out: output,
		}
	})
}

type format struct {
	out io.Writer
}

func (f *format) Event(e events.Event) error {
	return nil
}
