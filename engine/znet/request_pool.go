package znet


type URequestPool struct {
	* UStackPool
}
func NewRequestPool(size TSize) *URequestPool {
	stack := &URequestPool{
		&UStackPool{
			Size: size,
		},
	}
	stack.Init()
	return stack
}
func (stack * URequestPool) Pop() * URequest{
	object := stack.UStackPool.Pop()
	if object == nil {
		return stack.New()
	}
	return object.(*URequest)
}
func (stack * URequestPool) New() * URequest{
	return &URequest{}
}