package pretty

import (
	"fmt"
	"strings"
)

const ansiEscape = "\x1b"

// a color code type
type color int

// some ansi colors
const (
	c_black color = iota + 30
	c_red
	c_green
	c_yellow
	c_blue
	c_magenta
	c_cyan
	c_white
)

var colorizer = func(s interface{}, c color) string {
	return fmt.Sprintf("%s[%dm%v%s[0m", ansiEscape, c, s, ansiEscape)
}

var bold = func(s string) string {
	return strings.Replace(s, ansiEscape+"[", ansiEscape+"[1;", 1)
}

func red(s interface{}) string {
	return colorizer(s, c_red)
}

func green(s interface{}) string {
	return colorizer(s, c_green)
}

func cyan(s interface{}) string {
	return colorizer(s, c_cyan)
}

func black(s interface{}) string {
	return colorizer(s, c_black)
}

func yellow(s interface{}) string {
	return colorizer(s, c_yellow)
}

func white(s interface{}) string {
	return colorizer(s, c_white)
}
