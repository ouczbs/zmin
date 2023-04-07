package main

import (
	"github.com/ouczbs/zmin/component/base"
	"github.com/ouczbs/zmin/component/center"
	"github.com/ouczbs/zmin/engine/core/zlog"
	"os/exec"
	"path/filepath"
	"strconv"
	"time"
)

type (
	FServiceConfig = base.FServiceConfig
)

var (
	ServiceConfigFile = base.ServiceConfigFile
)

func main() {
	runCenter()
	runGate()
	runDispatcher()
	for {
		time.Sleep(time.Duration(1))
	}
}
func runCenter() {
	center := center.NewCenterService()
	go center.Run()
}
func runGate() {
	for _, config := range ServiceConfigFile.GateList {
		path := filepath.Join(base.AppPath, "run/base/gate/rungate.go")
		runService(path, &config)
	}
}
func runDispatcher() {
	for _, config := range ServiceConfigFile.DispatcherList {
		path := filepath.Join(base.AppPath, "run/base/dispatcher/rundispatcher.go")
		runService(path, &config)
	}
}
func runService(path string, config *FServiceConfig) {
	cid := strconv.Itoa(int(config.ComponentId))
	args := []string{"run", path, "-ComponentId", cid}
	cmd := exec.Command("go", args...)
	err := cmd.Start()
	if err != nil {
		zlog.Errorf("run service failed -ComponentId %d\n %s", cid, err)
	}
}
