package dispatcher

import (
	"github.com/ouczbs/Zmin/component/base"
	"github.com/ouczbs/Zmin/engine/zconf"
	"github.com/ouczbs/Zmin/engine/znet"
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
