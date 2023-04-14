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

	TCmd           = zconf.TCmd
	TMessageType   = zconf.TMessageType
	TComponentId   = zconf.TComponentId
	TComponentType = zconf.TComponentType
)

var (
	ownerType     = TComponentType(zconf.COMPONENT_TYPE_VERSION)
	reqHandleMaps = make(map[TCmd]FRequestHandle)

	clientMaps = make(map[TComponentId]*UClientProxy)
)

const (
	LoginType_Password = iota
	LoginType_Email
	LoginType_Phone
)
