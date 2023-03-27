package zmodel

import (
	_ "github.com/ouczbs/zmin/engine/zattr"
	_ "github.com/ouczbs/zmin/engine/zcache"
	_ "github.com/ouczbs/zmin/engine/zconf"
	_ "github.com/ouczbs/zmin/engine/zlog"
	_ "github.com/ouczbs/zmin/engine/zproto/zpb"
	_ "go.mongodb.org/mongo-driver/bson"
	_ "os"
	_ "strconv"
	_ "strings"
	"testing"
)

func TestTag(t *testing.T) {
	//service := &UService{
	//	Id: 2,
	//	Type: 1,
	//	ListenAddr: "127.0.0.1:9999",
	//}
	//zlog.Debug(service.M())
	//zcache.GetMongoClient().UpdateOrInsert(service,bson.M{"id":service.Id})
	//service.InsertOne()
	InitService()
}
