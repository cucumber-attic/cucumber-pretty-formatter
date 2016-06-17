package pretty

import (
	"fmt"
	"testing"
)

func (c color) name() string {
	switch c {
	case c_red:
		return "red"
	case c_cyan:
		return "cyan"
	case c_yellow:
		return "yellow"
	case c_black:
		return "black"
	case c_blue:
		return "blue"
	case c_green:
		return "green"
	case c_magenta:
		return "magenta"
	default:
		return "white"
	}
}

func init() {
	colorizer = func(s interface{}, c color) string {
		return fmt.Sprintf("<%s>%v</%s>", c.name(), s, c.name())
	}

	bold = func(s string) string {
		return "<bold>" + s + "</bold>"
	}
}

func TestPrintColor(t *testing.T) {
	cases := []struct {
		output string
		actual string
	}{
		{"<red>text</red>", red("text")},
		{"<green>text</green>", green("text")},
		{"<bold><white>text</white></bold>", bold(white("text"))},
	}

	for i, c := range cases {
		if c.output != c.actual {
			t.Errorf(`expected "%s", but got "%s" for case: %d`, c.output, c.actual, i)
		}
	}
}
