package zattr

import (
	"github.com/ouczbs/zmin/engine/core/zlog"
	"github.com/ouczbs/zmin/engine/data/zconf"
	"github.com/ouczbs/zmin/engine/net/zmessage"
	"github.com/ouczbs/zmin/engine/sync/zpb"
	"github.com/ouczbs/zmin/engine/sync/zproto"
	"strconv"
)

const (
	_MT_SYNC_PROXY_PROPERTY = zconf.MT_SYNC_PROXY_PROPERTY
	_MT_SET_REMOTE_PROPERTY = _MT_SYNC_PROXY_PROPERTY
	_MT_TO_ALL              = zconf.MT_TO_ALL
)
const (
	Property_Type_INT32 = 1 + iota
	Property_Type_String
	Property_Type_INT64
	Property_Type_BOOL
)

func SetRemoteProperty(proxy *UClientProxy, key TEnum, value interface{}) {
	property := WriteProperty(value)
	if property == nil {
		return
	}
	sync := &zpb.SET_REMOTE_PROPERTY{}
	sync.PropertyMapList = append(sync.PropertyMapList, &zpb.PropertyMap{Key: key, Value: property})
	request := zmessage.NewRequest(_MT_SET_REMOTE_PROPERTY, _MT_TO_ALL, sync)
	zproto.SendPbMessage(proxy, request)
}

// 设置链接属性
func SyncProxyProperty(proxy *UClientProxy, key TEnum, value interface{}) {
	property := WriteProperty(value)
	if property == nil {
		return
	}
	proxy.SetProperty(key, ReadProperty(property))
	sync := &zpb.SYNC_PROXY_PROPERTY{}
	sync.PropertyMapList = append(sync.PropertyMapList, &zpb.PropertyMap{Key: key, Value: property})
	request := zmessage.NewRequest(_MT_SYNC_PROXY_PROPERTY, _MT_TO_ALL, sync)
	zproto.SendPbMessage(proxy, request)
}
func SyncProxyPropertyMaps(proxy *UClientProxy) {
	sync := &zpb.SYNC_PROXY_PROPERTY{}
	for key, value := range proxy.Property {
		property := WriteProperty(value)
		if property == nil {
			continue
		}
		proxy.SetProperty(key, ReadProperty(property))
		sync.PropertyMapList = append(sync.PropertyMapList, &zpb.PropertyMap{Key: key, Value: property})
	}
	request := zmessage.NewRequest(_MT_SYNC_PROXY_PROPERTY, _MT_TO_ALL, sync)
	zproto.SendPbMessage(proxy, request)
}
func ReadProperty(p *zpb.Property) interface{} {
	switch p.Type {
	case Property_Type_INT32:
		return p.PInt
	case Property_Type_BOOL:
		return p.PBool
	case Property_Type_INT64:
		return p.PFloat
	default:
		return p.PString
	}
}
func ConvertProperty(ts string, s string) interface{} {
	t, err := strconv.Atoi(ts)
	if err != nil {
		return nil
	}
	switch t {
	case Property_Type_INT32:
		i, _ := strconv.ParseUint(s, 10, 32)
		return i
	case Property_Type_BOOL:
		b, _ := strconv.ParseBool(s)
		return b
	case Property_Type_String:
		f, _ := strconv.ParseFloat(s, 32)
		return f
	default:
		return s
	}
}
func WriteProperty(v interface{}) *zpb.Property {
	option := OptionOf(v)
	if option == nil {
		return nil
	}
	property := &zpb.Property{}
	option(property)
	return property
}

type FPropertyOption func(property *zpb.Property)

func WithBool(v bool) FPropertyOption {
	return func(property *zpb.Property) {
		property.PBool = v
		property.Type = Property_Type_BOOL
	}
}
func WithInt32(v int32) FPropertyOption {
	return func(property *zpb.Property) {
		property.PInt = v
		property.Type = Property_Type_INT32
	}
}
func WithInt64(v uint64) FPropertyOption {
	return func(property *zpb.Property) {
		property.PFloat = v
		property.Type = Property_Type_INT64
	}
}
func WithString(str string) FPropertyOption {
	return func(property *zpb.Property) {
		property.PString = str
		property.Type = Property_Type_String
	}
}
func OptionOf(v interface{}) FPropertyOption {
	switch t := v.(type) {
	case bool:
		return WithBool(v.(bool))
	case string:
		return WithString(v.(string))
	case []byte:
		return WithString(string(v.([]byte)))
	case uint32:
		return WithInt32(int32(v.(uint32)))
	case int32:
		return WithInt32(v.(int32))
	case int:
		return WithInt32(int32(v.(int)))
	case float32:
		return WithInt32(int32(v.(float32)))
	case float64:
		return WithInt64(uint64(v.(float64)))
	case int64:
		return WithInt64(uint64(v.(int64)))
	case uint64:
		return WithInt64(v.(uint64))
	default:
		zlog.Debug("OptionOf can't convert interface to property , interface type = ", v, t)
		return nil
	}
}
