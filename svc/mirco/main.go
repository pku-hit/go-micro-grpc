package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/server"
	"github.com/micro/go-micro/service/grpc"
	"github.com/pku-hit/consul"
	"github.com/pku-hit/go-micro-grpc/proto/helloworld"
	hs "github.com/pku-hit/go-micro-grpc/svc/service/helloworld"
	"google.golang.org/grpc/health/grpc_health_v1"
)

// health
type HealthImpl struct{}

// Check 实现健康检查接口，这里直接返回健康状态，这里也可以有更复杂的健康检查策略，比如根据服务器负载来返回
func (h *HealthImpl) Check(ctx context.Context, req *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {
	fmt.Printf("health checking\n")
	return &grpc_health_v1.HealthCheckResponse{
		Status: grpc_health_v1.HealthCheckResponse_SERVING,
	}, nil
}

func (h *HealthImpl) Watch(req *grpc_health_v1.HealthCheckRequest, w grpc_health_v1.Health_WatchServer) error {
	return nil
}

func main() {
	reg := consul.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{
			"127.0.0.1:8500",
		}
	})

	service := grpc.NewService(
		micro.Name("helloworld"),
		micro.Registry(reg),
	)
	metadata := make(map[string]string)
	metadata["gRPC.port"] = "10086"
	service.Server().Init(server.Address(":10086"), server.Metadata(metadata))
	helloworld.RegisterGreeterHandler(service.Server(), &hs.HelloWorldService{})
	// 添加健康检查
	grpc_health_v1.RegisterHealthServer(service, &HealthImpl{})

	// Run the server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
