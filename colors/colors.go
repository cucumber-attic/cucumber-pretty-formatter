package colors

import (
	"fmt"
	"strings"
)

const ansiEscape = "\x1b"

// a color code type
type Color int

// some ansi colors
const (
	black Color = iota + 30
	red
	green
	yellow
	blue
	magenta
	cyan
	white
)

func (c Color) Name() string {
	switch c {
	case red:
		return "red"
	case cyan:
		return "cyan"
	case yellow:
		return "yellow"
	case black:
		return "black"
	case blue:
		return "blue"
	case green:
		return "green"
	case magenta:
		return "magenta"
	default:
		return "white"
	}
}

var Colorize = func(s interface{}, c Color) string {
	return fmt.Sprintf("%s[%dm%v%s[0m", ansiEscape, c, s, ansiEscape)
}

var Bold = func(s string) string {
	return strings.Replace(s, ansiEscape+"[", ansiEscape+"[1;", 1)
}

func Red(s interface{}) string {
	return Colorize(s, red)
}

func Green(s interface{}) string {
	return Colorize(s, green)
}

func Cyan(s interface{}) string {
	return Colorize(s, cyan)
}

func Black(s interface{}) string {
	return Colorize(s, black)
}

func Yellow(s interface{}) string {
	return Colorize(s, yellow)
}

func White(s interface{}) string {
	return Colorize(s, white)
}
