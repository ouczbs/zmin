package center

import (
	"github.com/ouczbs/zmin/component/base"
	"github.com/ouczbs/zmin/engine/data/zconf"
	"github.com/ouczbs/zmin/engine/net/znet"
)

type UCenterService struct {
	*UService
}

func NewCenterService() *UCenterService {
	return &UCenterService{
		UService: base.NewService(reqHandleMaps),
	}
}
func (service *UCenterService) Run() {
	service.InitConfig()
	service.initService()
	go service.MessageLoop()
	service.MakeOwnerProxy(zconf.COMPONENT_TYPE_VERSION)
	znet.ServeTCPForever(service.Config.ListenAddr, service)
}
func (service *UCenterService) initService() {
	service.InitDownHandles()
}
