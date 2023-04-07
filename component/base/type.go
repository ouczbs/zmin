package base

import (
	"github.com/ouczbs/zmin/engine/core/zclass"
	"github.com/ouczbs/zmin/engine/data/zconf"
	"github.com/ouczbs/zmin/engine/net/zmessage"
	"github.com/ouczbs/zmin/engine/net/znet"
)

type (
	UClientProxy   = znet.UClientProxy
	UMessage       = zmessage.UMessage
	URequest       = zmessage.URequest
	UProperty      = zclass.UProperty
	FRequestHandle = znet.FRequestHandle

	IService = znet.IService

	TCmd           = zconf.TCmd
	TEnum          = zconf.TEnum
	TComponentType = zconf.TComponentType
)

type FServiceConfig struct {
	ListenAddr    string
	OwnerAddr     string
	ComponentType int32
	ComponentId   int32
	LogFile       string
	LogLevel      string
	Property      string
}

// KVDBConfig defines fields of KVDB config
type FKVDBConfig struct {
	Url    string // MongoDB
	DB     string // MongoDB
	Driver string // SQL Driver: e.x. mysql
}
