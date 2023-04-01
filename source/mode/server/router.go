package server

import (
	"net/http"

	"github.com/forbot161602/pbc-golang-lib/source/module/gbgin"

	"github.com/forbot161602/pbc-stream-origin/source/mode/server/view"
)

func GetRouter() *gbgin.Router {
	if mRouter == nil {
		mRouter = newRouter()
	}
	return mRouter
}

func newRouter() *gbgin.Router {
	router := (&routerBuilder{}).
		initialize().
		setMiddlewares().
		setRouterGroup().
		build()
	return router
}

var mRouter *gbgin.Router

var routerStems = []gbgin.RouterStem{
	{
		Path:     "internal/",
		Handlers: nil,
		Leaves: []gbgin.RouterLeaf{
			{Method: http.MethodGet, Path: "info/", Handlers: view.ViewInternalInfoHandlers},
		},
		Stems: nil,
	},
}

type routerBuilder struct {
	router *gbgin.Router
}

func (builder *routerBuilder) build() *gbgin.Router {
	return builder.router
}

func (builder *routerBuilder) initialize() *routerBuilder {
	builder.router = gbgin.NewRouter()
	return builder
}

func (builder *routerBuilder) setMiddlewares() *routerBuilder {
	builder.router.UseMiddlewares()
	return builder
}

func (builder *routerBuilder) setRouterGroup() *routerBuilder {
	builder.router.SetRouterGroup(routerStems...)
	return builder
}
