
package zmodel

import (
"github.com/ouczbs/zmin/engine/zconf"
"github.com/ouczbs/zmin/engine/zproto/zpb"
)

type UService struct {
	Id zconf.TComponentId
	Type zpb.COMPONENT_TYPE
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