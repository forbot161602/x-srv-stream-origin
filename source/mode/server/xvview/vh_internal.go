package xvview

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
