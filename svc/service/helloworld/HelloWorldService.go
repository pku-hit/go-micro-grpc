package helloworld

import (
	"context"
	"github.com/pku-hit/go-micro-grpc/proto/helloworld"
)

type HelloWorldService struct {
}

func (t *HelloWorldService) SayHello(ctx context.Context, req *helloworld.HelloRequest, resp *helloworld.HelloReply) error {
	resp.Message = "Go say: " + req.Name
	return nil
}
