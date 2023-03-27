package zproto

import (
	"github.com/ouczbs/zmin/engine/zconf"
	"github.com/ouczbs/zmin/engine/znet"
	"github.com/ouczbs/zmin/engine/zproto/zpb"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type (
	TCallId = zconf.TCallId
	TCode = zconf.TCode
	TCmd = zconf.TCmd
	TEnum = zconf.TEnum
	TMessageType = zconf.TMessageType

	IReflectMessage = protoreflect.ProtoMessage
	IReflectMessageType = protoreflect.MessageType
	IPbMessage = protoreflect.ProtoMessage

	UPacket = znet.UPacket
	URequest = znet.URequest
	UClientProxy = znet.UClientProxy

	UWrapMessage = zpb.WrapMessage

	FRequestHandle = znet.FRequestHandle
)

var (
	Unmarshal = proto.Unmarshal
	Marshal = proto.Marshal
)

const (
	_MT_INVALID = zconf.MT_INVALID
	_CMD_INVALID = TCmd(_MT_INVALID)
)

