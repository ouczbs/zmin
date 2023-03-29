package zproto

import (
	"github.com/ouczbs/zmin/engine/data/zconf"
	"github.com/ouczbs/zmin/engine/net/zmessage"
	"github.com/ouczbs/zmin/engine/net/znet"
	"github.com/ouczbs/zmin/engine/sync/zpb"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type (
	TCallId      = zconf.TCallId
	TCode        = zconf.TCode
	TCmd         = zconf.TCmd
	TEnum        = zconf.TEnum
	TMessageType = zconf.TMessageType

	IReflectMessage     = protoreflect.ProtoMessage
	IReflectMessageType = protoreflect.MessageType
	IPbMessage          = protoreflect.ProtoMessage

	UPacket      = zmessage.UPacket
	URequest     = zmessage.URequest
	UClientProxy = znet.UClientProxy

	UWrapMessage = zpb.WrapMessage

	FRequestHandle = znet.FRequestHandle
)

var (
	Unmarshal = proto.Unmarshal
	Marshal   = proto.Marshal
)

const (
	_MT_INVALID  = zconf.MT_INVALID
	_CMD_INVALID = TCmd(_MT_INVALID)
)
