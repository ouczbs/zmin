package zconf
type (
	TSize = int32
	TCmd = int32
	TCallId = int32
	TMessageType = uint16
	TSequence = int32
	TCode = int32
	TEnum = int32
	TComponentId = int32

)
type UServiceConfig struct{
	ListenAddr    string
	ComponentType int32
	ComponentId   int32
}
type UCenterConfig struct{
	ListenAddr    string
	AdvertiseAddr string
	HTTPAddr      string
	LogFile       string
	LogStderr     bool
	LogLevel      string
}
// DispatcherConfig defines fields of dispatcher config
type UDispatcherConfig struct {
	ListenAddr    string
	AdvertiseAddr string
	HTTPAddr      string
	LogFile       string
	LogStderr     bool
	LogLevel      string
}
type UGateConfig struct {
	ListenAddr    string
	AdvertiseAddr string
	HTTPAddr      string
	LogFile       string
	LogStderr     bool
	LogLevel      string
}
type ULoginConfig struct {
	ListenAddr    string
	AdvertiseAddr string
	HTTPAddr      string
	LogFile       string
	LogStderr     bool
	LogLevel      string
}
type UGameConfig struct {
	ListenAddr    string
	AdvertiseAddr string
	HTTPAddr      string
	LogFile       string
	LogStderr     bool
	LogLevel      string
}
// KVDBConfig defines fields of KVDB config
type UKVDBConfig struct {
	Type       string
	Url        string // MongoDB
	DB         string // MongoDB
	Collection string // MongoDB
	Driver     string // SQL Driver: e.x. mysql
}