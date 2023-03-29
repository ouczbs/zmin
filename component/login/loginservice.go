package login

import (
	"github.com/ouczbs/zmin/component/base"
	"github.com/ouczbs/zmin/engine/core/zlog"
	"github.com/ouczbs/zmin/engine/core/zutil"
	"github.com/ouczbs/zmin/engine/net/znet"
	"github.com/ouczbs/zmin/engine/sync/zattr"
	"net"
)

type ULoginService struct {
	*UService
}

func NewLoginService() *ULoginService {
	return &ULoginService{
		UService: base.NewService(reqHandleMaps),
	}
}
func (service *ULoginService) Run() {
	service.UService.Run()
	service.initService()
	go service.MessageLoop()
	service.ConnectToCenter()
	znet.ServeTCPForever(service.Config.ListenAddr, service)
}
func (service *ULoginService) initService() {
	logFile, ok := service.GetProperty(zattr.StringLogFile).(string)
	if !ok {
		logFile = base.LoginConfig.LogFile
	}
	zlog.SetOutput([]string{"stderr", logFile})
	service.InitDownHandles()
}
func (service *ULoginService) NewTcpConnection(conn net.Conn) {
	proxy := znet.NewClientProxy(service, conn)
	id := zutil.IncSequence()
	clientProxyMaps[id] = proxy
	proxy.SetProperty(zattr.Int32ComponentId, int32(id))
	proxy.Serve()
}
