package xvinfo

import (
	"github.com/forbot161602/x-lib-go/source/core/utility/xblogger"

	"github.com/forbot161602/x-srv-stream-origin/source/core/base/xvcfg"
)

func Execute() error {
	xblogger.WithFields(xblogger.Fields{
		"gitTag":    xvcfg.GetGitTag(),
		"gitCommit": xvcfg.GetGitCommit(),
	}).Info("Log info message.")
	return nil
}
