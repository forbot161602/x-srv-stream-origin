package xvview

import (
	"github.com/forbot161602/x-srv-stream-origin/source/core/base/xvcfg"
	"github.com/forbot161602/x-srv-stream-origin/source/core/base/xvmtmsg"
)

var ViewInternalInfoHandlers = []Handler{
	InternalRequestHandler,
	ViewInternalInfoHandler,
}

func ViewInternalInfoHandler(ctx *Context) {
	flow := &ViewInternalInfoHandlerFlow{}
	flow.Initiate(ctx)
	flow.SetResult()
}

type ViewInternalInfoHandlerFlow struct {
	APISIXFlow
}

func (flow *ViewInternalInfoHandlerFlow) SetResult() {
	if flow.HasError() {
		return
	}

	info := xvcfg.GetConfig()
	flow.RespondJSON(xvmtmsg.IAV200, info, nil)
	return
}
