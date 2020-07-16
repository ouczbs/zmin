package dispatcher

import (
	"Zmin/component/base"
	"Zmin/engine/zconf"
	"Zmin/engine/zlog"
	"Zmin/engine/znet"
)

type UDispatcherService struct {
	*UService
	Config              * zconf.UDispatcherConfig
}
func NewDispatcherService() *UDispatcherService {
	return &UDispatcherService{
		UService:base.NewService(reqHandleMaps),
		Config: zconf.GetDispatcherConfig(),
	}
}
func (service * UDispatcherService)Run() {
	service.initService()
	go service.MessageLoop()
	service.ConnectToCenter()
	znet.ServeTCPForever(service.Config.ListenAddr, service)
}
func (service *UDispatcherService) initService(){
	zlog.SetOutput([]string{ "stderr", service.Config.LogFile })
	service.InitDownHandles()
}