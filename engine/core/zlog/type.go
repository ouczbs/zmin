package zlog

import (
	"go.uber.org/zap/zapcore"
)

// Level is type of log levels
type Level = zapcore.Level

type FLogf = func(template string, args ...interface{})
type FLog = func(args ...interface{})

var (
	Debugf, Infof, Warnf, Errorf, Panicf, Fatalf FLogf
	Debug, Error, Panic, Fatal                   FLog
)
