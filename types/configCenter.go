package types

// ListenConfig 监听配置回调方法
type ListenConfig func(data string) string

// ConfigCenter 配置中心接口
type ConfigCenter interface {
	InitConfig(listenConfigCallback ListenConfig) string
}
