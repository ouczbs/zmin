package center

import (
	"Zmin/component/base"
	"Zmin/engine/zconf"
	"Zmin/engine/zlog"
	"Zmin/engine/znet"
)

type UCenterService struct {
	*UService
	Config              * zconf.UCenterConfig
}

func NewCenterService() *UCenterService {
	return &UCenterService{
		UService:base.NewService(reqHandleMaps),
		Config: zconf.GetCenterConfig(),
	}
}
func (service * UCenterService)Run() {
	service.initService()
	go service.MessageLoop()
	znet.ServeTCPForever(service.Config.ListenAddr, service)
}
func (service *UCenterService) initService(){
	zlog.SetOutput([]string{ "stderr", service.Config.LogFile })
	service.initDownHandles()
}