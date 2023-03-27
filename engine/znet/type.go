package znet

import (
	"github.com/ouczbs/zmin/engine/zclass"
	"github.com/ouczbs/zmin/engine/zconf"
	"github.com/ouczbs/zmin/engine/zproto/zpb"
	"google.golang.org/protobuf/reflect/protoreflect"
	"net"
)

type (
	TSequence = zconf.TSequence
	TSize = zconf.TSize
	TCode = zconf.TCode
	TCmd = zconf.TCmd
	TEnum =zconf.TEnum
	TComponentId = zconf.TComponentId

	TMessageType = zconf.TMessageType
	TCallId = zconf.TCallId

	IReflectMessage = protoreflect.ProtoMessage

	FRequestHandle func(*UClientProxy,* URequest)

	UStackPool = zclass.UStackPool
	UProperty = zclass.UProperty

	UWrapMessage = zpb.WrapMessage
)
type IService interface {
	NewTcpConnection(net.Conn)
	ClientDisconnect(*UClientProxy)
	RecvMessage(*UMessage)
	GetRequestHandle(TCmd)FRequestHandle
}
const (
	_CMaxPacketBuffer = zconf.CMaxPacketBuffer
	_CMinPacketBuffer = zconf.CMinPacketBuffer

	_CPacketHeadSize = zconf.CPacketHeadSize
	_CPacketMessageTypeSize = _CPacketHeadSize + zconf.CPacketMessageTypeSize
	_CPacketMessageHeadSize = _CPacketHeadSize + zconf.CPacketMessageTypeSize + zconf.CPacketRequestTypeSize
)