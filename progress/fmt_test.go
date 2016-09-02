package progress

import "testing"

func TestShouldMatchSupportedProtocol(t *testing.T) {
	type Case struct {
		vers    string
		matches bool
	}

	var cases = []Case{
		{"0.1.0", true},
		{"0.1.3", true},
		{"0.1.124", true},
		{"1.1.5", false},
	}

	for n, c := range cases {
		if supportedProtocol.MatchString(c.vers) != c.matches {
			t.Fatalf("case %d failed, expected %s to be %+v", n, c.vers, c.matches)
		}
	}
}
