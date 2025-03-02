package xvinfo

import (
	"github.com/starryck/strk-tc-x-lib-go/source/core/utility/xblogger"

	"github.com/forbot161602/x-srv-stream-origin/source/core/base/xvcfg"
)

func Execute() error {
	xblogger.WithFields(xblogger.Fields{
		"config": xvcfg.GetConfig(),
	}).Info("Log info message.")
	return nil
}
