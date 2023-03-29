package login

import (
	"github.com/ouczbs/zmin/component/base"
	"github.com/ouczbs/zmin/engine/data/zconf"
	"github.com/ouczbs/zmin/engine/net/zmessage"
	"github.com/ouczbs/zmin/engine/net/znet"
)

type (
	UClientProxy   = znet.UClientProxy
	UMessage       = zmessage.UMessage
	URequest       = zmessage.URequest
	UService       = base.UService
	UPacket        = zmessage.UPacket
	FRequestHandle = znet.FRequestHandle

	TCmd         = zconf.TCmd
	TMessageType = zconf.TMessageType
	TComponentId = zconf.TComponentId
)

var (
	reqHandleMaps = make(map[TCmd]FRequestHandle)
	centerProxy   *UClientProxy

	clientProxyMaps = make(map[TComponentId]*UClientProxy)
	gameMessageMaps = make(map[TMessageType]*UClientProxy)
)
