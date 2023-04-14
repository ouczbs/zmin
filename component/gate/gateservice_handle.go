package gate

import (
	"github.com/ouczbs/zmin/engine/core/zlog"
	"github.com/ouczbs/zmin/engine/data/zconf"
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
			case zconf.MT_TO_SERVER:
				zproto.PbMessageHandle(proxy, packet, message.Cmd)
			case zconf.MT_TO_CLIENT:
				service.ForwardToClient(packet)
			default:
				service.ForwardToDispatcher(proxy, packet)
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
	reqHandleMaps[zconf.CMD_SYNC_PROXY_PROPERTY] = service.SyncProxyProperty
	reqHandleMaps[zconf.CMD_ADD_ENGINE_COMPONENT_ACK] = service.AddEngineComponentAck
}
func (service *UGateService) AddEngineComponentAck(proxy *UClientProxy, request *URequest) {
	message, ok := request.ProtoMessage.(*zpb.ADD_ENGINE_COMPONENT_ACK)
	if !ok {
		zlog.Error("AddEngineComponent recv error request : ", proxy, request)
		return
	}
	proxy.SetProperty(zattr.Int32ComponentId, message.ComponentId)
	proxy.SetProperty(zattr.Int32ComponentType, ownerType)
	for _, dispatcher := range message.ComponentList {
		service.MakeClientProxy(dispatcher.ListenAddr, zconf.COMPONENT_TYPE_DISPATCHER)
		zlog.Debug("AddEngineComponentAck:login dispatcher listen addr: ", dispatcher.ListenAddr)
	}
}
func (service *UGateService) ForwardToDispatcher(proxy *UClientProxy, packet *UPacket) {
	id, ok := proxy.GetProperty(zattr.Int32ComponentId).(int32)
	if !ok {
		return
	}
	packet.AppendComponentId(id)
	dispatcher := dispatcherList[id%dispatcherSize]
	if dispatcher != nil {
		dispatcher.ForwardPacket(packet)
	}
}
func (service *UGateService) ForwardToClient(packet *UPacket) {
	id := packet.SubtractComponentId()
	proxy := clientMaps[id]
	if proxy != nil {
		proxy.ForwardPacket(packet)
	}
}
func (service *UGateService) SyncProxyProperty(proxy *UClientProxy, request *URequest) {
	service.UService.SyncProxyProperty(proxy, request)
	componentType, ok := proxy.GetProperty(zattr.Int32ComponentType).(int32)
	if !ok {
		return
	}
	if componentType == int32(zconf.COMPONENT_TYPE_DISPATCHER) {
		if componentId, ok := proxy.GetProperty(zattr.Int32ComponentId).(int32); ok {
			is := dispatcherMaps[componentId]
			if !is {
				dispatcherList = append(dispatcherList, proxy)
				dispatcherMaps[componentId] = true
			}
		}
	}
	if componentType == int32(zconf.COMPONENT_TYPE_CLIENT) {
		if componentId, ok := proxy.GetProperty(zattr.Int32ComponentId).(int32); ok {
			clientMaps[componentId] = proxy
		}
	}
}
