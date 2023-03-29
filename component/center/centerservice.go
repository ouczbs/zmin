package center

import (
	"github.com/ouczbs/zmin/component/base"
	"github.com/ouczbs/zmin/engine/core/zlog"
	"github.com/ouczbs/zmin/engine/net/znet"
	"github.com/ouczbs/zmin/engine/sync/zattr"
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
	service.UService.Run()
	service.initService()
	go service.MessageLoop()
	znet.ServeTCPForever(service.Config.ListenAddr, service)
}
func (service *UCenterService) initService() {
	logFile, ok := service.GetProperty(zattr.StringLogFile).(string)
	if !ok {
		logFile = base.CenterConfig.LogFile
	}
	zlog.SetOutput([]string{"stderr", logFile})
	service.InitDownHandles()
}
