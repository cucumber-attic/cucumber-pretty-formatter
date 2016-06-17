package fixtures

import (
	"fmt"

	"github.com/DATA-DOG/godog"
)

func passing() error {
	return nil
}

func failing() error {
	return fmt.Errorf("failed")
}

func skipping() error {
	return nil
}

func MainContext(s *godog.Suite) {
	s.Step(`^passing$`, passing)
	s.Step(`^failing$`, failing)
	s.Step(`^skipping$`, skipping)
}
