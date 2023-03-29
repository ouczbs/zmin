package gate

import (
	"github.com/ouczbs/zmin/engine/core/zlog"
	"github.com/ouczbs/zmin/engine/data/zconf"
	"github.com/ouczbs/zmin/engine/net/zmessage"
	"github.com/ouczbs/zmin/engine/sync/zattr"
	"github.com/ouczbs/zmin/engine/sync/zpb"
	"github.com/ouczbs/zmin/engine/sync/zproto"
)

func (service *UGateService) MessageLoop() {
	for {
		select {
		case message := <-service.MessageQueue:
			proxy, ok := message.Proxy.(*UClientProxy)
			if !ok {
				zlog.Error("MessageLoop Recv Unknown Proxy", proxy)
				break
			}
			messageType := message.MessageType
			packet := message.Packet
			switch messageType {
			case zconf.MT_TO_SERVER, zconf.MT_BROADCAST, zconf.MT_TO_ALL:
				zproto.PbMessageHandle(proxy, packet, message.Cmd)
			case zconf.MT_TO_CLIENT:
				service.ForwardToClient(packet)
			default:
				if messageType > zconf.MT_TO_GAME_START && messageType < zconf.MT_TO_GAME_END {
					service.ForwardToGame(proxy, messageType, packet)
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
func (service *UGateService) InitDownHandles() {
	service.UService.InitDownHandles()
	reqHandleMaps[TCmd(zpb.CommandList_MT_SYNC_PROXY_PROPERTY)] = service.SyncProxyProperty
	reqHandleMaps[TCmd(zpb.CommandList_MT_ADD_ENGINE_COMPONENT)] = service.AddEngineComponent
	reqHandleMaps[TCmd(zpb.CommandList_MT_ADD_ENGINE_COMPONENT_ACK)] = service.AddEngineComponentAck
}
func (service *UGateService) ConnectToCenter() {
	centerProxy = service.MakeCenterProxy()
	if centerProxy == nil {
		service.Close()
	}
	message := &zpb.ADD_ENGINE_COMPONENT{
		ComponentId: service.Config.ComponentId,
		Type:        zpb.COMPONENT_TYPE_GATE,
		ListenAddr:  service.Config.ListenAddr,
	}
	request := zmessage.NewRequest(TCmd(zpb.CommandList_MT_ADD_ENGINE_COMPONENT), zconf.MT_TO_SERVER, message)
	zproto.SendPbMessage(centerProxy, request)
	request.Release()
}
func (service *UGateService) AddEngineComponentAck(proxy *UClientProxy, request *URequest) {
	message, ok := request.ProtoMessage.(*zpb.ADD_ENGINE_COMPONENT_ACK)
	if !ok {
		zlog.Error("AddEngineComponent recv error request : ", proxy, request)
		return
	}
	proxy.SetProperty(zattr.Int32ComponentId, int32(message.ComponentId))
	proxy.SetProperty(zattr.Int32ComponentType, int32(zpb.COMPONENT_TYPE_CENTER))
	for _, game := range message.ComponentList {
		service.MakeClientProxy(string(game.ListenAddr), zpb.COMPONENT_TYPE_GAME)
		zlog.Debug("AddEngineComponentAck:login listen addr: ", game.ListenAddr)
	}
}
func (service *UGateService) AddEngineComponent(proxy *UClientProxy, request *URequest) {
	message, ok := request.ProtoMessage.(*zpb.ADD_ENGINE_COMPONENT)
	if !ok {
		zlog.Error("AddEngineComponent recv error request : ", proxy, request)
		return
	}
	service.MakeClientProxy(string(message.ListenAddr), zpb.COMPONENT_TYPE_GAME)
	zlog.Debug("AddEngineComponent ", message.Type, message.ListenAddr)
}
func (service *UGateService) ForwardToGame(proxy *UClientProxy, messageType TMessageType, packet *UPacket) {
	id, ok := proxy.GetProperty(zattr.Int32ComponentId).(int32)
	if !ok {
		return
	}
	packet.AppendComponentId(id)
	game := gameMessageMaps[messageType]
	if game != nil {
		game.ForwardPacket(packet)
	}
}
func (service *UGateService) ForwardToClient(packet *UPacket) {
	id := packet.SubtractComponentId()
	proxy := clientProxyMaps[id]
	if proxy != nil {
		proxy.ForwardPacket(packet)
	}
}
func (service *UGateService) SyncProxyProperty(proxy *UClientProxy, request *URequest) {
	service.UService.SyncProxyProperty(proxy, request)
	componentType, ok := proxy.GetProperty(zattr.Int32ComponentType).(int32)
	if ok && componentType == int32(zpb.COMPONENT_TYPE_GAME) {
		if messageType, ok := proxy.GetProperty(zattr.Int32MessageType).(int32); ok {
			gameMessageMaps[TMessageType(messageType)] = proxy
		}
	}
}
