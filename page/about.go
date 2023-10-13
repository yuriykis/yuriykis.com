package page

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/prop"
	"github.com/yuriykis/yuriykis.com/component"
)

type About struct {
	vecty.Core
}

func (a *About) Render() vecty.ComponentOrHTML {
	return elem.Div(
		&component.Header{
			Title: "Yuriy Kis",
		},
		elem.Section(
			vecty.Markup(vecty.Class("main")),
			elem.Heading1(
				vecty.Text(
					"Software and Systems Engineer, currently working at Luftansa Systems Poland.",
				),
			),
			elem.Heading1(
				vecty.Text(
					"My main interests are in distributed systems, software architecture and consensus algorithms.",
				),
			),
			elem.Heading1(
				vecty.Text(
					"You can find my projects and open source contributions on my ",
				),
				elem.Anchor(
					vecty.Markup(prop.Href("https://www.github.com/yuriykis/")),
					vecty.Text("GitHub  profile."),
				),
			),
			a.renderContact(),
		),
		&component.Footer{},
	)
}

func (p *About) renderContact() *vecty.HTML {
	return elem.Heading1(
		vecty.Text("If you want to contact me, please send me an email at:"),
		elem.Preformatted(
			vecty.Markup(vecty.Class("email")),
			vecty.Text("yurakis9@gmail.com"),
		),
		vecty.Text("or find me on "),
		elem.Anchor(
			vecty.Markup(prop.Href("https://www.linkedin.com/in/yuriy-kis/")),
			vecty.Text("LinkedIn profile."),
		),
	)
}
