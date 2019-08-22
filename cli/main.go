package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/service/grpc"
	"github.com/pku-hit/consul"
	"github.com/pku-hit/go-micro-grpc/proto/helloworld"
)

func main() {
	// Use consul
	reg := consul.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{
			"127.0.0.1:8500",
		}
	})

	// Create a new service. Optionally include some options here.
	service := grpc.NewService(
		micro.Name("greeter-cli"),
		micro.Registry(reg),
	)

	service.Init()

	// Create new greeter client
	// greeter := proto.NewGreeterService("cloud-grpc-server", service.Client())
	greeter := helloworld.NewGreeterService("helloworld", service.Client())

	// Call the greeter
	rsp, err := greeter.SayHello(context.Background(), &helloworld.HelloRequest{Name: "John"})
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print response
	fmt.Println(rsp.Message)
}
