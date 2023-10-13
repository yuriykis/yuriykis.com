package page

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	router "marwan.io/vecty-router"
)

type Body struct {
	vecty.Core
}

func (b *Body) Render() vecty.ComponentOrHTML {
	return elem.Body(
		router.NewRoute(
			"/", &About{}, router.NewRouteOpts{
				ExactMatch: true,
			},
		),
		router.NewRoute(
			"/cv", &CV{}, router.NewRouteOpts{},
		),
		router.NotFoundHandler(&About{}),
	)
}
