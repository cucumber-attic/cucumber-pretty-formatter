package colors

import (
	"fmt"
	"testing"
)

func init() {
	Colorize = func(s interface{}, c Color) string {
		return fmt.Sprintf("<%s>%v</%s>", c.Name(), s, c.Name())
	}

	Bold = func(s string) string {
		return "<bold>" + s + "</bold>"
	}
}

func TestPrintColor(t *testing.T) {
	cases := []struct {
		output string
		actual string
	}{
		{"<red>text</red>", Red("text")},
		{"<green>text</green>", Green("text")},
		{"<bold><white>text</white></bold>", Bold(White("text"))},
	}

	for i, c := range cases {
		if c.output != c.actual {
			t.Errorf(`expected "%s", but got "%s" for case: %d`, c.output, c.actual, i)
		}
	}
}
