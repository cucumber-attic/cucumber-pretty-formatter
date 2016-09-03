package progress

import (
	"fmt"
	"io"
	"math"
	"regexp"
	"time"

	"github.com/cucumber/cucumber-pretty-formatter"
	"github.com/cucumber/cucumber-pretty-formatter/colors"
	"github.com/cucumber/cucumber-pretty-formatter/events"
)

var supportedVersion = regexp.MustCompile(`0\.1\.[\d]+`)

const formaterDescription = "progress formatter"
const stepsPerRow = 70

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

	steps struct {
		passed,
		failed,
		undefined,
		skipped,
		ambiguous int
	}

	cases struct {
		passed,
		failed,
		undefined, // might not be used
		skipped, // might not be used
		ambiguous int // might not be used
	}

	total int
}

func (f *format) step(status string) (err error) {
	switch status {
	case "passed":
		_, err = fmt.Fprint(f.out, colors.Green("."))
		f.steps.passed++
	case "failed":
		_, err = fmt.Fprint(f.out, colors.Red("F"))
		f.steps.failed++
	case "skipped":
		_, err = fmt.Fprint(f.out, colors.Cyan("_"))
		f.steps.skipped++
	case "undefined":
		_, err = fmt.Fprint(f.out, colors.Yellow("U"))
		f.steps.undefined++
	case "ambiguous":
		_, err = fmt.Fprint(f.out, colors.Red("A")) // @TODO: check sign and color
		f.steps.ambiguous++
	}

	if err != nil {
		return
	}

	f.total++
	if math.Mod(float64(f.total), float64(stepsPerRow)) == 0 {
		_, err = fmt.Fprintf(f.out, " %d\n", f.total)
	}
	return
}

func (f *format) fcase(status string) {
	switch status {
	case "passed":
		f.cases.passed++
	case "failed":
		f.cases.failed++
	case "skipped":
		f.cases.skipped++
	case "undefined":
		f.cases.undefined++
	case "ambiguous":
		f.cases.ambiguous++
	}
}

func (f *format) summary(e events.TestRunFinished) error {
	return nil
}

func (f *format) Event(e events.Event) error {
	switch t := e.(type) {
	case events.TestRunStarted:
		f.started = t.Timestamp.Time
		if len(t.Version) > 0 && !supportedVersion.MatchString(t.Version) {
			return fmt.Errorf("event protocol version: %s is not supported - only 0.1.x versions are.", t.Version)
		}
	case events.TestStepFinished:
		if err := f.step(t.Status); err != nil {
			return fmt.Errorf("failed to print step status to formatter output: %v", err)
		}
	case events.TestCaseFinished:
		f.fcase(t.Status)
	case events.TestRunFinished:
		if err := f.summary(t); err != nil {
			return fmt.Errorf("failed to print progress formatter summary to output: %v", err)
		}
	}
	return nil
}
