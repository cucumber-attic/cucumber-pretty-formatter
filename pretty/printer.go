package pretty

import (
	"fmt"
	"io"
	"strings"

	"gopkg.in/cucumber/gherkin-go.v3"
)

type printer struct {
	output io.Writer
}

func (p *printer) tags(tags []*gherkin.Tag) error {
	if len(tags) <= 0 {
		return nil
	}

	var ln string
	for _, t := range tags {
		ln += s(t.Location.Column-len(ln)) + "@" + t.Name
	}

	_, err := io.WriteString(p.output, cyan(ln)+"\n")
	return err
}

func (p *printer) header(ft *gherkin.Feature) error {
	// @TODO: print language?
	// @TODO: print comments?
	if err := p.tags(ft.Tags); err != nil {
		return err
	}

	ln := bold(white(ft.Keyword+":")) + " " + ft.Name
	if len(ft.Description) > 0 {
		for _, descLine := range strings.Split(ft.Description, "\n") {
			ln += "\n" + s(2) + strings.TrimSpace(descLine)
		}
	}
	if _, err := io.WriteString(p.output, ln+"\n"); err != nil {
		return err
	}

	return nil
}

func (p *printer) container(st *step) error {
	max := st.maxLen()
	var headline string
	loc := st.container.(*gherkin.Node).Location
	switch t := st.container.(type) {
	case *gherkin.Background:
		headline = bold(white(t.Keyword+":")) + " " + t.Name
	case *gherkin.Scenario:
		headline = bold(white(t.Keyword+":")) + " " + t.Name
	case *gherkin.ScenarioOutline:
		headline = bold(white(t.Keyword+":")) + " " + t.Name
	}

	ln := s(loc.Column) + headline
	gap := max - len(headline)
	ln += s(gap) + black(" # "+fmt.Sprintf("%s:%d", st.id.Path, loc.Line))
	_, err := io.WriteString(p.output, "\n"+ln+"\n")
	return err
}

func (p *printer) step(st *step) error {
	max := st.maxLen()
	// determine color
	var clr func(interface{}) string
	switch st.state {
	case passed:
		clr = green
	case failed:
		clr = red
	case skipped:
		clr = cyan
	case undefined:
		clr = yellow
	}
	// @TODO: print arguments in bold, need it in event?
	text := st.step.Keyword + " " + st.step.Text
	ln := s(4) + clr(text)

	// step definition ref
	if len(st.defID) > 0 {
		gap := max - len(text)
		ln += s(gap) + " " + black("# "+st.defID)
	}

	// print step
	if _, err := io.WriteString(p.output, ln+"\n"); err != nil {
		return err
	}

	// argument
	ln = ""
	switch t := st.step.Argument.(type) {
	case *gherkin.DataTable:
		ln = p.table(t) + "\n"
	case *gherkin.DocString:
		var ct string
		if len(t.ContentType) > 0 {
			ct = " " + clr(t.ContentType)
		}
		ln = s(6) + clr(t.Delimitter) + ct
		for _, part := range strings.Split(t.Content, "\n") {
			ln += "\n" + s(6) + clr(part)
		}
		ln += "\n"
		ln += s(6) + clr(t.Delimitter) + "\n"
	}
	if _, err := io.WriteString(p.output, ln); err != nil {
		return err
	}

	// summary and details
	ln = ""
	switch st.state {
	case failed:
		ln = s(6) + clr(st.summary)
		for _, d := range strings.Split(st.details, "\n") {
			ln += "\n" + s(6) + clr(d)
		}
		ln += "\n"
	case undefined:
		ln = s(6) + clr(st.summary) + "\n"
	}

	if _, err := io.WriteString(p.output, ln); err != nil {
		return err
	}
	return nil
}

func (p *printer) table(t *gherkin.DataTable) string {
	var l = p.tableSize(t)
	var cols = make([]string, len(t.Rows[0].Cells))
	var rows []string
	for _, row := range t.Rows {
		for i, cell := range row.Cells {
			cols[i] = cell.Value + s(l[i]-len(cell.Value))
		}
		rows = append(rows, s(6)+"| "+strings.Join(cols, " | ")+" |")
	}
	return strings.Join(rows, "\n")
}

func (p *printer) tableSize(tbl interface{}) []int {
	var rows []*gherkin.TableRow
	switch t := tbl.(type) {
	case *gherkin.Examples:
		rows = append(rows, t.TableHeader)
		rows = append(rows, t.TableBody...)
	case *gherkin.DataTable:
		rows = append(rows, t.Rows...)
	}

	longest := make([]int, len(rows[0].Cells))
	for _, row := range rows {
		for i, cell := range row.Cells {
			if longest[i] < len(cell.Value) {
				longest[i] = len(cell.Value)
			}
		}
	}
	return longest
}
