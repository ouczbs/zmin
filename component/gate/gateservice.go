package gate

import (
	"github.com/ouczbs/zmin/component/base"
	"github.com/ouczbs/zmin/engine/data/zconf"
	"github.com/ouczbs/zmin/engine/net/znet"
)

type UGateService struct {
	*UService
}

func NewGateService() *UGateService {
	return &UGateService{
		UService: base.NewService(reqHandleMaps),
	}
}
func (service *UGateService) Run() {
	service.InitConfig()
	service.initService()
	go service.MessageLoop()
	service.MakeOwnerProxy(zconf.COMPONENT_TYPE_CENTER)
	znet.ServeTCPForever(service.Config.ListenAddr, service)
}
func (service *UGateService) initService() {
	service.InitDownHandles()
}
