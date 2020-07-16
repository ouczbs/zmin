package packet

type UPacketPool struct {
	pool chan * UPacket
}
func NewPacketPool(size TSize) *UPacketPool {
	return &UPacketPool{
		pool: make(chan *UPacket, size),
	}
}
func (stack * UPacketPool) Push(packet * UPacket){
	select {
	case stack.pool <- packet:
	default:
	}
}
func (stack * UPacketPool) Pop() * UPacket{
	var packet * UPacket
	select {
	case packet = <- stack.pool:
	default:
		packet = stack.New()
	}
	return packet
}
func (stack * UPacketPool) New() * UPacket{
	packet := &UPacket{}
	packet.Init()
	return packet
}