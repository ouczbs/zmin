package login

import (
	"github.com/ouczbs/zmin/component/base"
	"github.com/ouczbs/zmin/engine/zconf"
	"github.com/ouczbs/zmin/engine/znet"
)

type (
	UClientProxy = znet.UClientProxy
	UMessage = znet.UMessage
	URequest = znet.URequest
	UService = base.UService
	UPacket = znet.UPacket
	FRequestHandle = znet.FRequestHandle

	TCmd = zconf.TCmd
	TMessageType = zconf.TMessageType
	TComponentId = zconf.TComponentId
)

var (
	reqHandleMaps  = make(map[TCmd]FRequestHandle)
	centerProxy * UClientProxy

	clientProxyMaps = make(map[TComponentId]*UClientProxy)
	gameMessageMaps = make(map[TMessageType]*UClientProxy)

)
