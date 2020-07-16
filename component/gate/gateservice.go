package gate

import (
	"Zmin/component/base"
	"Zmin/engine/zconf"
	"Zmin/engine/zlog"
	"Zmin/engine/znet"
)

type UGateService struct {
	*UService
	Config              * zconf.UGateConfig
}

func NewGateService() *UGateService {
	return &UGateService{
		UService:base.NewService(reqHandleMaps),
		Config: zconf.GetGateConfig(),
	}
}
func (service * UGateService)Run() {
	service.initService()
	go service.MessageLoop()
	service.ConnectToCenter()
	znet.ServeTCPForever(service.Config.ListenAddr, service)
}
func (service *UGateService) initService(){
	zlog.SetOutput([]string{ "stderr", service.Config.LogFile })
	service.InitDownHandles()
}