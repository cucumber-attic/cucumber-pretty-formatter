package progress

import (
	"fmt"
	"io"
	"regexp"
	"time"

	"github.com/cucumber/cucumber-pretty-formatter"
	"github.com/cucumber/cucumber-pretty-formatter/events"
)

var supportedProtocol = regexp.MustCompile(`0\.1\.[\d]+`)

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
	out     io.Writer
	started time.Time
}

func (f *format) Event(e events.Event) error {
	switch t := e.(type) {
	case events.TestRunStarted:
		f.started = t.Timestamp.Time
		if !supportedProtocol.MatchString(t.ProtocolVersion) {
			return fmt.Errorf("event protocol version: %s is not supported - only 0.1.x versions are.", t.ProtocolVersion)
		}
	}
	return nil
}
