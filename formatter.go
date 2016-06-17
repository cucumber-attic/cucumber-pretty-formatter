package formatter

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/cucumber/cucumber-pretty-formatter/events"
)

// formatter is a registered formatter type
type formatter struct {
	name        string
	description string
	build       initializer
}

// formatters is a list type of registered formatters
type formatters []formatter

// find lookups a formatter
func (f formatters) find(name string) initializer {
	for _, fmt := range f {
		if fmt.name == name {
			return fmt.build
		}
	}
	return nil
}

// all registered formatters
var all formatters

// initializer is a func type witch used when
// registering a new Formatter and to initialize
// it with an output io.Writer. It should return a Formatter
// interface
type initializer func(io.Writer) Formatter

// Formatter is an interface for any event based formatter
// it must be able to consume events and build output
// based on these events
type Formatter interface {
	Event(events.Event) error
}

// Register registers a formatter by given name and description
func Register(name, description string, fn initializer) {
	all = append(all, formatter{name, description, fn})
}

// Run scans given input for events and runs it through configured
// formatters determined from option flags.
func Run(in io.Reader) error {
	// @TODO: will need to read flags and initialize
	// writers + stream events to all formatters configured
	build := all.find("progress")
	if nil == build {
		return fmt.Errorf("formatter: '%s' is not available", "progress")
	}
	// @TODO output should be configured from flags
	// @TODO ansicolor support for windows
	f := build(os.Stdout)

	// @TODO many formatters may be spawned in parallel if configured
	scanner := bufio.NewScanner(in)
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
