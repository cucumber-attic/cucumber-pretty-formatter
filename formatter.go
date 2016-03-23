package formatter

import (
	"io"

	"github.com/cucumber/cucumber-pretty-formatter/events"
)

var formatters = make(map[string]Initializer)

type Initializer func(io.Writer) Formatter

type Formatter interface {
	Event(events.Event) error
}

func Register(name string, fn Initializer) {
	formatters[name] = fn
}
