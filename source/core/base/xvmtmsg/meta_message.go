package xvmtmsg

import (
	"net/http"

	"github.com/starryck/x-lib-go/source/core/base/xbmtmsg"
)

var (
	IAV200 = xbmtmsg.NewMetaMessage(http.StatusOK,
		"IAV200", "RESTful view: OK.",
		"OK.")
)
