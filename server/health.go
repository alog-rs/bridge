package server

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/status"
)

// HealthServer contains methods which implement the grpc_health_v1 proto
type HealthServer struct{}

// NewHealthServer creates a new HealthServer
func NewHealthServer() HealthServer {
	return HealthServer{}
}

// Check simply returns SERVER if the server is healthy
func (h HealthServer) Check(context.Context, *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {
	res := &grpc_health_v1.HealthCheckResponse{
		Status: grpc_health_v1.HealthCheckResponse_SERVING,
	}

	return res, nil
}

// Watch is unimplemented
func (h HealthServer) Watch(*grpc_health_v1.HealthCheckRequest, grpc_health_v1.Health_WatchServer) error {
	return status.Errorf(codes.Unimplemented, "health check via Watch is not implemented")
}
