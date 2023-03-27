package center

import (
	"github.com/ouczbs/zmin/engine/zattr"
	"github.com/ouczbs/zmin/engine/zconf"
	"github.com/ouczbs/zmin/engine/zlog"
	"github.com/ouczbs/zmin/engine/znet"
	"github.com/ouczbs/zmin/engine/zproto"
	"github.com/ouczbs/zmin/engine/zproto/zpb"
)

func (service *UCenterService) MessageLoop() {
	for {
		select {
		case message := <-service.MessageQueue:
			proxy := message.Proxy
			messageType := message.MessageType
			packet := message.Packet
			switch messageType {
			case zconf.MT_TO_SERVER, zconf.MT_TO_ALL:
				zproto.PbMessageHandle(proxy, packet, message.Cmd)
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
func (service *UCenterService) UtilBroadcastAddEngineComponent(message *zpb.ADD_ENGINE_COMPONENT, componentMaps TProxyMap) {
	request := znet.NewRequest(TCmd(zpb.CommandList_MT_ADD_ENGINE_COMPONENT), zconf.MT_BROADCAST, message)
	packet := zproto.MakePbMessagePacket(request)
	for _, comp := range componentMaps {
		comp.SendPacket(packet)
	}
	packet.Release()
}

func (service *UCenterService) UtilAddEngineComponentAck(proxy *UClientProxy, componentMaps TProxyMap) {
	message := &zpb.ADD_ENGINE_COMPONENT_ACK{}
	message.ComponentId = service.Config.ComponentId
	if componentMaps != nil {
		for _, comp := range componentMaps {
			addr := comp.GetProperty(zattr.StringListenAddr).(string)
			id := comp.GetProperty(zattr.Int32ComponentId).(int32)
			component := &zpb.ADD_ENGINE_COMPONENT{ListenAddr: addr, ComponentId: id}
			message.ComponentList = append(message.ComponentList, component)
		}
	}
	request := znet.NewRequest(TCmd(zpb.CommandList_MT_ADD_ENGINE_COMPONENT_ACK), zconf.MT_TO_SERVER, message)
	zproto.ResponseMessage(proxy, request)
	request.Release()
}
func (service *UCenterService) AddEngineComponent(proxy *UClientProxy, request *URequest) {
	message, ok := request.ProtoMessage.(*zpb.ADD_ENGINE_COMPONENT)
	if !ok {
		zlog.Error("AddEngineComponent recv error request : ", proxy, request)
		return
	}
	sequence := message.ComponentId
	proxy.SetProperty(zattr.Int32ComponentId, int32(sequence))
	proxy.SetProperty(zattr.Int32ComponentType, int32(message.Type))
	proxy.SetProperty(zattr.StringListenAddr, string(message.ListenAddr))
	switch message.Type {
	case zpb.COMPONENT_TYPE_GAME:
		gameProxyMaps[sequence] = proxy
		service.UtilAddEngineComponentAck(proxy, dispatcherProxyMaps)
		service.UtilBroadcastAddEngineComponent(message, loginProxyMaps)
		service.UtilBroadcastAddEngineComponent(message, gateProxyMaps)
	case zpb.COMPONENT_TYPE_LOGIN:
		loginProxyMaps[sequence] = proxy
		service.UtilAddEngineComponentAck(proxy, gameProxyMaps)
		service.UtilBroadcastAddEngineComponent(message, gateProxyMaps)
	case zpb.COMPONENT_TYPE_DISPATCHER:
		dispatcherProxyMaps[sequence] = proxy
		service.UtilAddEngineComponentAck(proxy, nil)
		service.UtilBroadcastAddEngineComponent(message, gameProxyMaps)
	case zpb.COMPONENT_TYPE_GATE:
		gateProxyMaps[sequence] = proxy
		service.UtilAddEngineComponentAck(proxy, gameProxyMaps)
	}
	request.Release()
	zlog.Debug("AddEngineComponent:", sequence, message.Type)
}
func (service *UCenterService) InitDownHandles() {
	service.UService.InitDownHandles()
	reqHandleMaps[TCmd(zpb.CommandList_MT_ADD_ENGINE_COMPONENT)] = service.AddEngineComponent
}
