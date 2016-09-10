package progress

import (
	"fmt"
	"io"
	"math"
	"regexp"
	"strings"
	"time"

	"github.com/cucumber/cucumber-pretty-formatter"
	"github.com/cucumber/cucumber-pretty-formatter/colors"
	"github.com/cucumber/cucumber-pretty-formatter/events"
	"github.com/cucumber/cucumber-pretty-formatter/gherkin"
)

var supportedVersion = regexp.MustCompile(`0\.1\.[\d]+`)

var trim = strings.TrimSpace

const formaterDescription = "progress formatter"
const stepsPerRow = 70

// register progress formatter
// when this package is imported
func init() {
	formatter.Register("progress", formaterDescription, func(output io.Writer) formatter.Formatter {
		return &format{
			out:   output,
			steps: &stats{},
			cases: &stats{},
		}
	})
}

// stats are for step and scenario summary
type stats struct {
	passed,
	failed,
	undefined,
	pending,
	skipped,
	ambiguous int
}

func (s *stats) total() int {
	return s.passed + s.failed + s.undefined + s.skipped + s.ambiguous + s.pending
}

func (s *stats) summary(typ string) string {
	var parts []string

	if s.passed > 0 {
		parts = append(parts, colors.Green(fmt.Sprintf("%d passed", s.passed)))
	}

	if s.failed > 0 {
		parts = append(parts, colors.Red(fmt.Sprintf("%d failed", s.failed)))
	}

	if s.pending > 0 {
		parts = append(parts, colors.Yellow(fmt.Sprintf("%d pending", s.pending)))
	}

	if s.undefined > 0 {
		parts = append(parts, colors.Yellow(fmt.Sprintf("%d undefined", s.undefined)))
	}

	if s.ambiguous > 0 {
		parts = append(parts, colors.Cyan(fmt.Sprintf("%d ambiguous", s.ambiguous)))
	}

	if s.skipped > 0 {
		parts = append(parts, colors.Cyan(fmt.Sprintf("%d skipped", s.skipped)))
	}

	if s.total() == 0 {
		return fmt.Sprintf("No %s\n", typ)
	}
	return fmt.Sprintf("%d %s (%s)\n", s.total(), typ, strings.Join(parts, ", "))
}

type failure struct {
	step  *gherkin.Step
	event events.TestStepFinished
}

type format struct {
	out      io.Writer
	started  time.Time
	feature  *gherkin.Feature
	failures []*failure

	steps, cases *stats
}

func (f *format) locateStep(line int) *gherkin.Step {
	if f.feature.Background != nil {
		for _, step := range f.feature.Background.Steps {
			if step.Location.Line == line {
				return step
			}
		}
	}
	for _, def := range f.feature.ScenarioDefinitions {
		if outline, ok := def.(*gherkin.ScenarioOutline); ok {
			for _, step := range outline.Steps {
				if step.Location.Line == line {
					return step
				}
			}
		}

		if scenario, ok := def.(*gherkin.Scenario); ok {
			for _, step := range scenario.Steps {
				if step.Location.Line == line {
					return step
				}
			}
		}
	}
	return nil
}

func (f *format) step(e events.TestStepFinished) (err error) {
	switch e.Status {
	case "passed":
		_, err = fmt.Fprint(f.out, colors.Green("."))
		f.steps.passed++
	case "failed":
		_, err = fmt.Fprint(f.out, colors.Red("F"))
		f.steps.failed++
		if step := f.locateStep(e.Line); step != nil {
			f.failures = append(f.failures, &failure{step, e})
		} else {
			return fmt.Errorf("step could not be located in the feature by location: %s", e.Location)
		}
	case "skipped":
		_, err = fmt.Fprint(f.out, colors.Cyan("_"))
		f.steps.skipped++
	case "undefined":
		_, err = fmt.Fprint(f.out, colors.Yellow("U"))
		f.steps.undefined++
	case "pending":
		_, err = fmt.Fprint(f.out, colors.Yellow("P"))
		f.steps.pending++
	case "ambiguous":
		_, err = fmt.Fprint(f.out, colors.Red("A")) // @TODO: check sign and color
		f.steps.ambiguous++
	}

	if err != nil {
		return
	}

	if math.Mod(float64(f.steps.total()), float64(stepsPerRow)) == 0 {
		_, err = fmt.Fprintf(f.out, " %d\n", f.steps.total())
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
	var buf string
	if left := math.Mod(float64(f.steps.total()), float64(stepsPerRow)); left != 0 {
		buf += fmt.Sprintf("%s %d\n", strings.Repeat(" ", stepsPerRow-int(left)), f.steps.total())
	}

	if f.steps.total() > 0 {
		buf += "\n"
	}

	if len(f.failures) > 0 {
		buf += "\n--- " + colors.Red("Failed steps:") + "\n"
		for _, fail := range f.failures {
			buf += strings.Repeat(" ", 4)
			buf += colors.Red(trim(fail.step.Keyword) + " " + trim(fail.step.Text))
			buf += colors.Black(" # " + fail.event.Location)
			buf += "\n"

			buf += strings.Repeat(" ", 6)
			buf += colors.Red("Error: ")
			buf += colors.RedB(fail.event.Summary)
			buf += "\n"
		}
		buf += "\n"
	}

	buf += f.cases.summary("scenarios")
	buf += f.steps.summary("steps")

	finishedAt := e.Timestamp.Unix()
	usage := finishedAt.Sub(f.started).String()
	if len(e.Memory.String()) > 0 {
		usage += fmt.Sprintf(" (%s)", e.Memory)
	}

	buf += usage

	if len(e.Snippets) > 0 {
		buf += "\n\n" + colors.Yellow(e.Snippets)
	}
	_, err := fmt.Fprintln(f.out, buf)
	return err
}

func (f *format) Event(e events.Event) error {
	switch t := e.(type) {
	case events.TestRunStarted:
		f.started = t.Timestamp.Unix()
		if len(t.Version) > 0 && !supportedVersion.MatchString(t.Version) {
			return fmt.Errorf("event protocol version: %s is not supported - only 0.1.x versions are.", t.Version)
		}
	case events.TestSource:
		f.feature = t.Feature
	case events.TestStepFinished:
		if err := f.step(t); err != nil {
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
