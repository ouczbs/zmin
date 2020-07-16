package znet

type UMessagePool struct {
	* UStackPool
}
func NewMessagePool(size TSize) *UMessagePool {
	stack := &UMessagePool{
		&UStackPool{
			Size: size,
		},
	}
	stack.Init()
	return stack
}
func (stack * UMessagePool) Pop() * UMessage{
	object := stack.UStackPool.Pop()
	if object == nil {
		return stack.New()
	}
	return object.(*UMessage)
}
func (stack * UMessagePool) New() * UMessage{
	packet := &UMessage{}
	packet.Init()
	return packet
}