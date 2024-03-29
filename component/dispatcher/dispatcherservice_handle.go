package dispatcher

import (
	"github.com/ouczbs/zmin/engine/core/zlog"
	"github.com/ouczbs/zmin/engine/data/zcache"
	"github.com/ouczbs/zmin/engine/data/zconf"
	"github.com/ouczbs/zmin/engine/sync/zattr"
	"github.com/ouczbs/zmin/engine/sync/zpb"
	"github.com/ouczbs/zmin/engine/sync/zproto"
	"go.mongodb.org/mongo-driver/bson"
)

func (service *UDispatcherService) MessageLoop() {
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
			case zconf.MT_TO_SERVER:
				zproto.PbMessageHandle(proxy, packet, message.Cmd)
			case zconf.MT_TO_CLIENT:
				service.ForwardToClient(packet)
			case zconf.MT_TO_SCENE:
				service.ForwardToScene(proxy, packet)
			default:
				service.ForwardToGame(messageType, packet)
			}
			//case <-service.ticker:
			//	post.Tick()
			//	service.sendEntitySyncInfosToGames()
			//	break
			//	default:
		}
	}
}
func (service *UDispatcherService) InitDownHandles() {
	service.UService.InitDownHandles()
	reqHandleMaps[zconf.CMD_ADD_ENGINE_COMPONENT_ACK] = service.AddEngineComponentAck
}
func (service *UDispatcherService) AddEngineComponentAck(proxy *UClientProxy, request *URequest) {
	message, ok := request.ProtoMessage.(*zpb.ADD_ENGINE_COMPONENT_ACK)
	if !ok {
		zlog.Error("AddEngineComponent recv error request : ", proxy, request)
		return
	}
	proxy.SetProperty(zattr.Int32ComponentId, message.ComponentId)
	proxy.SetProperty(zattr.Int32ComponentType, int32(zconf.COMPONENT_TYPE_CENTER))
}
func (service *UDispatcherService) ForwardToScene(proxy *UClientProxy, packet *UPacket) {
	id := packet.SubtractComponentId()
	role := roleMaps[id]
	if role == nil {
		if err := zcache.GetMongoClient().FindOne(role, bson.M{"id": id}); err != nil {
			return
		}
		proxyId, ok := proxy.GetProperty(zattr.Int32ComponentId).(int32)
		if !ok {
			return
		}
		role.GateID = proxyId
	}
	game := gameSceneMaps[role.SceneID]
	if game != nil {
		game.ForwardPacket(packet)
	}
}
func (service *UDispatcherService) ForwardToGame(messageType TMessageType, packet *UPacket) {
	game := gameTypeMaps[messageType]
	if game != nil {
		game.ForwardPacket(packet)
	}
}
func (service *UDispatcherService) ForwardToClient(packet *UPacket) {
	id := packet.SubtractComponentId()
	role := roleMaps[id]
	if role == nil {
		return
	}
	proxy := gateMaps[role.GateID]
	if proxy != nil {
		proxy.ForwardPacket(packet)
	}
}
