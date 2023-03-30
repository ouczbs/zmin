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

type UServiceConfig struct {
	ListenAddr    string
	ComponentType int32
	ComponentId   int32
}
type UCenterConfig struct {
	ListenAddr    string
	AdvertiseAddr string
	HTTPAddr      string
	LogFile       string
	LogStderr     bool
	LogLevel      string
}

// DispatcherConfig defines fields of dispatcher config
type UDispatcherConfig struct {
	ListenAddr    string
	AdvertiseAddr string
	HTTPAddr      string
	LogFile       string
	LogStderr     bool
	LogLevel      string
}
type UGateConfig struct {
	ListenAddr    string
	AdvertiseAddr string
	HTTPAddr      string
	LogFile       string
	LogStderr     bool
	LogLevel      string
}
type ULoginConfig struct {
	ListenAddr    string
	AdvertiseAddr string
	HTTPAddr      string
	LogFile       string
	LogStderr     bool
	LogLevel      string
}
type UGameConfig struct {
	ListenAddr    string
	AdvertiseAddr string
	HTTPAddr      string
	LogFile       string
	LogStderr     bool
	LogLevel      string
}

// KVDBConfig defines fields of KVDB config
type UKVDBConfig struct {
	Type       string
	Url        string // MongoDB
	DB         string // MongoDB
	Collection string // MongoDB
	Driver     string // SQL Driver: e.x. mysql
}
