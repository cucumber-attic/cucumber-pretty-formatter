package events

import (
	"fmt"
	"strconv"
	"strings"
)

type identifier struct {
	Location string `json:"location"` // feature identifier
	Path     string `json:"-"`        // feature path
	Line     int    `json:"-"`        // feature identification line
}

func (i *identifier) parseLocation() (err error) {
	i.Path, i.Line, err = Location(i.Location)
	return
}

func Location(s string) (string, int, error) {
	delimIdx := strings.LastIndex(s, ":")
	if delimIdx == -1 {
		return "", 0, fmt.Errorf("could not parse location, line delimiter not found in: %s", s)
	}

	line, err := strconv.Atoi(s[delimIdx+1:])
	if err != nil {
		return "", 0, fmt.Errorf("could not parse line number from: \"%s\" as integer: %v", s, err)
	}

	return s[:delimIdx], line, nil
}
