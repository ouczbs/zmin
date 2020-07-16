package zconf
type (
	TSize = uint32
	TCmd = uint32
	TCallId = uint32
	TMessageType = uint16
	TSequence = uint32
	TCode = uint32
	TEnum = uint32
	TComponentId = uint32

)
type UBaseConfig struct{
	ListenAddr    string
	ComponentType uint32
	ComponentId   uint32
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