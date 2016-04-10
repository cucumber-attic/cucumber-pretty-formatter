package formatter

import (
	"bufio"
	"fmt"
	"io"
	"os"

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

func Run() error {
	// @TODO: will need to read flags and initialize
	// writers + stream events to all formatters configured
	initer, ok := formatters["pretty"]
	if !ok {
		return fmt.Errorf("formatter: '%s' is not available", "pretty")
	}
	f := initer(os.Stdout)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		ev, err := events.Read(scanner.Bytes())
		if err != nil {
			return err
		}
		if err := f.Event(ev); err != nil {
			return err
		}
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}
