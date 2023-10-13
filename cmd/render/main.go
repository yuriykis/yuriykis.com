package main

import (
	"github.com/hexops/vecty"
	"github.com/yuriykis/yuriykis.com/page"
)

func main() {
	vecty.SetTitle("Yuriy Kis")
	vecty.RenderBody(&page.Body{})
}
