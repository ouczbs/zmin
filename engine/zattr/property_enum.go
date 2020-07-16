package zattr
type PropertyType = TEnum

// bytes uint32 uint64 bool -- > s i f b
const (
	// -- service proxy property
	Invalid  PropertyType = 0 + iota
	Uint32ComponentId
	Uint32ComponentType
	StringListenAddr
	// -- service property
	BoolIsLoadedService
	StringCenterAddr
)