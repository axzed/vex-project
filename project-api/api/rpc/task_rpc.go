package rpc

import (
	"github.com/axzed/project-api/config"
	"github.com/axzed/project-common/discovery"
	"github.com/axzed/project-common/logs"
	"github.com/axzed/project-grpc/task"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
	"log"
)

var TaskServiceClient task.TaskServiceClient

// InitTaskRpcClient 初始化Task gRPC 的客户端连接
func InitTaskRpcClient() {
	etcdRegister := discovery.NewResolver(config.AppConf.EtcdConfig.Addrs, logs.LG)
	resolver.Register(etcdRegister)
	// 注: task的服务也由project-project服务提供
	conn, err := grpc.Dial("etcd:///project", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	TaskServiceClient = task.NewTaskServiceClient(conn)
}
