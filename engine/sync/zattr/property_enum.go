package zattr

type TPropertyType = TEnum

// bytes uint32 uint64 bool -- > s i f b
const (
	serviceIndex = iota
	messageIndex
)

const (
	//proxy
	Int32ComponentId = serviceIndex*256 + iota
	Int32ComponentType
	StringListenAddr
	BoolIsOwnerProxy
	BoolIsPlayerProxy
)

const (
	// message
	Int32MessageType = 256*messageIndex + iota
	// -- service property
)
