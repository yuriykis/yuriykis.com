package page

import (
	"time"

	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/yuriykis/yuriykis.com/component"
)

func (p *CV) renderEducation() *vecty.HTML {
	return elem.Span(
		vecty.Markup(vecty.Class("cv-entry")),
		elem.UnorderedList(
			&component.Experience{
				BeginDate:   time.Date(2021, time.February, 1, 0, 0, 0, 0, time.UTC),
				EndDate:     time.Date(2022, time.October, 1, 0, 0, 0, 0, time.UTC),
				Location:    "Gdańsk, Poland",
				JobTitle:    "Master of Science in Computer Science",
				Company:     "Gdańsk University of Technology",
				Description: "",
			},
		),
		elem.UnorderedList(
			&component.Experience{
				BeginDate:   time.Date(2016, time.October, 1, 0, 0, 0, 0, time.UTC),
				EndDate:     time.Date(2021, time.February, 1, 0, 0, 0, 0, time.UTC),
				Location:    "Gdańsk, Poland",
				JobTitle:    "Bachelor of Science in Computer Science",
				Company:     "Gdańsk University of Technology",
				Description: "",
			},
		),
	)
}
