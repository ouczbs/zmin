package main

import (
	"github.com/ouczbs/zmin/component/base"
	"github.com/ouczbs/zmin/component/gate"
	"github.com/ouczbs/zmin/engine/core/zlog"
)

func main() {
	zlog.Debug(base.AppPath, base.ComponentId)
	service := gate.NewGateService()
	service.Run()
}
