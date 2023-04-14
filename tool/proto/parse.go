package main

import (
	"encoding/json"
	"fmt"
	"github.com/ouczbs/zmin/engine/core/zlog"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func ParseJson(str string, obj interface{}) {
	if str == "" {
		return
	}
	_bytes := String2Bytes(str)
	if err := json.Unmarshal(_bytes, obj); err != nil {
		zlog.Debugf("%s error:: %s", str, err)
	}
}
func ParseGroup(doc string) *FGroupData {
	pattern := "//@group "
	line := fmt.Sprintf("%s(.*)\n", pattern)
	re := regexp.MustCompile(line)
	res := re.FindString(doc)
	if res == "" {
		return nil
	}
	data := strings.Trim(res[len(pattern):], "\r")
	var group FGroupData
	ParseJson(data, &group)
	group.Idx = int32(len(groupList))
	groupList = append(groupList, &group)
	return &group
}
func ParseProto(doc string) {
	group := ParseGroup(doc)
	pattern1, pattern2 := "//@message", "message "
	l1, l2 := len(pattern1), len(pattern2)
	line := fmt.Sprintf("%s(.*)\n%s([^{]*)", pattern1, pattern2)
	re := regexp.MustCompile(line)
	res := re.FindAllString(doc, -1)
	if res == nil {
		return
	}
	var protoList []*FProtoData
	for _, v := range res {
		lines := strings.Split(v, "\n")
		if lines == nil || len(lines) < 2 {
			continue
		}
		data := strings.Trim(lines[0][l1:], "\r")
		name := strings.Trim(lines[1][l2:], "\r")
		var message FMessageData
		ParseJson(data, &message)
		protoList = append(protoList, &FProtoData{Name: name, Message: &message, Group: group})
	}
	protoMapList = append(protoMapList, protoList)
}
func ParseAllProto(files []os.DirEntry) {
	for _, file := range files {
		name := file.Name()
		if filepath.Ext(name) != ".proto" {
			continue
		}
		zlog.Debugf("parse %s", name)
		doc := ReadFile(filepath.Join(inDir, file.Name()))
		ParseProto(doc)
	}
}
