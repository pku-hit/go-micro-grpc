package hello

import (
	"context"
	pb "github.com/pku-hit/go-micro-grpc/proto/hello"
	"log"
)

type HelloService struct{}

func (s *HelloService) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.Name)
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}
