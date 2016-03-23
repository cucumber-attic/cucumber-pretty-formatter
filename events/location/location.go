package location

import (
	"fmt"
	"strconv"
	"strings"
)

type Location struct {
	Line int
	Path string
}

func From(s string) (*Location, error) {
	delimIdx := strings.LastIndex(s, ":")
	if delimIdx == -1 {
		return nil, fmt.Errorf("could not parse location, line delimiter not found from: %s", s)
	}

	line, err := strconv.Atoi(s[delimIdx+1:])
	if err != nil {
		return nil, fmt.Errorf("could not parse line number from: \"%s\" as integer: %s", s, err)
	}

	return &Location{
		Line: line,
		Path: s[:delimIdx],
	}, nil
}
