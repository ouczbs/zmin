package version

import (
	"github.com/ouczbs/zmin/engine/core/zlog"
	"github.com/ouczbs/zmin/engine/data/zconf"
	"github.com/ouczbs/zmin/engine/net/zmessage"
	"github.com/ouczbs/zmin/engine/sync/zattr"
	"github.com/ouczbs/zmin/engine/sync/zpb"
	"github.com/ouczbs/zmin/engine/sync/zproto"
)

func (service *UVersionService) MessageLoop() {
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

func (service *UVersionService) UtilAddEngineComponentAck(proxy *UClientProxy, componentMaps TProxyMap) {
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
	request := zmessage.NewRequest(zconf.MT_ADD_ENGINE_COMPONENT_ACK, zconf.MT_TO_SERVER, message)
	zproto.ResponseMessage(proxy, request)
	request.Release()
}
func (service *UVersionService) AddEngineComponent(proxy *UClientProxy, request *URequest) {
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
	case zconf.COMPONENT_TYPE_LOGIN:
		loginProxyMaps[sequence] = proxy
	case zconf.COMPONENT_TYPE_CENTER:
		centerProxyMaps[sequence] = proxy
		service.UtilAddEngineComponentAck(proxy, nil)
	}
	request.Release()
	zlog.Debug("AddEngineComponent:", sequence, message.Type)
}
func (service *UVersionService) InitDownHandles() {
	service.UService.InitDownHandles()
	reqHandleMaps[zconf.MT_ADD_ENGINE_COMPONENT] = service.AddEngineComponent
}
