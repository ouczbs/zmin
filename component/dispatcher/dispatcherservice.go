package dispatcher

import (
	"github.com/ouczbs/zmin/component/base"
	"github.com/ouczbs/zmin/engine/data/zconf"
	"github.com/ouczbs/zmin/engine/net/znet"
)

type UDispatcherService struct {
	*UService
}

func NewDispatcherService() *UDispatcherService {
	return &UDispatcherService{
		UService: base.NewService(reqHandleMaps),
	}
}
func (service *UDispatcherService) Run() {
	service.InitConfig()
	service.initService()
	go service.MessageLoop()
	service.MakeOwnerProxy(zconf.COMPONENT_TYPE_CENTER)
	znet.ServeTCPForever(service.Config.ListenAddr, service)
}
func (service *UDispatcherService) initService() {
	service.InitDownHandles()
}
