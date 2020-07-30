package zcache

import (
	"github.com/ouczbs/Zmin/engine/zlog"
	"testing"
)

func TestConnect(t *testing.T) {
	redis := NewRedisClient("111.229.54.9:6379" , "ouczbs" , 0)
	redis.SetValue("framename" , "Zinx")
	name , err := redis.GetValue("framename")
	zlog.Debug(name , err)

	table , key := "loginserver" , "2"
	s1 , err := redis.SetTableValue(table , key, "f1" , "f2" , "f3" , "f4")
	zlog.Debug(s1 , err)
	sv ,err := redis.GetTableValue(table, key )
	zlog.Debug(sv , err)
	redis.SetTableField(table , key , "field" , 30)

}