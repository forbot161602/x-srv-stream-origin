package mtmsg

import (
	"net/http"

	"github.com/forbot161602/pbc-golang-lib/source/core/base/gbmtmsg"
)

var (
	IAV200 = gbmtmsg.NewMetaMessage(http.StatusOK,
		"IAV200", "RESTful view: OK.",
		"OK.")
)
