package preset

import (
	"fmt"

	"github.com/joho/godotenv"

	"github.com/forbot161602/x-lib-go/source/core/base/xbcfg"

	"github.com/forbot161602/x-srv-stream-origin/source/core/base/config"
)

func init() {
	if err := godotenv.Load(); err == nil {
		fmt.Println("[INFO] The .env file has been successfully loaded.")
	}
	xbcfg.SetConfig(config.GetConfig())
}
