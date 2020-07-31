package login

import (
	"github.com/ouczbs/Zmin/engine/zattr"
	"github.com/ouczbs/Zmin/engine/zconf"
	"github.com/ouczbs/Zmin/engine/zlog"
	"github.com/ouczbs/Zmin/engine/znet"
	"github.com/ouczbs/Zmin/engine/zproto"
	"github.com/ouczbs/Zmin/engine/zproto/pb"
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
				zproto.PbMessageHandle(proxy, packet,message.Cmd)
			default:
				if messageType > zconf.MT_TO_GAME_START && messageType < zconf.MT_TO_GAME_END {
					id ,ok := proxy.GetProperty(zattr.Int32ComponentId).(int32)
					if ok {
						packet.AppendComponentId(id)
						service.ForwardToGame(messageType , packet)
					}
				}
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
	service.ReqHandleMaps[TCmd(pb.CommandList_MT_SYNC_PROXY_PROPERTY)] = service.SyncProxyProperty
}
func (service *ULoginService) ConnectToCenter() {
	centerProxy = service.MakeCenterProxy()
	if centerProxy == nil {
		service.Close()
	}
	message := &pb.ADD_ENGINE_COMPONENT{
		ComponentId: service.Config.ComponentId,
		Type:       pb.COMPONENT_TYPE_LOGIN,
		ListenAddr: service.Config.ListenAddr,
	}
	request := znet.NewRequest(TCmd(pb.CommandList_MT_ADD_ENGINE_COMPONENT), zconf.MT_TO_CENTER,message)
	zproto.SendPbMessage(centerProxy, request)
	request.Release()
}
func (service *ULoginService) AddEngineComponentAck(proxy *UClientProxy, request *URequest) {
	message, ok := request.ProtoMessage.(*pb.ADD_ENGINE_COMPONENT_ACK)
	if !ok {
		zlog.Error("AddEngineComponent recv error request : ", proxy, request)
		return
	}
	proxy.SetProperty(zattr.Int32ComponentId, int32(message.ComponentId))
	proxy.SetProperty(zattr.Int32ComponentType, int32(pb.COMPONENT_TYPE_CENTER))
	for _, login := range message.ComponentList {
		gameProxyMaps[login.ComponentId] = service.MakeClientProxy(string(login.ListenAddr),pb.COMPONENT_TYPE_GAME)
	}
}
func (service *ULoginService) AddEngineComponent(proxy *UClientProxy, request *URequest) {
	message, ok := request.ProtoMessage.(*pb.ADD_ENGINE_COMPONENT)
	if !ok {
		zlog.Error("AddEngineComponent recv error request : ", proxy, request)
		return
	}
	gameProxyMaps[message.ComponentId] = service.MakeClientProxy(string(message.ListenAddr),pb.COMPONENT_TYPE_GAME)
	zlog.Debug("AddEngineComponent " , message.Type , message.ListenAddr)
}
func (service *ULoginService) ForwardToGame(messageType TMessageType , packet * UPacket) {
	proxy := gameMessageMaps[messageType]
	proxy.ForwardPacket(packet)
}
func (service *ULoginService) SyncProxyProperty(proxy *UClientProxy, request * URequest) {
	service.UService.SyncProxyProperty(proxy ,request)
	componentType , ok := proxy.GetProperty(zattr.Int32ComponentType).(int32)
	if ok && componentType == int32(pb.COMPONENT_TYPE_GAME){
		if messageType , ok := proxy.GetProperty(zattr.Int32MessageType).(int32); ok{
			gameProxyMaps[messageType] = proxy
		}
	}
}