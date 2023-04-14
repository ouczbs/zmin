package login

import (
	"github.com/ouczbs/zmin/engine/core/zlog"
	"github.com/ouczbs/zmin/engine/data/zcache"
	"github.com/ouczbs/zmin/engine/data/zconf"
	"github.com/ouczbs/zmin/engine/data/zmodel"
	"github.com/ouczbs/zmin/engine/sync/zpb"
	"github.com/ouczbs/zmin/engine/sync/zproto"
	"go.mongodb.org/mongo-driver/bson"
)

func (service *ULoginService) MessageLoop() {
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
			message.Release()
			//case <-service.ticker:
			//	post.Tick()
			//	service.sendEntitySyncInfosToGames()
			//	break
			//	default:
		}
	}
}
func (service *ULoginService) InitDownHandles() {
	service.UService.InitDownHandles()
	reqHandleMaps[zconf.CMD_ADD_ENGINE_COMPONENT_ACK] = service.AddEngineComponentAck
	service.ReqHandleMaps[zconf.CMD_SYNC_PROXY_PROPERTY] = service.LoginAccount
	service.ReqHandleMaps[zconf.CMD_SYNC_PROXY_PROPERTY] = service.RegisterAccount
}
func (service *ULoginService) LoginAccount(proxy *UClientProxy, request *URequest) {
	message, ok := request.ProtoMessage.(*zpb.LoginAccount)
	if !ok {
		zlog.Error("AddEngineComponent recv error request : ", proxy, request)
		return
	}
	var account *zmodel.FAccount
	var query bson.M
	switch message.Type {
	case LoginType_Password:
		query["UserName"] = message.Name
		query["Password"] = message.Password
	default:

	}
	zcache.GetMongoClient().FindOne(account, query)
	if account == nil {
		zproto.ResponseMessage(proxy, request)
	}
}
func (service *ULoginService) RegisterAccount(proxy *UClientProxy, request *URequest) {
	//todo
}
func (service *ULoginService) CreateRole(proxy *UClientProxy, request *URequest) {
	//todo
}
func (service *ULoginService) EnterGame(proxy *UClientProxy, request *URequest) {
	//todo
}
