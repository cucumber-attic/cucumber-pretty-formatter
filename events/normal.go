// +build !testing

package events

import (
	"strconv"
	"strings"
	"time"
)

func (t *Timestamp) UnmarshalJSON(b []byte) error {
	i, err := strconv.ParseInt(string(b), 10, 64)
	if err != nil {
		return err
	}
	*t = Timestamp(i * int64(time.Millisecond))
	return nil
}

func (m *Memory) UnmarshalJSON(b []byte) error {
	*m = Memory(strings.Trim(string(b), `"`))
	return nil
}
