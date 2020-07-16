package login

import (
	"Zmin/engine/zattr"
	"Zmin/engine/zconf"
	"Zmin/engine/zlog"
	"Zmin/engine/znet"
	"Zmin/engine/zproto"
	"Zmin/engine/zproto/pb"
)

func (service ULoginService) MessageLoop() {
	for {
		select {
		case message := <-service.MessageQueue:
			proxy := message.Proxy
			messageType := message.MessageType
			packet := message.Packet
			switch messageType {
			case zconf.MT_FROM_CENTER, zconf.MT_BROADCAST:
				zproto.PbMessageHandle(proxy, packet)
			default:

			}
			//case <-service.ticker:
			//	post.Tick()
			//	service.sendEntitySyncInfosToGames()
			//	break
			//	default:
		}
	}
}
func (service *ULoginService) InitDownHandles() {
	service.UService.InitDownHandles()
	reqHandleMaps[TCmd(pb.CommandList_MT_ADD_ENGINE_COMPONENT)] = service.AddEngineComponent
	reqHandleMaps[TCmd(pb.CommandList_MT_ADD_ENGINE_COMPONENT_ACK)] = service.AddEngineComponentAck
}
func (service *ULoginService) ConnectToCenter() {
	centerProxy = service.MakeCenterProxy()
	if centerProxy == nil {
		service.Close()
	}
	request := znet.NewRequest(TCmd(pb.CommandList_MT_ADD_ENGINE_COMPONENT), zconf.MT_TO_CENTER)
	message := &pb.ADD_ENGINE_COMPONENT{
		Type:       pb.COMPONENT_TYPE_LOGIN,
		ListenAddr: service.Config.ListenAddr,
	}
	zproto.SendPbMessage(centerProxy, message, request)
	request.Release()
}
func (service *ULoginService) AddEngineComponentAck(proxy *UClientProxy, request *URequest) {
	message, ok := request.ProtoMessage.(*pb.ADD_ENGINE_COMPONENT_ACK)
	if !ok {
		zlog.Error("AddEngineComponent recv error request : ", proxy, request)
		return
	}
	proxy.SetProperty(zattr.Uint32ComponentId, uint32(message.ComponentId))
	proxy.SetProperty(zattr.Uint32ComponentType, uint32(pb.COMPONENT_TYPE_CENTER))
	for _, login := range message.ComponentList {
		gameProxyMaps[login.ComponentId] = service.MakeClientProxy(string(login.ListenAddr))
	}
}
func (service *ULoginService) AddEngineComponent(proxy *UClientProxy, request *URequest) {
	message, ok := request.ProtoMessage.(*pb.ADD_ENGINE_COMPONENT)
	if !ok {
		zlog.Error("AddEngineComponent recv error request : ", proxy, request)
		return
	}
	gameProxyMaps[message.ComponentId] = service.MakeClientProxy(string(message.ListenAddr))
}