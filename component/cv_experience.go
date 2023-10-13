package component

import (
	"fmt"
	"time"

	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
)

type Experience struct {
	vecty.Core
	BeginDate   time.Time
	EndDate     time.Time
	Location    string
	JobTitle    string
	Company     string
	Description string
}

func (e *Experience) Render() vecty.ComponentOrHTML {
	return elem.Span(
		elem.ListItem(
			e.renderDate(),
			vecty.If(e.JobTitle != "", elem.Strong(vecty.Text(e.JobTitle))),
			vecty.If(e.JobTitle == "", elem.Strong(vecty.Text(e.Company))),
			vecty.If(
				e.JobTitle != "",
				vecty.Text(fmt.Sprintf(" at %s, %s", e.Company, e.Location)),
			),
			vecty.If(e.Description != "", elem.BlockQuote(vecty.Text(e.Description))),
		),
	)
}

func (e *Experience) renderDate() *vecty.HTML {
	beginDate := formatDate(e.BeginDate)
	endDate := formatDate(e.EndDate)
	if endDate == "" {
		return elem.Span(
			vecty.Text(fmt.Sprintf("%s - ", beginDate)),
			elem.Emphasis(vecty.Text("Present")),
			vecty.Text(", "),
		)
	}
	return vecty.Text(fmt.Sprintf("%s - %s, ", beginDate, endDate))
}

func formatDate(d time.Time) string {
	if d == (time.Time{}) {
		return ""
	}
	return fmt.Sprintf("%s %d", d.Month().String()[:3], d.Year())
}
