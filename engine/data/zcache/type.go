package zcache

import (
	"github.com/ouczbs/zmin/engine/data/zconf"
	"sync"
)

const (
	CUnderline = "_"
)

type IModel = zconf.IModel

var once sync.Once
