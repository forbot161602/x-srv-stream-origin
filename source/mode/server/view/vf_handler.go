package view

import (
	"github.com/forbot161602/x-lib-go/source/module/xbgin"
)

type (
	Flow     = xbgin.Flow
	KongFlow = xbgin.KongFlow
	Handler  = xbgin.Handler
	Context  = xbgin.Context
)

func InternalRequestHandler(ctx *Context) {
	flow := &InternalRequestHandlerFlow{}
	flow.Initiate(ctx)
	if !flow.IsInternalRequest() {
		flow.SetNotFoundError()
	}
}

type InternalRequestHandlerFlow struct {
	KongFlow
}
