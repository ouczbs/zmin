package base

import (
	"github.com/ouczbs/Zmin/engine/zclass"
	"github.com/ouczbs/Zmin/engine/zconf"
	"github.com/ouczbs/Zmin/engine/znet"
)

type (
	UClientProxy = znet.UClientProxy
	UMessage = znet.UMessage
	URequest = znet.URequest
	UProperty = zclass.UProperty
	UServiceConfig = zconf.UServiceConfig
	FRequestHandle = znet.FRequestHandle

	IService = znet.IService


	TCmd = zconf.TCmd
	TEnum = zconf.TEnum
)
