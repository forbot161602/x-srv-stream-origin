package view

import (
	"github.com/forbot161602/x-srv-stream-origin/source/core/base/config"
	"github.com/forbot161602/x-srv-stream-origin/source/core/base/mtmsg"
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
	KongFlow
}

func (flow *ViewInternalInfoHandlerFlow) SetResult() {
	if flow.HasError() {
		return
	}

	info := config.GetConfig()
	flow.RespondJSON(mtmsg.IAV200, info, nil)
	return
}
