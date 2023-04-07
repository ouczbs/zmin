package base

import (
	"github.com/ouczbs/zmin/engine/core/zclass"
	"github.com/ouczbs/zmin/engine/core/zlog"
	"github.com/ouczbs/zmin/engine/data/zconf"
	"github.com/ouczbs/zmin/engine/net/znet"
	"github.com/ouczbs/zmin/engine/sync/zattr"
	"github.com/ouczbs/zmin/engine/sync/zpb"
	"net"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type UService struct {
	*UProperty
	ReqHandleMaps map[TCmd]FRequestHandle
	MessageQueue  chan *UMessage
	Config        *FServiceConfig
}

func NewService(reqHandleMaps map[TCmd]FRequestHandle) *UService {
	return &UService{
		UProperty:     zclass.NewProperty(),
		ReqHandleMaps: reqHandleMaps,
		MessageQueue:  make(chan *UMessage, zconf.CQueueMessageSize),
		Config:        nil,
	}
}
func (service *UService) ClientDisconnect(proxy *UClientProxy) {
	zlog.Debugf("ClientDisconnect %s", proxy)
	t, ok := proxy.GetProperty(zattr.Int32ComponentType).(int32)
	zlog.Infof(" ClientDisconnect   ", t)
	if ok && t == int32(zconf.COMPONENT_TYPE_CENTER) {
		zlog.Infof(" Center client notify exit process !!!")
		service.Close()
	}
}
func (service *UService) Close() {
	os.Exit(1)
}
func (service *UService) NewTcpConnection(conn net.Conn) {
	proxy := znet.NewClientProxy(service, conn)
	proxy.Serve()
}
func (service *UService) RecvMessage(message *UMessage) {
	service.MessageQueue <- message
}
func (service *UService) GetRequestHandle(cmd TCmd) FRequestHandle {
	return service.ReqHandleMaps[cmd]
}
func (service *UService) SyncProxyProperty(proxy *UClientProxy, request *URequest) {
	message, ok := request.ProtoMessage.(*zpb.SYNC_PROXY_PROPERTY)
	if !ok {
		zlog.Error("AddEngineComponent recv error request : ", proxy, request)
		return
	}
	for _, property := range message.PropertyMapList {
		proxy.SetProperty(property.Key, zattr.ReadProperty(property.Value))
		zlog.Debug("SyncProxyProperty:attr k:", property.Key)
	}
	request.Release()
}
func (service *UService) InitDownHandles() {
	service.ReqHandleMaps[zconf.MT_SYNC_PROXY_PROPERTY] = service.SyncProxyProperty
}
func (service *UService) MakeClientProxy(addr string, componentType TComponentType) *znet.UClientProxy {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		zlog.Infof(" MakeClientProxy error , addr %s , err %s", addr, err)
		return nil
	}
	proxy := znet.NewClientProxy(service, conn)
	proxy.SetProperty(zattr.Int32ComponentType, int32(componentType))
	go proxy.Serve()
	return proxy
}
func (service *UService) MakeOwnerProxy(ownerType TComponentType) *znet.UClientProxy {
	ownerProxy := service.MakeClientProxy(service.Config.OwnerAddr, ownerType)
	return ownerProxy
}
func (service *UService) InitConfig() {
	config := getServiceConfig(ComponentId)
	if ComponentId == 0 || config == nil {
		zlog.Debugf("start process failed , componentId is %d", ComponentId)
		zlog.Debugf("app path is %s", AppPath)
		os.Exit(0)
	}
	service.Config = config
	if config.LogFile != "" {
		logfile := filepath.Join(AppPath, "run/log", config.LogFile)
		zlog.SetOutput([]string{"stderr", logfile})
	}
	property := strings.Fields(config.Property)
	l := len(property)
	for i := 0; i < l; i += 3 {
		k, err := strconv.Atoi(property[i])
		if err != nil {
			continue
		}
		t := property[i+1]
		v := property[i+2]
		service.SetProperty(TEnum(k), zattr.ConvertProperty(t, v))
	}
}
