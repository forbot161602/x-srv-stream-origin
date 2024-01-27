package run_server

import (
	"fmt"
	"net/http"

	"github.com/forbot161602/x-lib-go/source/utility/xbspvs"

	"github.com/forbot161602/x-srv-stream-origin/source/core/base/config"
	"github.com/forbot161602/x-srv-stream-origin/source/mode/server"
)

func Execute() error {
	supervisor := xbspvs.GetSupervisor(nil)
	supervisor.Handle(&ServerProcess{})
	supervisor.RunForever()
	return nil
}

type ServerProcess struct {
	xbspvs.ServerProcess
}

func (process *ServerProcess) Setup() error {
	process.SetServer(&http.Server{
		Addr:    fmt.Sprintf(":%d", config.GetServicePort()),
		Handler: server.GetRouter().GetEngine(),
	})
	return nil
}
