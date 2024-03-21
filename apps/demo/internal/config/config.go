package config

import (
	"github.com/zeromicro/go-zero/rest"
	"go-zero-demo/middleware/myNacos"
)

type Config struct {
	rest.RestConf
	NacosConfig myNacos.NacosConfig `yaml:"Nacos"`
}
