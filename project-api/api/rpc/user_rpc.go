package rpc

import (
	"github.com/axzed/project-api/config"
	"github.com/axzed/project-common/discovery"
	"github.com/axzed/project-common/logs"
	"github.com/axzed/project-grpc/user/login"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
	"log"
)

var LoginServiceClient login.LoginServiceClient

// InitUserRpcClient 初始化grpc的客户端连接
func InitUserRpcClient() {
	etcdRegister := discovery.NewResolver(config.AppConf.EtcdConfig.Addrs, logs.LG)
	resolver.Register(etcdRegister)
	conn, err := grpc.Dial("etcd:///user", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	LoginServiceClient = login.NewLoginServiceClient(conn)
}
