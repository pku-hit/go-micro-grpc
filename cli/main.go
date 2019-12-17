package main

import (
	"context"
	"github.com/pku-hit/go-micro-grpc/cli/consul"
	pb "github.com/pku-hit/go-micro-grpc/proto/hello"
	"google.golang.org/grpc"
	"log"
	"os"
	"time"
)

/*func main() {
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
}*/

const (
	target      = "consul://127.0.0.1:8500/helloworld"
	defaultName = "world"
)

func main() {
	consul.Init()

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	conn, err := grpc.DialContext(ctx, target, grpc.WithBlock(), grpc.WithInsecure(), grpc.WithBalancerName("round_robin"))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewHelloClient(conn)

	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[0]
	}
	for {
		ctx, _ := context.WithTimeout(context.Background(), time.Second)
		r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		log.Printf("Greeting: %s", r.Message)
		time.Sleep(time.Second * 2)
	}
}
