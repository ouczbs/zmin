package packet

import (
	"Zmin/engine/zclass"
)

type UPacketStackPool struct {
	* zclass.UStackPool
}
func NewPacketStackPool(size TSize) *UPacketStackPool {
	stack := &UPacketStackPool{
		&zclass.UStackPool{
			Size: size,
			IsStatic: true,
		},
	}
	stack.Init()
	return stack
}
func (stack * UPacketStackPool) Pop() * UPacket{
	object := stack.UStackPool.Pop()
	if object == nil {
		return stack.New()
	}
	return object.(*UPacket)
}
func (stack * UPacketStackPool) New() * UPacket{
	packet := &UPacket{}
	packet.Init()
	return packet
}