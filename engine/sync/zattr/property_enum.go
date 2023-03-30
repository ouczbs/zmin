package zattr

type TPropertyType = TEnum

// bytes uint32 uint64 bool -- > s i f b
const (
	Int32ComponentId = 1 + iota
	Int32ComponentType
	Int32MessageType
	StringListenAddr
	// -- service property
	BoolIsLoadedService
	StringCenterAddr
	StringComponentName
	StringLogFile
)
