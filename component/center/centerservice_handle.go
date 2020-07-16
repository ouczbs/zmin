package center

import (
	"Zmin/engine/zconf"
	"Zmin/engine/zlog"
	"Zmin/engine/zattr"
	"Zmin/engine/znet"
	"Zmin/engine/zproto"
	"Zmin/engine/zproto/pb"
	"Zmin/engine/zutil"
)

func (service UCenterService) MessageLoop() {
	for {
		select {
		case message := <-service.MessageQueue:
			proxy := message.Proxy
			messageType := message.MessageType
			packet := message.Packet
			switch messageType {
			case zconf.MT_TO_CENTER:
				zproto.PbMessageHandle(proxy , packet)
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
func (service *UCenterService)UtilBroadcastAddEngineComponent(message *pb.ADD_ENGINE_COMPONENT, componentMaps TProxyMap){
	request := znet.NewRequest(TCmd(pb.CommandList_MT_ADD_ENGINE_COMPONENT) , zconf.MT_BROADCAST)
	packet := zproto.MakePbMessagePacket(message , request)
	for _,comp := range componentMaps {
		comp.SendPacket(packet)
	}
	packet.Release()
}

func (service *UCenterService) UtilAddEngineComponentAck(proxy *UClientProxy , componentMaps TProxyMap){
	message := &pb.ADD_ENGINE_COMPONENT_ACK{}
	message.ComponentId = proxy.GetProperty(zattr.Uint32ComponentId).(uint32)
	if componentMaps != nil {
		for _,comp := range componentMaps {
			addr := comp.GetProperty(zattr.StringListenAddr).(string)
			id := comp.GetProperty(zattr.Uint32ComponentId).(uint32)
			component := &pb.ADD_ENGINE_COMPONENT{ ListenAddr: addr , ComponentId: id}
			message.ComponentList = append(message.ComponentList, component)
		}
	}
	request := znet.NewRequest(TCmd(pb.CommandList_MT_ADD_ENGINE_COMPONENT_ACK) , zconf.MT_FROM_CENTER)
	zproto.ResponseMessage(proxy , message , request)
	request.Release()
}
func (service *UCenterService) AddEngineComponent(proxy *UClientProxy, request * URequest) {
	message , ok := request.ProtoMessage.(* pb.ADD_ENGINE_COMPONENT)
	if !ok {
		zlog.Error("AddEngineComponent recv error request : " , proxy , request)
		return
	}
	sequence := zutil.IncSequence()
	if 	message.ComponentId !=0 {
		sequence = message.ComponentId
	}
	proxy.SetProperty(zattr.Uint32ComponentId , uint32(sequence))
	proxy.SetProperty(zattr.Uint32ComponentType , uint32(message.Type))
	proxy.SetProperty(zattr.StringListenAddr , string(message.ListenAddr))
	switch message.Type {
	case pb.COMPONENT_TYPE_GAME:
		gameProxyMaps[sequence] = proxy
		service.UtilAddEngineComponentAck(proxy , dispatcherProxyMaps)
		service.UtilBroadcastAddEngineComponent(message , loginProxyMaps)
	case pb.COMPONENT_TYPE_LOGIN:
		loginProxyMaps[sequence] = proxy
		service.UtilAddEngineComponentAck(proxy , gameProxyMaps)
		service.UtilBroadcastAddEngineComponent(message , gateProxyMaps)
	case pb.COMPONENT_TYPE_DISPATCHER:
		dispatcherProxyMaps[sequence] = proxy
		service.UtilAddEngineComponentAck(proxy , nil)
		service.UtilBroadcastAddEngineComponent(message , gameProxyMaps)
	case pb.COMPONENT_TYPE_GATE:
		gateProxyMaps[sequence] = proxy
		service.UtilAddEngineComponentAck(proxy , loginProxyMaps)
	}
	request.Release()
	zlog.Debug("AddEngineComponent:" ,sequence ,message.Type)
}
func (service *UCenterService) initDownHandles(){
	service.UService.InitDownHandles()
	reqHandleMaps[TCmd(pb.CommandList_MT_ADD_ENGINE_COMPONENT)] = service.AddEngineComponent
}
