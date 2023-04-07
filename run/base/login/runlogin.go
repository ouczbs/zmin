package main

import (
	"github.com/ouczbs/zmin/component/base"
	"github.com/ouczbs/zmin/component/login"
	"github.com/ouczbs/zmin/engine/core/zlog"
)

func main() {
	zlog.Debug(base.AppPath, base.ComponentId)
	service := login.NewLoginService()
	service.Run()
}
