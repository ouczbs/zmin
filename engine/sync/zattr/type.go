package zattr

import (
	"github.com/ouczbs/zmin/engine/data/zconf"
	"github.com/ouczbs/zmin/engine/net/zmessage"
	"github.com/ouczbs/zmin/engine/net/znet"
	"github.com/ouczbs/zmin/engine/sync/zpb"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type (
	TCode = zconf.TCode
	TCmd  = zconf.TCmd
	TEnum = zconf.TEnum

	IReflectMessage     = protoreflect.ProtoMessage
	IReflectMessageType = protoreflect.MessageType
	IPbMessage          = protoreflect.ProtoMessage

	UPacket      = zmessage.UPacket
	URequest     = zmessage.URequest
	UClientProxy = znet.UClientProxy

	UWrapMessage = zpb.WrapMessage
)
