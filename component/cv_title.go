package component

import (
	"fmt"

	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
)

type Title struct {
	vecty.Core
	Title string
}

func (t *Title) Render() vecty.ComponentOrHTML {
	return elem.Heading2(
		vecty.Markup(vecty.Class("cv-title")),
		vecty.Text(fmt.Sprintf("— %s", t.Title)),
	)
}
