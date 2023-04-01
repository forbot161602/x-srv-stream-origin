package view

import (
	"github.com/forbot161602/pbc-golang-lib/source/module/gbgin"
)

type (
	Handler  = gbgin.Handler
	Context  = gbgin.Context
	KongFlow = gbgin.KongFlow
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
