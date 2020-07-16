package zlog

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"strings"
)

// Level is type of log levels
type Level = zapcore.Level

var (
	cfg          zap.Config
	logger       *zap.Logger
	sugar        *zap.SugaredLogger
	source       string
	currentLevel Level
)
type FLogf = func(template string, args ...interface{})
type FLog = func(args ...interface{})
var (
	Debugf ,Infof ,Warnf ,Errorf , Panicf ,Fatalf FLogf
	Debug ,Error ,Panic ,Fatal  FLog
)
func init() {
	currentLevel = zap.DebugLevel
	cfg = zap.NewDevelopmentConfig()
	cfg.Development = true
	rebuildLoggerFromCfg()
}

// SetSource sets the component name (dispatcher/gate/game) of gwlog module
func SetSource(source_ string) {
	source = source_
	rebuildLoggerFromCfg()
}

func SetParseLevel(slv string){
	lv := ParseLevel(slv)
	SetLevel(lv)
}
// SetLevel sets the log level
func SetLevel(lv Level) {
	currentLevel = lv
	cfg.Level.SetLevel(lv)
}

// GetLevel get the current log level
func GetLevel() Level {
	return currentLevel
}

// SetOutput sets the output writer
func SetOutput(outputs []string) {
	cfg.OutputPaths = outputs
	rebuildLoggerFromCfg()
}

// ParseLevel converts string to Levels
func ParseLevel(s string) Level {
	if strings.ToLower(s) == "debug" {
		return zap.DebugLevel
	} else if strings.ToLower(s) == "info" {
		return zap.InfoLevel
	} else if strings.ToLower(s) == "warn" || strings.ToLower(s) == "warning" {
		return zap.WarnLevel
	} else if strings.ToLower(s) == "error" {
		return zap.ErrorLevel
	} else if strings.ToLower(s) == "panic" {
		return zap.PanicLevel
	} else if strings.ToLower(s) == "fatal" {
		return zap.FatalLevel
	}
	Errorf("ParseLevel: unknown level: %s", s)
	return zap.DebugLevel
}

func rebuildLoggerFromCfg() {
	newLogger, err := cfg.Build();
	if err != nil {
		panic(err);return
	}
	if logger != nil {
		logger.Sync()
	}
	logger = newLogger
	if source != "" {
		logger = logger.With(zap.String("source", source))
	}
	sugar = logger.Sugar()
	initFLog()
}
func initFLog(){
	Debugf = sugar.Debugf
	Infof = sugar.Infof
	Warnf = sugar.Warnf
	Errorf = sugar.Errorf
	Panicf = sugar.Panicf
	Fatalf = sugar.Fatalf

	Debug = sugar.Debug
	Error = sugar.Error
	Panic = sugar.Panic
	Fatal = sugar.Fatal
}