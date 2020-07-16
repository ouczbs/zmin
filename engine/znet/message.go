package znet

import "Zmin/engine/zconf"

var (
	messagePool = NewMessagePool(zconf.CPoolMessageSize)
)

type UMessage struct{
	Proxy 		 * UClientProxy
	MessageType  TMessageType
	Packet      * UPacket

	isReleased bool
}
func NewMessage(proxy * UClientProxy ,messageType TMessageType ,packet * UPacket )*UMessage{
	message := messagePool.Pop()
	message.Proxy = proxy
	message.MessageType = messageType
	message.Packet = packet
	message.isReleased = false
	return message
}
func (message * UMessage) Init() {

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