package version

import (
	"github.com/ouczbs/zmin/engine/core/zlog"
	"github.com/ouczbs/zmin/engine/data/zconf"
	"github.com/ouczbs/zmin/engine/net/zmessage"
	"github.com/ouczbs/zmin/engine/sync/zattr"
	"github.com/ouczbs/zmin/engine/sync/zpb"
	"github.com/ouczbs/zmin/engine/sync/zproto"
	"math/rand"
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
			packet := message.Packet
			zproto.PbMessageHandle(proxy, packet, message.Cmd)
			//case <-service.ticker:
			//	post.Tick()
			//	service.sendEntitySyncInfosToGames()
			//	break
			//	default:
		}
	}
}
func (service *UVersionService) AddEngineComponent(proxy *UClientProxy, request *URequest) {
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
	case zconf.COMPONENT_TYPE_LOGIN:
		loginList = append(loginList, proxy)
		service.UtilAddEngineComponentAck(proxy, nil)
	case zconf.COMPONENT_TYPE_CENTER:
		centerMaps[id] = proxy
		service.UtilAddEngineComponentAck(proxy, nil)
	case zconf.COMPONENT_TYPE_CLIENT:
		service.ToAddClientAck(proxy)
	}
	zlog.Debug("AddEngineComponent:", id, message.Type)
}
func (service *UVersionService) InitDownHandles() {
	service.UService.InitDownHandles()
	reqHandleMaps[zconf.CMD_ADD_ENGINE_COMPONENT] = service.AddEngineComponent
}
func (service *UVersionService) ToAddClientAck(proxy *UClientProxy) {
	login := loginList[rand.Intn(len(loginList))]
	if login == nil {
		return
	}
	message := &zpb.ADD_CLIENT_ACK{}
	message.ComponentId = clientID.Inc()
	message.LoginAddr = login.GetProperty(zattr.StringListenAddr).(string)
	request := zmessage.NewRequest(zconf.CMD_ADD_CLIENT_ACK, zconf.MT_TO_CLIENT, message)
	zproto.SendPbMessage(proxy, request)
	request.Release()
}
