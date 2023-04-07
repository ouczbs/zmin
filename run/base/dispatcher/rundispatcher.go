package main

import (
	"github.com/ouczbs/zmin/component/base"
	"github.com/ouczbs/zmin/component/dispatcher"
	"github.com/ouczbs/zmin/engine/core/zlog"
)

func main() {
	zlog.Debug(base.AppPath, base.ComponentId)
	service := dispatcher.NewDispatcherService()
	service.Run()
}
