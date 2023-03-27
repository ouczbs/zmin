package zclass

import (
	"github.com/ouczbs/zmin/engine/zconf"
	"sync"
)

type (
	TSize = zconf.TSize
)
type IPoolObject = interface {}

type UStackPool struct {
	Size TSize
	IsStatic bool

	head TSize
	tail TSize
	mutex sync.Mutex

	Pool []IPoolObject
}
func (stack * UStackPool) Init(){
	stack.IsStatic = false
	stack.Pool = make([]IPoolObject , stack.Size)
}
func (stack * UStackPool) Len()TSize{
	return (stack.head - stack.tail + stack.Size) % stack.Size
}
func (stack * UStackPool) MakeAdd(size TSize){
	stack.head = stack.Size
	stack.tail = 0
	stack.Size += size
	pool := make([]IPoolObject , size)
	stack.Pool = append(stack.Pool , pool...)
}
func (stack * UStackPool) SyncPush(object IPoolObject){
	head := stack.head
	stack.Pool[head] = object
	stack.head = ( head + 1 ) % stack.Size
	if stack.head == stack.tail { //栈满
		if stack.IsStatic {
			return
		}
		stack.MakeAdd(stack.Size)
	}
}
func (stack * UStackPool) Push(object IPoolObject){
	stack.mutex.Lock()
	stack.SyncPush(object)
	stack.mutex.Unlock()
}
func (stack * UStackPool) SyncPop()IPoolObject{
	tail := stack.tail
	if tail == stack.head { //栈空
		return nil
	}
	stack.tail = (tail + stack.Size - 1) % stack.Size
	return stack.Pool[tail]
}
func (stack * UStackPool) Pop() IPoolObject{
	stack.mutex.Lock()
	object := stack.SyncPop()
	stack.mutex.Unlock()
	return object
}
