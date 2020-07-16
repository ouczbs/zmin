package login

import (
	"Zmin/component/base"
	"Zmin/engine/zconf"
	"Zmin/engine/zlog"
	"Zmin/engine/znet"
)

type ULoginService struct {
	*UService
	Config              * zconf.ULoginConfig
}

func NewLoginService() *ULoginService {
	return &ULoginService{
		UService:base.NewService(reqHandleMaps),
		Config: zconf.GetLoginConfig(),
	}
}
func (service * ULoginService)Run() {
	service.initService()
	go service.MessageLoop()
	service.ConnectToCenter()
	znet.ServeTCPForever(service.Config.ListenAddr, service)
}
func (service *ULoginService) initService(){
	zlog.SetOutput([]string{ "stderr", service.Config.LogFile })
	service.InitDownHandles()
}