package rpc

import (
	"github.com/axzed/project-api/config"
	"github.com/axzed/project-common/discovery"
	"github.com/axzed/project-common/logs"
	"github.com/axzed/project-grpc/user/login"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
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
	conn, err := grpc.Dial("etcd:///user",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()), // 为 jaeger 加上grpc客户端支持
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	LoginServiceClient = login.NewLoginServiceClient(conn)
}
