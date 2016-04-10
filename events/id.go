package events

import (
	"fmt"
	"strconv"
	"strings"
)

type Identifier struct {
	Location string // feature identifier
	Suite    string // on which suite the feature runs in
	Path     string // feature path
	Line     int    // feature identification line
}

func (i *Identifier) parseLocation() (err error) {
	i.Path, i.Line, err = Location(i.Location)
	return
}

func Location(s string) (string, int, error) {
	delimIdx := strings.LastIndex(s, ":")
	if delimIdx == -1 {
		return "", 0, fmt.Errorf("could not parse location, line delimiter not found from: %s", s)
	}

	line, err := strconv.Atoi(s[delimIdx+1:])
	if err != nil {
		return "", 0, fmt.Errorf("could not parse line number from: \"%s\" as integer: %s", s, err)
	}

	return s[:delimIdx], line, nil
}
