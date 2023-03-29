package zmessage

import (
	"github.com/ouczbs/zmin/engine/core/zclass"
	"github.com/ouczbs/zmin/engine/data/zconf"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type (
	TCode        = zconf.TCode
	TSize        = zconf.TSize
	TComponentId = zconf.TComponentId
	TMessageType = zconf.TMessageType
	TCallId      = zconf.TCallId
	TCmd         = zconf.TCmd

	IReflectMessage = protoreflect.ProtoMessage

	UStackPool = zclass.UStackPool
)

type IClientProxy interface {
}

const (
	CPoolMessageSize = zconf.CPoolMessageSize
	CPoolRequestSize = zconf.CPoolRequestSize
	CQueuePacketSize = zconf.CQueuePacketSize

	_CPacketHeadSize        = zconf.CPacketHeadSize
	_CPacketMessageTypeSize = _CPacketHeadSize + zconf.CPacketMessageTypeSize
	_CPacketMessageHeadSize = _CPacketHeadSize + zconf.CPacketMessageTypeSize + zconf.CPacketRequestTypeSize

	_CMaxPacketBuffer = zconf.CMaxPacketBuffer
	_CMinPacketBuffer = zconf.CMinPacketBuffer
)
