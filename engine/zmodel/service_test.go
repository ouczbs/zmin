package zmodel

import (
	"github.com/ouczbs/Zmin/engine/zattr"
	"github.com/ouczbs/Zmin/engine/zcache"
	"github.com/ouczbs/Zmin/engine/zconf"
	"github.com/ouczbs/Zmin/engine/zlog"
	"github.com/ouczbs/Zmin/engine/zproto/pb"
	"go.mongodb.org/mongo-driver/bson"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestTag(t *testing.T) {
	//service := &UService{
	//	Id: 2,
	//	Type: 1,
	//	ListenAddr: "127.0.0.1:9999",
	//}
	//zlog.Debug(service.M())
	//zcache.MongoClient.UpdateOrInsert(service,bson.M{"id":service.Id})
	//service.InsertOne()
	InitService()
}
