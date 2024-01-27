package server

import (
	"net/http"

	"github.com/forbot161602/x-lib-go/source/module/xbgin"

	"github.com/forbot161602/x-srv-stream-origin/source/mode/server/view"
)

func GetRouter() *xbgin.Router {
	if mRouter == nil {
		mRouter = newRouter()
	}
	return mRouter
}

func newRouter() *xbgin.Router {
	router := (&routerBuilder{}).
		initialize().
		setMiddlewares().
		setRouterGroup().
		build()
	return router
}

var mRouter *xbgin.Router

var routerStems = []xbgin.RouterStem{
	{
		Path: "_internal/",
		Leaves: []xbgin.RouterLeaf{
			{Method: http.MethodGet, Path: "info/", Handlers: view.ViewInternalInfoHandlers},
		},
	},
}

type routerBuilder struct {
	router *xbgin.Router
}

func (builder *routerBuilder) build() *xbgin.Router {
	return builder.router
}

func (builder *routerBuilder) initialize() *routerBuilder {
	builder.router = xbgin.NewRouter()
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
