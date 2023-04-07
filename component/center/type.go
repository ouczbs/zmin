package center

import (
	"github.com/ouczbs/zmin/component/base"
	"github.com/ouczbs/zmin/engine/core/zclass"
	"github.com/ouczbs/zmin/engine/data/zconf"
	"github.com/ouczbs/zmin/engine/net/zmessage"
	"github.com/ouczbs/zmin/engine/net/znet"
)

type (
	UClientProxy   = znet.UClientProxy
	UMessage       = zmessage.UMessage
	URequest       = zmessage.URequest
	UService       = base.UService
	UProperty      = zclass.UProperty
	FRequestHandle = znet.FRequestHandle

	TCmd         = zconf.TCmd
	TEnum        = zconf.TEnum
	TComponentId = zconf.TComponentId
	TSequence    = zconf.TSequence
	TMessageType = zconf.TMessageType

	TProxyMap = map[TComponentId]*UClientProxy
)

var (
	reqHandleMaps = make(map[TCmd]FRequestHandle)

	versionProxy *UClientProxy

	gateProxyMaps       = make(TProxyMap)
	loginProxyMaps      = make(TProxyMap)
	gameProxyMaps       = make(TProxyMap)
	dispatcherProxyMaps = make(TProxyMap)
)
