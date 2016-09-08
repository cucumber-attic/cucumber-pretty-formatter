package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/cucumber/cucumber-pretty-formatter"
	_ "github.com/cucumber/cucumber-pretty-formatter/progress"
)

const defaultFormat = "progress"

var formats formatters

var options = formatter.Options{
	EventInputStream: struct {
		io.Reader
		io.Closer
	}{os.Stdin, ioutil.NopCloser(nil)},
}

func init() {
	flag.BoolVar(&options.NoColors, "no-colors", false, "Disable ansi colors.")
	flag.Var(&formats, "format", "Format to use, currently only: progress")
	flag.Var(&formats, "f", "Format to use, currently only: progress")
}

func main() {
	flag.Parse()

	options.Formats = formats
	if status, err := formatter.Run(options); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	} else {
		os.Exit(status)
	}
}

// formatters is a custom flag to support list of formats
type formatters []string

func (f *formatters) String() string {
	return defaultFormat
}

func (f *formatters) Set(formatter string) error {
	*f = append(*f, formatter)
	return nil
}
