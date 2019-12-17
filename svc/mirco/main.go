package main

import (
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/server"
	"github.com/micro/go-micro/service/grpc"
	"github.com/pku-hit/consul"
	"github.com/pku-hit/go-micro-grpc/proto/helloworld"
	hs "github.com/pku-hit/go-micro-grpc/svc/service/helloworld"
	"time"
)

func main() {
	reg := consul.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{
			"127.0.0.1:8500",
		}
	})

	service := grpc.NewService(
		micro.Name("helloworld"),
		micro.Registry(reg),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*15),
	)
	metadata := make(map[string]string)
	metadata["gRPC.port"] = "10086"
	service.Server().Init(server.Address(":10086"), server.Metadata(metadata))
	helloworld.RegisterGreeterHandler(service.Server(), &hs.HelloWorldService{})

	// Run the server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
