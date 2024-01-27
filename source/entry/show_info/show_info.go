package show_info

import (
	"github.com/forbot161602/x-lib-go/source/core/utility/xblogger"

	"github.com/forbot161602/x-srv-stream-origin/source/core/base/config"
)

func Execute() error {
	xblogger.WithFields(xblogger.Fields{
		"gitTag":    config.GetGitTag(),
		"gitCommit": config.GetGitCommit(),
	}).Info("Log info message.")
	return nil
}
