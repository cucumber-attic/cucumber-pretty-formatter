package formatter

import (
	"fmt"

	"github.com/DATA-DOG/godog"
)

func passing() error {
	return nil
}

func failing() error {
	return fmt.Errorf("step failed error")
}

func pending() error {
	return godog.ErrPending
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^passing$`, passing)
	s.Step(`^failing$`, failing)
	s.Step(`^pending$`, pending)
}
