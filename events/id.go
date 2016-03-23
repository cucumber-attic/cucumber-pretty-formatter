package events

import "github.com/cucumber/cucumber-pretty-formatter/events/location"

type id struct {
	ID    string // feature identifier
	Suite string // on which suite the feature runs in
	Path  string // feature path
	Line  int    // feature identification line
}

func (i *id) parseID() error {
	loc, err := location.From(i.ID)
	if err != nil {
		return err
	}
	i.Path = loc.Path
	i.Line = loc.Line
	return nil
}
