package page

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/yuriykis/yuriykis.com/component"
)

type CV struct {
	vecty.Core
}

func (p *CV) Render() vecty.ComponentOrHTML {
	return elem.Div(
		&component.Header{
			Title: "Yuriy Kis",
		},
		elem.Section(
			vecty.Markup(vecty.Class("main")),
			// Experience section
			&component.Title{Title: "Experience"},
			p.renderExperience(),
			// // Side Projects section
			// &component.Title{Title: "Side Projects"},
			// p.renderProjects(),
			// Skills section
			&component.Title{Title: "Skills"},
			p.renderSkills(),
			// Education section
			&component.Title{Title: "Education"},
			p.renderEducation(),
			// // Certifications section
			// &component.Title{Title: "Certifications"},
			// p.renderCertification(),
			// // Interest section
			// &component.Title{Title: "Interests"},
			// p.renderInterest(),
		),
		&component.Footer{},
	)
}
