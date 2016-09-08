package formatter

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/cucumber/cucumber-pretty-formatter/colors"
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
func Run(opts Options) (status int, err error) {
	defer opts.EventInputStream.Close()

	var output io.Writer = os.Stdout
	if opts.NoColors {
		output = colors.Uncolored(output)
	} else {
		output = colors.Colored(output)
	}

	// @TODO output should be configured with formatter name
	formats := make([]Formatter, len(opts.Formats))
	for i, name := range opts.Formats {
		if build := all.find(name); build == nil {
			return status, fmt.Errorf("formatter: '%s' is not available", name)
		} else {
			formats[i] = build(output)
		}
	}

	// @TODO many formatters may be spawned in parallel if configured
	scanner := bufio.NewScanner(opts.EventInputStream)
	for scanner.Scan() {
		ev, err := events.Read(scanner.Bytes())
		if err != nil {
			return status, err
		}
		for _, f := range formats {
			if err := f.Event(ev); err != nil {
				return status, err
			}
		}
	}
	return status, scanner.Err()
}
