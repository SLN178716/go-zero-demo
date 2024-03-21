package myNacos

import (
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
)

// NacosConfig Nacos配置信息
type NacosConfig struct {
	Client  constant.ClientConfig
	Servers []constant.ServerConfig
	DataId  string
	Group   string
}

// Service 服务信息
type Service struct {
	ServiceName string
	Host        string
	Port        uint64
	Metadata    map[string]string
	Weight      float64
}

type Nacos struct {
	cfg     NacosConfig
	service Service
}
