package preset

import (
	"fmt"

	"github.com/joho/godotenv"

	"github.com/forbot161602/pbc-golang-lib/source/core/base/gbcfg"

	"github.com/forbot161602/pbc-stream-origin/source/core/base/config"
)

func init() {
	if err := godotenv.Load(); err == nil {
		fmt.Println("[INFO] The .env file has been successfully loaded.")
	}
	gbcfg.SetConfig(config.GetConfig())
}
