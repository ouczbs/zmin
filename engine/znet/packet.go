package znet

import (
	"github.com/ouczbs/Zmin/engine/zconf"
	"github.com/ouczbs/Zmin/engine/zlog"
	"encoding/binary"
	"unsafe"
)

var (
	packetPool = NewPacketPool(zconf.CQueuePacketSize)
	packetBufferPools = make(map[TSize]*UPacketBufferPool)
	packetEndian = binary.LittleEndian
)

type UPacket struct {
	IsReleased bool
	Size 		 TSize

	byteSize     TSize
	bytes        []byte
}

// NewPacket allocates a new packet
func NewPacket() *UPacket {
	var packet * UPacket
	packet = packetPool.Pop()
	packet.IsReleased = false
	return packet
}
func (packet *UPacket) Init()  {
	packet.MakeSize(_CMinPacketBuffer)
	packet.Size = _CPacketHeadSize
}
func (packet *UPacket) Release()  {
	if packet.IsReleased {
		return
	}
	packet.IsReleased = true
	packetBufferPool := _makePacketBufferPool(packet.byteSize)
	packetBufferPool.Push(packet.bytes)
	packetPool.Push(packet)
}
func _makePacketBufferPool(ByteSize TSize)*UPacketBufferPool{
	packetBufferPool := packetBufferPools[ByteSize]
	if packetBufferPool == nil {
		packetBufferPool = NewPacketBufferPool(ByteSize)
		packetBufferPools[ByteSize] = packetBufferPool
	}
	return packetBufferPool
}
func (packet *UPacket) MakeSize(size TSize){
	packet.AssureBytesSize(size)
}
func (packet *UPacket) SetSize(size TSize){
	packet.Size = size
	packet.AssureBytesSize(size)
}
func (packet *UPacket) AssureBytesSize(size TSize){
	if size <= packet.byteSize{
		return
	}
	beforeBytesSize := packet.byteSize
	beforeBytes := packet.bytes
	byteSize := beforeBytesSize * 2
	if byteSize < _CMinPacketBuffer {
		byteSize = _CMinPacketBuffer
	}else if byteSize > _CMaxPacketBuffer {
		zlog.Panic("Packet length out of max packet buffer")
	}
	packetBufferPool := _makePacketBufferPool(byteSize)
	packet.bytes = packetBufferPool.Pop()
	packet.byteSize = byteSize
	if beforeBytesSize >= _CMinPacketBuffer {
		curPacketBufferPool := _makePacketBufferPool(beforeBytesSize)
		curPacketBufferPool.Push(beforeBytes)
		copy(packet.bytes[:packet.Size], beforeBytes[:packet.Size])
	}
}
func (packet *UPacket) ReadSize() TSize {
	return TSize(packetEndian.Uint16(packet.bytes[:_CPacketHeadSize]))
}
func (packet *UPacket) WriteSize(size TSize) {
	*(*uint16)(unsafe.Pointer(&packet.bytes[0])) = uint16(size)
}
func (packet *UPacket) ReadMessageType() TMessageType {
	return packetEndian.Uint16(packet.bytes[_CPacketHeadSize:_CPacketMessageHeadSize])
}
func (packet *UPacket) MessagePayload() []byte {
	return packet.bytes[_CPacketMessageHeadSize:]
}
func (packet *UPacket) WriteMessageType(messageType TMessageType) {
	packetEndian.PutUint16(packet.bytes[_CPacketHeadSize:_CPacketMessageHeadSize] , messageType)
	packet.Size = _CPacketMessageHeadSize
}
func (packet *UPacket) AppendBytes(buf []byte) {
	size := packet.Size
	bufSize := TSize(len(buf))
	packet.SetSize(bufSize + size)
	copy(packet.bytes[size:size+bufSize], buf)
}

