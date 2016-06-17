package pretty

import (
	"bytes"
	"testing"

	"gopkg.in/cucumber/gherkin-go.v3"
)

func TestPrintsSimpleStepWithoutArguments(t *testing.T) {
	s := &step{
		id:    "feature:4",
		state: passed,
		defID: "def.go:13",
		step: &gherkin.Step{
			Keyword: "Given",
			Text:    "is passing",
		},
	}

	var b bytes.Buffer
	p := &printer{&b}
	p.step(s, 16)

	expected := "    <green>Given is passing</green> <black># def.go:13</black>\n"
	actual := b.String()
	if expected != actual {
		t.Errorf("expected output:\n'%s'\nbut got:\n'%s'", expected, actual)
	}
}

func TestPrintsSkippedStepWithDocStringArgument(t *testing.T) {
	s := &step{
		id:    "feature:4",
		state: skipped,
		defID: "def.go:13",
		step: &gherkin.Step{
			Keyword: "Then",
			Text:    "is skipped:",
			Argument: &gherkin.DocString{
				Content:    `{"json": "object"}`,
				Delimitter: `"`,
			},
		},
	}

	var b bytes.Buffer
	p := &printer{&b}
	p.step(s, 16)

	expected := `    <cyan>Then is skipped:</cyan> <black># def.go:13</black>
      <cyan>"""</cyan>
      <cyan>{"json": "object"}</cyan>
      <cyan>"""</cyan>
`
	actual := b.String()
	if expected != actual {
		t.Errorf("expected output:\n'%s'\nbut got:\n'%s'", expected, actual)
	}
}
