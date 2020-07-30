package znet

import "github.com/ouczbs/Zmin/engine/zconf"

var (
	requestPool = NewRequestPool(zconf.CPoolRequestSize)
)

type URequest struct {
	ProtoMessage IReflectMessage
	Code         TCode
	Next		 bool
	// Request
	MessageType  TMessageType
	Cmd 		 TCmd
	Request      TCallId

	isReleased bool
}

func NewRequest(cmd TCmd, messageType TMessageType)*URequest{
	request := requestPool.Pop()
	request.isReleased = false
	request.Next = true
	request.Cmd = cmd
	request.MessageType = messageType
	return request
}
func (request * URequest) Init() {

}
func (request * URequest) Release() {
	if request.isReleased {
		return
	}
	request.isReleased = true
	requestPool.Push(request)
}
