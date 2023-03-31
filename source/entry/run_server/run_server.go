package run_server

import (
	"fmt"
	"net/http"

	"github.com/forbot161602/pbc-golang-lib/source/utility/gbspvs"

	"github.com/forbot161602/pbc-stream-origin/source/core/base/config"
	"github.com/forbot161602/pbc-stream-origin/source/mode/server"
)

func Execute() error {
	supervisor := gbspvs.GetSupervisor(nil)
	supervisor.Handle(&ServerProcess{})
	supervisor.RunForever()
	return nil
}

type ServerProcess struct {
	gbspvs.ServerProcess
}

func (process *ServerProcess) Setup() error {
	process.SetServer(&http.Server{
		Addr:    fmt.Sprintf(":%d", config.GetServicePort()),
		Handler: server.GetRouter().GetEngine(),
	})
	return nil
}
