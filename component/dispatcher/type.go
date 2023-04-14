package dispatcher

import (
	"github.com/ouczbs/zmin/component/base"
	"github.com/ouczbs/zmin/engine/data/zcache"
	"github.com/ouczbs/zmin/engine/data/zconf"
	"github.com/ouczbs/zmin/engine/data/zmodel"
	"github.com/ouczbs/zmin/engine/net/zmessage"
	"github.com/ouczbs/zmin/engine/net/znet"
)

type (
	UClientProxy   = znet.UClientProxy
	UMessage       = zmessage.UMessage
	URequest       = zmessage.URequest
	UService       = base.UService
	FRequestHandle = znet.FRequestHandle
	UPacket        = zmessage.UPacket
	TCmd           = zconf.TCmd
	TMessageType   = zconf.TMessageType
	TComponentType = zconf.TComponentType
	TComponentId   = zconf.TComponentId
	FRole          = zmodel.FRole
)

var (
	ownerType     = TComponentType(zconf.COMPONENT_TYPE_VERSION)
	reqHandleMaps = make(map[TCmd]FRequestHandle)
	gameTypeMaps  = make(map[TMessageType]*UClientProxy)
	gameSceneMaps = make(map[TComponentId]*UClientProxy)
	gateMaps      = make(map[TComponentId]*UClientProxy)
	roleMaps      = make(map[TComponentId]*FRole)
)

func init() {
	mongodb := base.ServiceConfigFile.MongoDB
	zcache.InitMongoClient(mongodb.Url, mongodb.DB)
}
