package zattr

import (
	"github.com/ouczbs/Zmin/engine/zconf"
	"github.com/ouczbs/Zmin/engine/zlog"
	"github.com/ouczbs/Zmin/engine/znet"
	"github.com/ouczbs/Zmin/engine/zproto"
	"github.com/ouczbs/Zmin/engine/zproto/zpb"
	"strconv"
)

const (
	_MT_SYNC_PROXY_PROPERTY =  znet.TCmd(zpb.CommandList_MT_SYNC_PROXY_PROPERTY)
	_MT_SET_REMOTE_PROPERTY =  _MT_SYNC_PROXY_PROPERTY
	_MT_TO_ALL = zconf.MT_TO_ALL
)

func SetRemoteProperty(proxy * UClientProxy , key TEnum, value interface{})  {
	property := WriteProperty(value)
	if property == nil {
		return
	}
	sync := &zpb.SET_REMOTE_PROPERTY{}
	sync.PropertyMapList = append(sync.PropertyMapList, &zpb.PropertyMap{Key: key,Value: property})
	request := znet.NewRequest(_MT_SET_REMOTE_PROPERTY , _MT_TO_ALL,sync)
	zproto.SendPbMessage(proxy  , request)
}
//设置链接属性
func SyncProxyProperty(proxy * UClientProxy , key TEnum, value interface{}) {
	property := WriteProperty(value)
	if property == nil {
		return
	}
	proxy.SetProperty(key , ReadProperty(property))
	sync := &zpb.SYNC_PROXY_PROPERTY{}
	sync.PropertyMapList = append(sync.PropertyMapList, &zpb.PropertyMap{Key: key,Value: property})
	request := znet.NewRequest(_MT_SYNC_PROXY_PROPERTY , _MT_TO_ALL , sync)
	zproto.SendPbMessage(proxy , request)
}
func SyncProxyPropertyMaps(proxy * UClientProxy) {
	sync := &zpb.SYNC_PROXY_PROPERTY{}
	for key , value := range proxy.Property{
		property := WriteProperty(value)
		if property == nil {
			continue
		}
		proxy.SetProperty(key , ReadProperty(property))
		sync.PropertyMapList = append(sync.PropertyMapList, &zpb.PropertyMap{Key: key,Value: property})
	}
	request := znet.NewRequest(_MT_SYNC_PROXY_PROPERTY , _MT_TO_ALL , sync)
	zproto.SendPbMessage(proxy , request)
}
func ReadProperty(p * zpb.Property) interface{}{
	switch p.Type {
	case zpb.Property_Type_INT32:
		return p.PInt
	case zpb.Property_Type_String:
		return p.PString
	case zpb.Property_Type_BOOL:
		return p.PBool
	default:
		return p.PFloat
	}
}
func ConvertProperty(ts string , s string) interface{}{
	t,err := strconv.Atoi(ts)
	if err != nil{
		return nil
	}
	switch zpb.Property_Type(t) {
	case zpb.Property_Type_INT32:
		i,_ := strconv.ParseUint(s,10,32)
		return i
	case zpb.Property_Type_String:
		return s
	case zpb.Property_Type_BOOL:
		b,_ := strconv.ParseBool(s)
		return b
	default:
		f,_ := strconv.ParseFloat(s,32)
		return f
	}
}
func WriteProperty(v interface{})* zpb.Property{
	option := OptionOf(v)
	if option == nil {return nil}
	property := &zpb.Property{}
	option(property)
	return property
}
type FPropertyOption func(property *zpb.Property)
func WithBool(v bool)FPropertyOption{
	return func(property *zpb.Property) {
		property.PBool = v
		property.Type = zpb.Property_Type_BOOL
	}
}
func WithInt32(v int32)FPropertyOption{
	return func(property *zpb.Property) {
		property.PInt = v
		property.Type = zpb.Property_Type_INT32
	}
}
func WithInt64(v uint64)FPropertyOption{
	return func(property *zpb.Property) {
		property.PFloat = v
		property.Type = zpb.Property_Type_INT64
	}
}
func WithString(str string)FPropertyOption{
	return func(property *zpb.Property) {
		property.PString = str
		property.Type = zpb.Property_Type_String
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
	case uint32 :
		return  WithInt32(int32(v.(uint32)))
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
		zlog.Debug("OptionOf can't convert interface to property , interface type = " ,v , t)
		return nil
	}
}