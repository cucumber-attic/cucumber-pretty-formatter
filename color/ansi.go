package color

import "fmt"

// a color code type
type color int

const ansiEscape = "\x1b"

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

// colorizes foreground s with color c
func cl(s interface{}, c color) string {
	return fmt.Sprintf("%s[%dm%v%s[0m", ansiEscape, c, s, ansiEscape)
}

// colorizes foreground s with bold color c
func bcl(s interface{}, c color) string {
	return fmt.Sprintf("%s[1;%dm%v%s[0m", ansiEscape, c, s, ansiEscape)
}

func Red(s interface{}) string {
	return cl(s, red)
}

func Green(s interface{}) string {
	return cl(s, green)
}

func Yellow(s interface{}) string {
	return cl(s, yellow)
}

func BoldWhite(s interface{}) string {
	return bcl(s, white)
}
