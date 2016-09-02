package events

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type identifier struct {
	Location string `json:"location"` // feature identifier
	RunID    string `json:"run_id"`
	Path     string `json:"-"` // feature path
	Line     int    `json:"-"` // feature identification line
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

type UnixTimestampMS struct {
	time.Time
}

func (t UnixTimestampMS) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.Time.UnixNano() / int64(time.Millisecond))
}

func (t *UnixTimestampMS) UnmarshalJSON(b []byte) error {
	tm, err := strconv.ParseInt(string(b), 10, 64)
	if err != nil {
		return err
	}
	t.Time = time.Unix(0, tm*int64(time.Millisecond))
	return nil
}
