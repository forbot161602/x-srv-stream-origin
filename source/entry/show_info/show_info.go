package show_info

import (
	"github.com/forbot161602/pbc-golang-lib/source/core/utility/gblogger"

	"github.com/forbot161602/pbc-stream-origin/source/core/base/config"
)

func Execute() error {
	gblogger.WithFields(gblogger.Fields{
		"gitTag":    config.GetGitTag(),
		"gitCommit": config.GetGitCommit(),
	}).Info("Log info message.")
	return nil
}
