package main

import (
	"github.com/ouczbs/zmin/component/base"
	"github.com/ouczbs/zmin/component/version"
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
	centerConfig      FServiceConfig
)

func main() {
	runVersion()
	runLogin()
	runCenter()
	runGate()
	runDispatcher()
	for {
		time.Sleep(time.Duration(1))
	}
}
func runVersion() {
	base.ComponentId = 1
	version := version.NewVersionService()
	go version.Run()
}
func runLogin() {
	for _, config := range ServiceConfigFile.LoginList {
		path := filepath.Join(base.AppPath, "run/base/login/runlogin.go")
		runService(path, &config)
	}
}
func runCenter() {
	for _, config := range ServiceConfigFile.CenterList {
		path := filepath.Join(base.AppPath, "run/base/center/runcenter.go")
		runService(path, &config)
	}
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
