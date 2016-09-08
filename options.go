package formatter

import "io"

type Options struct {
	NoColors         bool
	EventInputStream io.ReadCloser
	Formats          []string
}
