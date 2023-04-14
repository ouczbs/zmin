package main

import (
	"fmt"
	"sort"
)

type FGoProtoHelp struct{}

func (g *FGoProtoHelp) GenProto() {
	g.GenPkgHeader()
	g.SortProto()
	g.GenProtoGroup()
	g.GenProtoCMDConsts()
	g.GenProtoCmd2Pb()
}
func (g *FGoProtoHelp) SortProto() {
	sort.SliceStable(groupList, func(i, j int) bool {
		if groupList[i].Id != groupList[j].Id {
			return groupList[i].Id < groupList[j].Id //从小到大排列
		}
		return groupList[i].Idx < groupList[j].Idx //从小到大排列
	})
}

func (g *FGoProtoHelp) GenPkgHeader() {
	bytes.WriteString("package zconf\r\n\r\n")
}
func (g *FGoProtoHelp) GenProtoGroup() {
	bytes.WriteString("const (\r\n")
	var line string
	for k, v := range groupList {
		if k == 0 {
			line = fmt.Sprintf("\tCMD_%s = iota * 256 \r\n", v.Name)
		} else if v.Id == 0 {
			line = fmt.Sprintf("\tCMD_%s \r\n", v.Name)
		} else {
			line = fmt.Sprintf("\tCMD_%s = %d * 256\r\n", v.Name, v.Id)
		}
		bytes.WriteString(line)
	}
	bytes.WriteString(")\r\n")
}
func (g *FGoProtoHelp) GenProtoCMDConsts() {
	bytes.WriteString("\r\n")
	for _, protoList := range protoMapList {
		g.GenProtoCMD2Int(protoList)
	}
	bytes.WriteString("\r\n")
}
func (g *FGoProtoHelp) GenProtoCMD2Int(protoList []*FProtoData) {
	bytes.WriteString("const (\r\n")
	var line string
	for k, v := range protoList {
		gName, id := v.Group.Name, v.Message.Id
		if k == 0 {
			line = fmt.Sprintf("\tCMD_%s = CMD_%s + iota\r\n", v.Name, gName)
		} else if id == 0 {
			line = fmt.Sprintf("\tCMD_%s \r\n", v.Name)
		} else {
			line = fmt.Sprintf("\tCMD_%s = CMD_%s + %d\r\n", v.Name, gName, id)
		}
		bytes.WriteString(line)
	}
	bytes.WriteString(")\r\n")
}
func (g *FGoProtoHelp) GenProtoCmd2Pb() {
	bytes.WriteString("var (\r\n\tCommandList_name = map[int32]string{\r\n")
	for _, protoList := range protoMapList {
		for _, v := range protoList {
			line := fmt.Sprintf("\t\tCMD_%s : \"zpb.%s\",\r\n", v.Name, v.Name)
			bytes.WriteString(line)
		}
	}
	bytes.WriteString("\t}\r\n)\r\n")
}
