package zattr

type TPropertyType = TEnum

// bytes uint32 uint64 bool -- > s i f b
const (
	// -- service proxy property
	Invalid TPropertyType = 0 + iota
	Int32ComponentId
	Int32ComponentType
	Int32MessageType
	StringListenAddr
	// -- service property
	BoolIsLoadedService
	StringCenterAddr
	StringComponentName
	StringLogFile
)
