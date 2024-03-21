package myNacos

import (
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-demo/types"
)

func NewNacos(cfg *NacosConfig, service *Service) types.ConfigCenter {
	return &Nacos{
		cfg:     *cfg,
		service: *service,
	}
}

func (n *Nacos) InitConfig(listenConfigCallback types.ListenConfig) string {
	// 创建服务发现客户端
	namingClient, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &n.cfg.Client,
			ServerConfigs: n.cfg.Servers,
		},
	)
	if err != nil {
		panic(err)
	}

	success, err := namingClient.RegisterInstance(vo.RegisterInstanceParam{
		Ip:          n.service.Host,
		Port:        n.service.Port,
		ServiceName: n.service.ServiceName,
		Weight:      n.service.Weight,
		Enable:      true,
		Healthy:     true,
		Ephemeral:   true,
		Metadata:    n.service.Metadata,
		ClusterName: n.cfg.Client.ClusterName, // 默认值DEFAULT
		GroupName:   n.cfg.Group,              // 默认值DEFAULT_GROUP
	})
	if err != nil {
		panic(err)
	} else if !success {
		panic("Register to nacos fail!")
	}

	// 创建动态配置客户端
	configClient, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &n.cfg.Client,
			ServerConfigs: n.cfg.Servers,
		},
	)
	if err != nil {
		panic(err)
	}

	//获取配置中心内容
	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: n.cfg.DataId,
		Group:  n.cfg.Group,
	})
	if err != nil {
		panic(err)
	}

	//设置配置监听
	if err = configClient.ListenConfig(vo.ConfigParam{
		DataId: n.cfg.DataId,
		Group:  n.cfg.Group,
		OnChange: func(namespace, group, dataId, data string) {
			//配置文件产生变化就会触发
			if len(data) == 0 {
				logx.Errorf("listen nacos data nil error ,  namespace : %s，group : %s , dataId : %s , data : %s!")
				return
			}
			listenConfigCallback(data)
		},
	}); err != nil {
		panic(err)
	}

	if len(content) == 0 {
		panic("read nacos config content err, content is nil!")
	}

	return content
}
