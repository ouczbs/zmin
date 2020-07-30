package zproto

import (
	"github.com/ouczbs/Zmin/engine/zconf"
	"github.com/ouczbs/Zmin/engine/znet"
	"github.com/ouczbs/Zmin/engine/zproto/pb"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type (
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

	UWrapMessage = pb.WrapMessage

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

