package events

import (
	"encoding/json"
	"testing"
)

func TestMemoryUnmarshal(t *testing.T) {
	cases := []string{
		"2M",
		"0B",
		"",
		"100G200M",
	}

	for i, mem := range cases {
		data, err := json.Marshal(&struct {
			Mem string `json:"mem"`
		}{mem})

		if err != nil {
			t.Fatal(err)
		}

		exp := struct {
			Mem Memory `json:"mem"`
		}{}

		if err := json.Unmarshal(data, &exp); err != nil {
			t.Fatal(err)
		}

		if mem != exp.Mem.String() {
			t.Fatalf("expected case %d - mem %s to match %s", i, mem, exp.Mem.String())
		}
	}
}
