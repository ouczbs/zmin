package version

import (
	"github.com/ouczbs/zmin/component/base"
	"github.com/ouczbs/zmin/engine/net/znet"
)

type UVersionService struct {
	*UService
}

func NewVersionService() *UVersionService {
	return &UVersionService{
		UService: base.NewService(reqHandleMaps),
	}
}
func (service *UVersionService) Run() {
	service.InitConfig()
	service.initService()
	go service.MessageLoop()
	znet.ServeTCPForever(service.Config.ListenAddr, service)
}
func (service *UVersionService) initService() {
	service.InitDownHandles()
}
