package main

import (
	"context"
	"fmt"
	pb "github.com/pku-hit/go-micro-grpc/proto/hello"
	"github.com/pku-hit/go-micro-grpc/svc/consul"
	"github.com/pku-hit/go-micro-grpc/svc/service/hello"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
	"log"
	"net"
)

const (
	port = ":50051"
)

func RegisterToConsul() {
	consul.RegisterService("127.0.0.1:8500", &consul.ConsulService{
		IP:   "127.0.0.1",
		Port: 50051,
		Tag:  []string{"helloworld"},
		Name: "helloworld",
	})
}

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
	lis, error := net.Listen("tcp", port)
	if error != nil {
		log.Fatalf("failed to listen: %v", error)
	}
	s := grpc.NewServer()
	pb.RegisterHelloServer(s, &hello.HelloService{})
	grpc_health_v1.RegisterHealthServer(s, &HealthImpl{})
	RegisterToConsul()
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
