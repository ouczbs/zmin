package zconf
func GetCenterConfig()*UCenterConfig{
	return &UCenterConfig{
		ListenAddr:"127.0.0.1:9999",
		HTTPAddr:"127.0.0.1:9999",
		LogFile:"center.log",
		LogLevel:"debug",
	}
}
func GetLoginConfig()*ULoginConfig{
	return &ULoginConfig{
		ListenAddr:"127.0.0.1:11001",
		HTTPAddr:"127.0.0.1:11001",
		LogFile:"dispatcher.log",
		LogLevel:"debug",
	}
}
func GetGateConfig()*UGateConfig{
	return  &UGateConfig{
		ListenAddr:"127.0.0.1:13001",
		HTTPAddr:"127.0.0.1:13001",
		LogFile:"dispatcher.log",
		LogLevel:"debug",
	}
}
func GetDispatcherConfig()*UDispatcherConfig{
	return &UDispatcherConfig{
		ListenAddr:"127.0.0.1:12001",
		HTTPAddr:"127.0.0.1:12001",
		LogFile:"dispatcher.log",
		LogLevel:"debug",
	}
}
func GetGameConfig() *UGameConfig {
	return &UGameConfig{
		ListenAddr:"127.0.0.1:14001",
		HTTPAddr:"127.0.0.1:14001",
		LogFile:"dispatcher.log",
		LogLevel:"debug",
	}
}
var Database =&UKVDBConfig{
	Url: "mongodb://111.229.54.9:27017/tree",
	DB: "tree",
	Collection: "__kv__",
}