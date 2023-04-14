package gate

import (
	"github.com/ouczbs/zmin/component/base"
	"github.com/ouczbs/zmin/engine/core/zutil"
	"github.com/ouczbs/zmin/engine/data/zconf"
	"github.com/ouczbs/zmin/engine/net/zmessage"
	"github.com/ouczbs/zmin/engine/net/znet"
)

type (
	UClientProxy   = znet.UClientProxy
	UMessage       = zmessage.UMessage
	URequest       = zmessage.URequest
	UService       = base.UService
	FRequestHandle = znet.FRequestHandle
	UPacket        = zmessage.UPacket
	TCmd           = zconf.TCmd
	TMessageType   = zconf.TMessageType
	TComponentId   = zconf.TComponentId
	TComponentType = zconf.TComponentType
)

var (
	sequence      = zutil.NewSequence()
	ownerType     = TComponentType(zconf.COMPONENT_TYPE_VERSION)
	reqHandleMaps = make(map[TCmd]FRequestHandle)

	clientMaps = make(map[TComponentId]*UClientProxy)

	dispatcherMaps = make(map[TComponentId]bool)
	dispatcherList []*UClientProxy
	dispatcherSize int32 = 1
)
