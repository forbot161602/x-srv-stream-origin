package server

import "github.com/forbot161602/pbc-golang-lib/source/module/gbgin"

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
		build()
	return router
}

var mRouter *gbgin.Router

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
