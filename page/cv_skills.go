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
				Title:  "Computer Science, Distributed Systems, Orchestration, Blockchain Development, Cryptography",
			},
			&component.ListItem{
				Prefix: "Programming:",
				Title:  "Go, Python",
			},
			&component.ListItem{
				Prefix: "Technologies:",
				Title:  "Protobuf, RPC, REST, Microservices, SQL",
			},
			&component.ListItem{
				Prefix: "System Administration:",
				Title:  "Docker, Kubernetes, Prometheus, Grafana, Linux",
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
