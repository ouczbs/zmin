package zproto

import (
	"github.com/ouczbs/zmin/engine/data/zconf"
	"github.com/ouczbs/zmin/engine/net/zmessage"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
)

var (
	pbMessageTypes  = make(map[TCmd]IReflectMessageType)
	CommandListName = zconf.CommandList_name
)

func NewRequest(wrap *UWrapMessage, pb IReflectMessage) *URequest {
	request := zmessage.NewRequest(_CMD_INVALID, _MT_INVALID, pb)
	request.Request = wrap.Request
	return request
}
func NewPbMessage(cmd TCmd) (IReflectMessage, error) {
	messageType := pbMessageTypes[cmd]
	if messageType != nil {
		return messageType.New().Interface(), nil
	}
	pbName := CommandListName[cmd]
	messageType, err := protoregistry.GlobalTypes.FindMessageByName(protoreflect.FullName(pbName))
	if err != nil {
		return nil, err
	}
	pbMessageTypes[cmd] = messageType
	return messageType.New().Interface(), err
}
func PbMessageHandle(proxy *UClientProxy, packet *UPacket, cmd TCmd) {
	wrapBytes := packet.MessagePayload()
	wrapMessage := &UWrapMessage{}
	Unmarshal(wrapBytes, wrapMessage)
	handle, globalHandle := proxy.GetRequestHandles(wrapMessage.Response, cmd)
	if cmd == 0 || (handle == nil && globalHandle == nil) {
		return
	}
	pbMessage, _ := NewPbMessage(cmd)
	Unmarshal(wrapMessage.Content, pbMessage)
	request := NewRequest(wrapMessage, pbMessage)
	proxy.Then(handle, request).Then(globalHandle, request)
	request.Release()
}
