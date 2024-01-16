package httpController

import (
	"github.com/fasthttp/router"
)

type HTTPHandler interface {
	RegisterRouter(*router.Group)
}

type httpController struct {
	bioHandler HTTPHandler
}

func NewHttpController(bioHandler HTTPHandler) HTTPHandler {
	return &httpController{bioHandler: bioHandler}
}

func (hC *httpController) RegisterRouter(r *router.Group) {
	bioGroup := r.Group("/bio")

	hC.bioHandler.RegisterRouter(bioGroup)
}

const (
	JsonType = "application/json"
)
