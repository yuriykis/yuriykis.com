package page

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/yuriykis/yuriykis.com/component"
)

func (p *CV) renderSkills() *vecty.HTML {
	return elem.Span(
		vecty.Markup(vecty.Class("cv-entry")),
		elem.UnorderedList(
			&component.ListItem{
				Prefix: "General:",
				Title:  "Computer Science, Distibuted Systems, Blockchain Development, Cryptography",
			},
			&component.ListItem{
				Prefix: "Programming:",
				Title:  "Go, Python",
			},
			&component.ListItem{
				Prefix: "Tech:",
				Title:  "Protobuf, RPC, REST, Microservices",
			},
			&component.ListItem{
				Prefix: "OS:",
				Title:  "GNU/Linux",
			},
			&component.ListItem{
				Prefix: "Business:",
				Title:  "Agile, SCRUM, Kanban, UML",
			},
			&component.ListItem{
				Prefix: "Languages:",
				Title:  "Polish (Native), Ukrainian (Native), Russian (Fluent), English (Fluent)",
			},
		),
	)
}
