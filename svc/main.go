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
)

type Data struct {
}

func (t *Data) SayHello(ctx context.Context, req *helloworld.HelloRequest, resp *helloworld.HelloReply) error {
	resp.Message = "Go say: " + req.Name
	return nil
}

func main() {
	// Use consul
	reg := consul.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{
			"127.0.0.1:8500",
		}
	})

	// Create a new service. Optionally include some options here.
	service := grpc.NewService(
		micro.Name("cloud-grpc-server"),
		micro.Registry(reg),
	)
	metadata := make(map[string]string)
	metadata["gRPC.port"] = "10086"
	service.Server().Init(server.Address("127.0.0.1:10086"), server.Metadata(metadata))
	// service.Server().Init(server.Address("127.0.0.1:10086"),)
	// service.Server().Init()

	// Init will parse the command line flags. Any flags set will
	// override the above settings. Options defined here will
	// override anything set on the command line.
	// service.Init()

	// By default we'll run the server unless the flags catch us

	// Setup the server

	// Register handler
	// proto.RegisterGreeterHandler(service.Server(), new(Greeter))
	helloworld.RegisterGreeterHandler(service.Server(), new(Data))

	// Run the server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
