package mtmsg

import (
	"net/http"

	"github.com/forbot161602/x-lib-go/source/core/base/xbmtmsg"
)

var (
	IAV200 = xbmtmsg.NewMetaMessage(http.StatusOK,
		"IAV200", "RESTful view: OK.",
		"OK.")
)
