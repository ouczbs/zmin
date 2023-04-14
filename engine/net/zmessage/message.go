package zmessage

var (
	messagePool = NewMessagePool(CPoolMessageSize)
)

type UMessage struct {
	Proxy       IClientProxy
	MessageType TMessageType
	Cmd         TCmd
	Packet      *UPacket

	isReleased bool
}

func NewMessage(proxy IClientProxy, packet *UPacket) *UMessage {
	message := messagePool.Pop()
	message.Proxy = proxy
	message.MessageType = packet.ReadMessageType()
	message.Cmd = packet.ReadMessageCmd()
	message.Packet = packet
	message.isReleased = false
	return message
}
func (message *UMessage) Release() {
	if message.isReleased {
		return
	}
	message.isReleased = true
	message.Proxy = nil
	message.Packet.Release()
	message.Packet = nil
	messagePool.Push(message)
}
