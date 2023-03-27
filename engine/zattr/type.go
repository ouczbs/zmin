package zattr

import (
	"github.com/ouczbs/zmin/engine/zconf"
	"github.com/ouczbs/zmin/engine/znet"
	"github.com/ouczbs/zmin/engine/zproto/zpb"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type (
	TCode = zconf.TCode
	TCmd = zconf.TCmd
	TEnum = zconf.TEnum

	IReflectMessage = protoreflect.ProtoMessage
	IReflectMessageType = protoreflect.MessageType
	IPbMessage = protoreflect.ProtoMessage

	UPacket = znet.UPacket
	URequest = znet.URequest
	UClientProxy = znet.UClientProxy

	UWrapMessage = zpb.WrapMessage

)
