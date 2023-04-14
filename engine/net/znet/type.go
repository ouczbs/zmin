package znet

import (
	"github.com/ouczbs/zmin/engine/core/zclass"
	"github.com/ouczbs/zmin/engine/data/zconf"
	"github.com/ouczbs/zmin/engine/net/zmessage"
	"google.golang.org/protobuf/reflect/protoreflect"
	"net"
	"time"
)

type (
	TSequence    = zconf.TSequence
	TSize        = zconf.TSize
	TCmd         = zconf.TCmd
	TEnum        = zconf.TEnum
	TComponentId = zconf.TComponentId

	TMessageType = zconf.TMessageType
	TCallId      = zconf.TCallId

	IReflectMessage = protoreflect.ProtoMessage

	FRequestHandle func(*UClientProxy, *zmessage.URequest)

	UStackPool = zclass.UStackPool
	UProperty  = zclass.UProperty
)
type IService interface {
	NewTcpConnection(net.Conn)
	ClientDisconnect(*UClientProxy)
	RecvMessage(*zmessage.UMessage)
	GetRequestHandle(TCmd) FRequestHandle
}

const (
	_RESTART_TCP_SERVER_INTERVAL = 3 * time.Second
	_RESTART_UDP_SERVER_INTERVAL = 3 * time.Second

	_CPacketHeadSize = zconf.CPacketHeadSize
)
