package hello

import (
	"context"
	"github.com/pku-hit/go-micro-grpc/proto/helloworld"
)

type HelloService struct {
}

func (t *HelloService) SayHello(ctx context.Context, req *helloworld.HelloRequest, resp *helloworld.HelloReply) error {
	resp.Message = "Go say: " + req.Name
	return nil
}
