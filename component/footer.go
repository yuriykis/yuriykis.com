package component

import (
	"fmt"
	"time"

	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/prop"
)

type Footer struct {
	vecty.Core
}

func (f *Footer) Render() vecty.ComponentOrHTML {
	return elem.Footer(
		vecty.Markup(vecty.Class("footer")),
		elem.Div(
			vecty.Text("This website is "),
			elem.Anchor(
				vecty.Markup(prop.Href("https://www.github.com/yuriykis/yuriykis.com")),
				vecty.Text("open source"),
			),
			vecty.Text(" and built with "),
			elem.Anchor(
				vecty.Markup(prop.Href("https://github.com/hexops/vecty")),
				vecty.Text("Vecty"),
			),
		),
		f.renderSignature(),
	)
}

func (p *Footer) renderSignature() *vecty.HTML {
	return elem.Preformatted(
		vecty.Text(fmt.Sprintf("© %d Yuriy Kis", time.Now().Year())),
	)
}
