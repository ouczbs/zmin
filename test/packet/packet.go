package packet

import "github.com/ouczbs/Zmin/engine/zconf"
var (
	packetPool = NewPacketPool(zconf.CQueuePacketSize)
	packetStackPool = NewPacketStackPool(zconf.CQueuePacketSize)
	stackPoolType = 1 // 1 -> packetChanPool 2 -> packetStackPool
	popCount = 0
	pushCount = 0
)

type UPacket struct {
	IsReleased bool
	Size 		 TSize
	bytes        []byte
}

// NewPacket allocates a new packet
func NewPacket() *UPacket {
	var packet * UPacket
	if stackPoolType == 1 {
		packet = packetPool.Pop()
	}else {
		packet =  packetStackPool.Pop()
	}
	popCount++
	packet.IsReleased = false
	return packet
}
func (packet *UPacket) Init()  {
	packet.bytes = make([]byte , 128)
}
func (packet *UPacket) Release()  {
	if packet.IsReleased {
		return
	}
	packet.IsReleased = true
	pushCount++
	if stackPoolType == 1 {
		packetPool.Push(packet)
	}else {
		packetStackPool.Push(packet)
	}
}

