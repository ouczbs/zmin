package base

import (
	"Zmin/engine/zclass"
	"Zmin/engine/zconf"
	"Zmin/engine/znet"
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
