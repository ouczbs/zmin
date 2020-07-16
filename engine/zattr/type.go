package zattr

import (
	"Zmin/engine/zconf"
	"Zmin/engine/znet"
	"Zmin/engine/zproto/pb"
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

	UWrapMessage = pb.WrapMessage

)
