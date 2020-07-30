
package zmodel

import (
"Zmin/engine/zconf"
"Zmin/engine/zproto/pb"
)

type UService struct {
	Id zconf.TComponentId
	Type pb.COMPONENT_TYPE
	ListenAddr string
	Path string
	Property string
}

func (service * UService)Table()string{
	return "service"
}

func (service * UService) M()map[string]interface{}{
	return M(service)
}
var Service = &UService{}
func init(){
	Schema(Service)
}