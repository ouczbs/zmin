package main

import (
	"strings"
)

type FGroupData struct {
	Name string `json:"name"`
	Id   int32  `json:"id"`
	Idx  int32
}
type FMessageData struct {
	Id int32 `json:"id"`
}
type FProtoData struct {
	Name    string
	Message *FMessageData
	Group   *FGroupData
}
type IProtoHelp interface {
	GenProto()
}

const (
	INDIR   = "engine/sync/zpb"
	OUTFILE = "engine/data/zconf/proto.go"
)

var (
	protoHelp    IProtoHelp
	inDir        string
	outFile      string
	groupList    []*FGroupData
	protoMapList [][]*FProtoData
	bytes        strings.Builder
)

func main() {
	AppPath := MakeAppPath()
	files := GetProtoFiles(AppPath)
	ParseAllProto(files)
	protoHelp = &FGoProtoHelp{}
	protoHelp.GenProto()
	FlushFile()
}
