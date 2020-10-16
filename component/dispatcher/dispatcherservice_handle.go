package dispatcher

import (
	"github.com/ouczbs/Zmin/engine/zattr"
	"github.com/ouczbs/Zmin/engine/zconf"
	"github.com/ouczbs/Zmin/engine/zlog"
	"github.com/ouczbs/Zmin/engine/znet"
	"github.com/ouczbs/Zmin/engine/zproto"
	"github.com/ouczbs/Zmin/engine/zproto/zpb"
)

func (service *UDispatcherService) MessageLoop() {
	for {
		select {
		case message := <-service.MessageQueue:
			proxy := message.Proxy
			messageType := message.MessageType
			packet := message.Packet
			switch messageType {
			case zconf.MT_TO_SERVER, zconf.MT_BROADCAST:
				zproto.PbMessageHandle(proxy, packet , message.Cmd)
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
	request := znet.NewRequest(TCmd(zpb.CommandList_MT_ADD_ENGINE_COMPONENT), zconf.MT_TO_SERVER , message)
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
