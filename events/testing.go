// +build testing

package events

func (t *Timestamp) UnmarshalJSON(b []byte) error {
	*t = Timestamp(0)
	return nil
}

func (m *Memory) UnmarshalJSON(b []byte) error {
	*m = Memory("0B")
	return nil
}
