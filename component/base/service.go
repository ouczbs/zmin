package base

import (
	"github.com/ouczbs/Zmin/engine/zattr"
	"github.com/ouczbs/Zmin/engine/zclass"
	"github.com/ouczbs/Zmin/engine/zconf"
	"github.com/ouczbs/Zmin/engine/zlog"
	"github.com/ouczbs/Zmin/engine/znet"
	"github.com/ouczbs/Zmin/engine/zproto/pb"
	"flag"
	"net"
	"os"
	"strconv"
)

type UService struct {
	*UProperty
	ReqHandleMaps map[TCmd]FRequestHandle
	MessageQueue  chan *UMessage
	Config        *UServiceConfig
}

func NewService(reqHandleMaps map[TCmd]FRequestHandle) *UService {
	return &UService{
		UProperty:     zclass.NewProperty(),
		ReqHandleMaps: reqHandleMaps,
		MessageQueue:  make(chan *UMessage, zconf.CQueueMessageSize),
		Config:        &UServiceConfig{},
	}
}
func (service *UService) ClientDisconnect(proxy *UClientProxy) {
	zlog.Debugf("ClientDisconnect %s", proxy)
	t, ok := proxy.GetProperty(zattr.Int32ComponentType).(int32)
	zlog.Infof(" ClientDisconnect   " , t)
	if ok && t == int32(pb.COMPONENT_TYPE_CENTER) {
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
	message, ok := request.ProtoMessage.(*pb.SYNC_PROXY_PROPERTY)
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
	service.SetProperty(zattr.StringCenterAddr, zconf.GetCenterConfig().ListenAddr)
	service.ReqHandleMaps[TCmd(pb.CommandList_MT_SYNC_PROXY_PROPERTY)] = service.SyncProxyProperty
}
func (service *UService) MakeClientProxy(addr string, componentType pb.COMPONENT_TYPE) *znet.UClientProxy {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		zlog.Infof(" MakeClientProxy error , addr %s , err %s", addr, err)
		return nil
	}
	proxy := znet.NewClientProxy(service, conn)
	proxy.SetProperty(zattr.Int32ComponentType , int32(componentType))
	go proxy.Serve()
	return proxy
}
func (service *UService) MakeCenterProxy() *znet.UClientProxy {
	addr, ok := service.GetProperty(zattr.StringCenterAddr).(string)
	if !ok {
		zlog.Error("ConnectToCenter :attr k:", zattr.StringListenAddr)
		return nil
	}
	centerProxy := service.MakeClientProxy(addr , pb.COMPONENT_TYPE_CENTER)
	return centerProxy
}
func (service *UService) ParseCmd() bool {
	componentId := flag.Int("ComponentId", 0, "component id")
	listenAddr := flag.String("ListenAddr", "", "listenAddr")
	flag.Parse()
	if *componentId == 0 || *listenAddr == "" {
		zlog.Debug("args not enough")
		zlog.Debug(*componentId, *listenAddr)
		return false
	}
	config := service.Config
	config.ListenAddr = *listenAddr
	config.ComponentId = int32(*componentId)
	args := flag.Args()
	l := len(args)
	for i := 0; i < l; i += 3 {
		k, err := strconv.Atoi(args[i])
		if err != nil {
			continue
		}
		t := args[i+1]
		v := args[i+2]
		service.SetProperty(TEnum(k), zattr.ConvertProperty(t, v))
	}
	zlog.Debug(flag.Args(), config)
	return true
}
func (service *UService) Run() {
	ok := service.ParseCmd()
	if !ok {
		zlog.Debugf("restart process failed")
		os.Exit(0)
	}
}
