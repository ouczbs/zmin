package center

import (
	"Zmin/component/base"
	"Zmin/engine/zclass"
	"Zmin/engine/zconf"
	"Zmin/engine/znet"
)

type (
	UClientProxy = znet.UClientProxy
	UMessage = znet.UMessage
	URequest = znet.URequest
	UService = base.UService
	UProperty = zclass.UProperty
	FRequestHandle = znet.FRequestHandle

	TCmd = zconf.TCmd
	TEnum = zconf.TEnum
	TComponentId = zconf.TComponentId
	TSequence = zconf.TSequence
	TMessageType = zconf.TMessageType

	TProxyMap = map[TComponentId]*UClientProxy
)

var (
	reqHandleMaps  = make(map[TCmd]FRequestHandle)

	gateProxyMaps = make(TProxyMap)
	loginProxyMaps = make(TProxyMap)
	gameProxyMaps = make(TProxyMap)
	dispatcherProxyMaps = make(TProxyMap)
)
