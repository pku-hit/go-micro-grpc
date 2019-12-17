package health

import (
	"context"
	"fmt"
	"github.com/pku-hit/go-micro-grpc/proto/health"
)

type HealthService struct {
}

func (hs *HealthService) Check(ctx context.Context, in *health.HealthCheckRequest, out *health.HealthCheckResponse) error {
	fmt.Printf("health checking\n")
	out.Status = health.HealthCheckResponse_SERVING
	return nil
}

func (hs *HealthService) Watch(ctx context.Context, in *health.HealthCheckRequest, ws health.Health_WatchStream) error {
	return nil
}
