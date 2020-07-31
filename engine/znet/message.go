package znet

import "github.com/ouczbs/Zmin/engine/zconf"

var (
	messagePool = NewMessagePool(zconf.CPoolMessageSize)
)

type UMessage struct{
	Proxy 		 * UClientProxy
	MessageType  TMessageType
	Cmd 	 	 	TCmd
	Packet      * UPacket

	isReleased bool
}
func NewMessage(proxy * UClientProxy ,packet * UPacket )*UMessage{
	message := messagePool.Pop()
	message.Proxy = proxy
	message.MessageType = packet.ReadMessageType()
	message.Cmd = packet.ReadMessageCmd()
	message.Packet = packet
	message.isReleased = false
	return message
}
func (message * UMessage) Release() {
	if message.isReleased {
		return
	}
	message.isReleased = true
	message.Proxy = nil
	message.Packet.Release()
	messagePool.Push(message)
}