package component

import (
	"fmt"

	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/prop"
)

// ListItemWithURL represents the component that a list element in the CV with an URL
type ListItemWithURL struct {
	vecty.Core
	Prefix string
	Title  string
	Affix  string
	URL    string
}

func (w *ListItemWithURL) Render() vecty.ComponentOrHTML {
	return elem.ListItem(
		vecty.If(w.Prefix != "", vecty.Text(fmt.Sprintf("%s ", w.Prefix))),
		vecty.If(w.Title != "", elem.Anchor(
			vecty.Markup(vecty.Class("xp")),
			vecty.Markup(prop.Href(w.URL)),
			vecty.Text(w.Title)),
		),
		vecty.If(w.Affix != "", vecty.Text(fmt.Sprintf(", %s", w.Affix))),
	)
}

// ListItem represents the component that a list element in the CV
type ListItem struct {
	vecty.Core
	Prefix string
	Title  string
	Affix  string
}

func (w *ListItem) Render() vecty.ComponentOrHTML {
	return elem.ListItem(
		vecty.If(w.Prefix != "", elem.Strong(vecty.Text(fmt.Sprintf("%s ", w.Prefix)))),
		vecty.Text(w.Title),
		vecty.If(w.Affix != "", vecty.Text(fmt.Sprintf(", %s", w.Affix))),
	)
}
