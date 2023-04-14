package base

import (
	"github.com/ouczbs/zmin/engine/core/zlog"
	"github.com/ouczbs/zmin/engine/data/zconf"
	"github.com/ouczbs/zmin/engine/net/zmessage"
	"github.com/ouczbs/zmin/engine/net/znet"
	"github.com/ouczbs/zmin/engine/sync/zattr"
	"github.com/ouczbs/zmin/engine/sync/zpb"
	"github.com/ouczbs/zmin/engine/sync/zproto"
	"net"
)

func (service *UService) MakeClientProxy(addr string, componentType TComponentType) *znet.UClientProxy {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		zlog.Infof(" MakeClientProxy error , addr %s , err %s", addr, err)
		return nil
	}
	proxy := znet.NewClientProxy(service, conn)
	proxy.SetProperty(zattr.Int32ComponentType, int32(componentType))
	go proxy.Serve()
	return proxy
}
func (service *UService) MakeOwnerProxy(ownerType TComponentType) *znet.UClientProxy {
	ownerProxy := service.MakeClientProxy(service.Config.OwnerAddr, ownerType)
	if ownerProxy == nil {
		service.Close()
	}
	ownerProxy.SetProperty(zattr.BoolIsOwnerProxy, true)
	message := &zpb.ADD_ENGINE_COMPONENT{
		ComponentId: service.Config.ComponentId,
		Type:        service.Config.ComponentType,
		ListenAddr:  service.Config.ListenAddr,
	}
	request := zmessage.NewRequest(zconf.CMD_ADD_ENGINE_COMPONENT, zconf.MT_TO_SERVER, message)
	zproto.SendPbMessage(ownerProxy, request)
	request.Release()
	return ownerProxy
}
func (service *UService) SyncProxyProperty(proxy *UClientProxy, request *URequest) {
	message, ok := request.ProtoMessage.(*zpb.SYNC_PROXY_PROPERTY)
	if !ok {
		zlog.Error("AddEngineComponent recv error request : ", proxy, request)
		return
	}
	for _, property := range message.PropertyMapList {
		proxy.SetProperty(property.Key, zattr.ReadProperty(property.Value))
		zlog.Debug("SyncProxyProperty:attr k:", property.Key)
	}
	request.Release()
}
func (service *UService) AddEngineComponentAck(proxy *UClientProxy, request *URequest) {
	message, ok := request.ProtoMessage.(*zpb.ADD_ENGINE_COMPONENT_ACK)
	if !ok {
		zlog.Error("AddEngineComponent recv error request : ", proxy, request)
		return
	}
	proxy.SetProperty(zattr.Int32ComponentId, message.ComponentId)
	proxy.SetProperty(zattr.Int32ComponentType, message.ComponentId)
}
func (service *UService) UtilAddEngineComponentAck(proxy *UClientProxy, componentMaps TProxyMap) {
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
	request := zmessage.NewRequest(zconf.CMD_ADD_ENGINE_COMPONENT_ACK, zconf.MT_TO_SERVER, message)
	zproto.SendPbMessage(proxy, request)
	request.Release()
}
func (service *UService) UtilBroadcastAddEngineComponent(message *zpb.ADD_ENGINE_COMPONENT, componentMaps TProxyMap) {
	request := zmessage.NewRequest(zconf.CMD_ADD_ENGINE_COMPONENT, zconf.MT_TO_SERVER, message)
	packet := zproto.MakePbMessagePacket(request)
	for _, comp := range componentMaps {
		comp.SendPacket(packet)
	}
	packet.Release()
}
func (service *UService) AddEngineComponent(proxy *UClientProxy, request *URequest) {
	message, ok := request.ProtoMessage.(*zpb.ADD_ENGINE_COMPONENT)
	if !ok {
		zlog.Error("AddEngineComponent recv error request : ", proxy, request)
		return
	}
	service.MakeClientProxy(message.ListenAddr, message.Type)
	zlog.Debug("AddEngineComponent dispatcher ", message.Type, message.ListenAddr)
}

func (service *UService) InitDownHandles() {
	service.ReqHandleMaps[zconf.CMD_SYNC_PROXY_PROPERTY] = service.SyncProxyProperty
	service.ReqHandleMaps[zconf.CMD_ADD_ENGINE_COMPONENT] = service.AddEngineComponent
}
