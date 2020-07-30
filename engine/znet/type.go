package znet

import (
	"github.com/ouczbs/Zmin/engine/zclass"
	"github.com/ouczbs/Zmin/engine/zconf"
	"github.com/ouczbs/Zmin/engine/zproto/pb"
	"google.golang.org/protobuf/reflect/protoreflect"
	"net"
)

type (
	TSize = zconf.TSize
	TCode = zconf.TCode
	TCmd = zconf.TCmd
	TEnum =zconf.TEnum

	TMessageType = zconf.TMessageType
	TCallId = zconf.TCallId

	IReflectMessage = protoreflect.ProtoMessage

	FRequestHandle func(*UClientProxy,* URequest)

	UStackPool = zclass.UStackPool
	UProperty = zclass.UProperty

	UWrapMessage = pb.WrapMessage
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
	_CPacketMessageHeadSize = _CPacketHeadSize + zconf.CPacketMessageTypeSize
)