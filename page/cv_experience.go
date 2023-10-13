package page

import (
	"time"

	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/yuriykis/yuriykis.com/component"
)

func (p *CV) renderExperience() *vecty.HTML {
	return elem.Span(
		vecty.Markup(vecty.Class("cv-entry")),
		elem.UnorderedList(
			&component.Experience{
				BeginDate:   time.Date(2021, time.February, 1, 0, 0, 0, 0, time.UTC),
				Location:    "Hybrid",
				JobTitle:    "Senior Systems Engineer",
				Company:     "Lufthansa Systems Poland",
				Description: "Systems Engineer for the several airlines within the Lufthansa Group.",
			},
			&component.Experience{
				BeginDate:   time.Date(2022, time.March, 1, 0, 0, 0, 0, time.UTC),
				Location:    "Remote",
				JobTitle:    "Software Engineer (Part-time)",
				Company:     "Seaeasy",
				Description: "Software Engineer for the Seaeasy service.",
			},
			&component.Experience{
				BeginDate:   time.Date(2020, time.July, 1, 0, 0, 0, 0, time.UTC),
				EndDate:     time.Date(2021, time.February, 1, 0, 0, 0, 0, time.UTC),
				Location:    "Gdynia, Poland",
				JobTitle:    "Software Developer",
				Company:     "General Engineering Solutions",
				Description: "Software Engineer for the GES Cloud Computing service.",
			},
			&component.Experience{
				BeginDate:   time.Date(2018, time.February, 1, 0, 0, 0, 0, time.UTC),
				EndDate:     time.Date(2020, time.June, 1, 0, 0, 0, 0, time.UTC),
				Location:    "Gdańsk, Poland",
				JobTitle:    "Systems Engineer",
				Company:     "Call Center Poland",
				Description: "",
			},
		),
	)
}
