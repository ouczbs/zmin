package dispatcher

import (
	"Zmin/component/base"
	"Zmin/engine/zconf"
	"Zmin/engine/znet"
)

type (
	UClientProxy = znet.UClientProxy
	UMessage = znet.UMessage
	URequest = znet.URequest
	UService = base.UService
	FRequestHandle = znet.FRequestHandle

	TCmd = zconf.TCmd
	TMessageType = zconf.TMessageType
)

var (
	reqHandleMaps  = make(map[TCmd]FRequestHandle)

	centerProxy * UClientProxy
)
