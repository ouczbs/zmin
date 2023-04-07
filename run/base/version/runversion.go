package main

import (
	"github.com/ouczbs/zmin/component/base"
	"github.com/ouczbs/zmin/component/version"
	"github.com/ouczbs/zmin/engine/core/zlog"
)

func main() {
	zlog.Debug(base.AppPath, base.ComponentId)
	service := version.NewVersionService()
	service.Run()
}
