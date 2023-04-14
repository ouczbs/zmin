package center

import (
	"github.com/ouczbs/zmin/engine/core/zlog"
	"github.com/ouczbs/zmin/engine/data/zconf"
	"github.com/ouczbs/zmin/engine/sync/zattr"
	"github.com/ouczbs/zmin/engine/sync/zpb"
	"github.com/ouczbs/zmin/engine/sync/zproto"
)

func (service *UCenterService) MessageLoop() {
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
			default:
				zlog.Debug("un handle type ", messageType)
			}
			//case <-service.ticker:
			//	post.Tick()
			//	service.sendEntitySyncInfosToGames()
			//	break
			//	default:
		}
	}
}
func (service *UCenterService) AddEngineComponent(proxy *UClientProxy, request *URequest) {
	message, ok := request.ProtoMessage.(*zpb.ADD_ENGINE_COMPONENT)
	if !ok {
		zlog.Error("AddEngineComponent recv error request : ", proxy, request)
		return
	}
	id := message.ComponentId
	proxy.SetProperty(zattr.Int32ComponentId, id)
	proxy.SetProperty(zattr.Int32ComponentType, message.Type)
	proxy.SetProperty(zattr.StringListenAddr, message.ListenAddr)
	switch message.Type {
	case zconf.COMPONENT_TYPE_GAME:
		gameProxyMaps[id] = proxy
		service.UtilAddEngineComponentAck(proxy, nil)
		service.UtilBroadcastAddEngineComponent(message, dispatcherProxyMaps)
	case zconf.COMPONENT_TYPE_DISPATCHER:
		dispatcherProxyMaps[id] = proxy
		service.UtilAddEngineComponentAck(proxy, nil)
		service.UtilBroadcastAddEngineComponent(message, gateProxyMaps)
	case zconf.COMPONENT_TYPE_GATE:
		gateProxyMaps[id] = proxy
		service.UtilAddEngineComponentAck(proxy, dispatcherProxyMaps)
	}
	zlog.Debug("AddEngineComponent:", id, message.Type)
}
func (service *UCenterService) InitDownHandles() {
	service.UService.InitDownHandles()
	reqHandleMaps[zconf.CMD_ADD_ENGINE_COMPONENT] = service.AddEngineComponent
}
