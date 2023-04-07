package main

import (
	"github.com/ouczbs/zmin/component/base"
	"github.com/ouczbs/zmin/component/center"
	"github.com/ouczbs/zmin/engine/core/zlog"
)

func main() {
	zlog.Debug(base.AppPath, base.ComponentId)
	service := center.NewCenterService()
	service.Run()
}
