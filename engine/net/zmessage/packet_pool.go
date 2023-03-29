package zmessage

type UPacketPool struct {
	*UStackPool
}

func NewPacketPool(size TSize) *UPacketPool {
	stack := &UPacketPool{
		&UStackPool{
			Size: size,
		},
	}
	stack.Init()
	return stack
}
func (stack *UPacketPool) Pop() *UPacket {
	object := stack.UStackPool.Pop()
	if object == nil {
		return stack.New()
	}
	return object.(*UPacket)
}
func (stack *UPacketPool) New() *UPacket {
	return &UPacket{}
}

type UPacketBufferPool struct {
	*UStackPool
	bufferSize TSize
}

func (stack *UPacketBufferPool) Pop() []byte {
	object := stack.UStackPool.Pop()
	if object == nil {
		return stack.New()
	}
	return object.([]byte)
}
func (stack *UPacketBufferPool) New() []byte {
	return make([]byte, stack.bufferSize)
}
func NewPacketBufferPool(bufferSize TSize) *UPacketBufferPool {
	stack := &UPacketBufferPool{
		&UStackPool{Size: 1},
		bufferSize,
	}
	stack.Init()
	return stack
}
