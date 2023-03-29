package dispatcher

import (
	"github.com/ouczbs/zmin/engine/core/zlog"
	"github.com/ouczbs/zmin/engine/data/zconf"
	"github.com/ouczbs/zmin/engine/net/zmessage"
	"github.com/ouczbs/zmin/engine/sync/zattr"
	"github.com/ouczbs/zmin/engine/sync/zpb"
	"github.com/ouczbs/zmin/engine/sync/zproto"
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
			case zconf.MT_TO_SERVER, zconf.MT_BROADCAST:
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
func (service *UDispatcherService) InitDownHandles() {
	service.UService.InitDownHandles()
	reqHandleMaps[TCmd(zpb.CommandList_MT_ADD_ENGINE_COMPONENT_ACK)] = service.AddEngineComponentAck
}
func (service *UDispatcherService) ConnectToCenter() {
	centerProxy = service.MakeCenterProxy()
	if centerProxy == nil {
		service.Close()
	}

	message := &zpb.ADD_ENGINE_COMPONENT{
		ComponentId: service.Config.ComponentId,
		Type:        zpb.COMPONENT_TYPE_DISPATCHER,
		ListenAddr:  service.Config.ListenAddr,
	}
	request := zmessage.NewRequest(TCmd(zpb.CommandList_MT_ADD_ENGINE_COMPONENT), zconf.MT_TO_SERVER, message)
	zproto.SendPbMessage(centerProxy, request)
	request.Release()
}
func (service *UDispatcherService) AddEngineComponentAck(proxy *UClientProxy, request *URequest) {
	message, ok := request.ProtoMessage.(*zpb.ADD_ENGINE_COMPONENT_ACK)
	if !ok {
		zlog.Error("AddEngineComponent recv error request : ", proxy, request)
		return
	}
	proxy.SetProperty(zattr.Int32ComponentId, int32(message.ComponentId))
	proxy.SetProperty(zattr.Int32ComponentType, int32(zpb.COMPONENT_TYPE_CENTER))
}
