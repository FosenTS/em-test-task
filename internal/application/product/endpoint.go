package product

import "github.com/fasthttp/router"

type Endpoint struct {
	*Controllers
}

func NewEndpoint(controllers *Controllers) *Endpoint {
	return &Endpoint{Controllers: controllers}
}

func (e *Endpoint) RegisterRouter(r *router.Group) {
	e.Controllers.httpController.RegisterRouter(r)
}
