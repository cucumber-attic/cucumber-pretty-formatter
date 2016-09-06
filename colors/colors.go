package colors

import (
	"fmt"
	"strings"
)

const ansiEscape = "\x1b"

// a color code type
type color int

// some ansi colors
const (
	black color = iota + 30
	red
	green
	yellow
	blue
	magenta
	cyan
	white
)

func colorize(s interface{}, c color) string {
	return fmt.Sprintf("%s[%dm%v%s[0m", ansiEscape, c, s, ansiEscape)
}

func bold(s string) string {
	return strings.Replace(s, ansiEscape+"[", ansiEscape+"[1;", 1)
}

func Red(s interface{}) string {
	return colorize(s, red)
}

func RedB(s interface{}) string {
	return bold(Red(s))
}

func Green(s interface{}) string {
	return colorize(s, green)
}

func Cyan(s interface{}) string {
	return colorize(s, cyan)
}

func Black(s interface{}) string {
	return colorize(s, black)
}

func Yellow(s interface{}) string {
	return colorize(s, yellow)
}

func White(s interface{}) string {
	return colorize(s, white)
}

func WhiteB(s interface{}) string {
	return bold(White(s))
}
