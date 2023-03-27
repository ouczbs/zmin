package base

import (
	"github.com/ouczbs/zmin/engine/zclass"
	"github.com/ouczbs/zmin/engine/zconf"
	"github.com/ouczbs/zmin/engine/znet"
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
