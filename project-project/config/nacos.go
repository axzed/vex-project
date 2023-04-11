package config

import (
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"log"
)

// NaocsClient nacos客户端
type NacosClient struct {
	confClient config_client.IConfigClient
	group      string
}

// InitNacosClient 初始化nacos客户端
func InitNacosClient() *NacosClient {
	// 从bootstrap 初始化nacos配置
	bootConf := InitBootstrap()
	// 初始化nacos客户端
	clientConfig := constant.ClientConfig{
		NamespaceId:         bootConf.NacosConfig.Namespace, //we can create multiple clients with different namespaceId to support multiple namespace.When namespace is public, fill in the blank string here.
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		LogLevel:            "debug",
	}
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr:      bootConf.NacosConfig.IpAddr,
			ContextPath: bootConf.NacosConfig.ContextPath,
			Port:        uint64(bootConf.NacosConfig.Port),
			Scheme:      bootConf.NacosConfig.Scheme,
		},
	}
	configClient, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)
	if err != nil {
		log.Fatalln(err)
	}
	nc := &NacosClient{
		confClient: configClient,
		group:      bootConf.NacosConfig.Group,
	}
	return nc
}
