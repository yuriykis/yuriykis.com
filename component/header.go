package component

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/prop"
)

type Header struct {
	vecty.Core
	Title string
}

func (h *Header) Render() vecty.ComponentOrHTML {
	return elem.Header(
		elem.Heading1(
			vecty.Markup(vecty.Class("logo")),
			elem.Anchor(
				vecty.Markup(prop.Href("cv")),
				vecty.Text(h.Title)),
		),
	)
}
