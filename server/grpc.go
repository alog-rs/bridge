package server

import (
	"net"

	"google.golang.org/grpc"
)

// GRPCServer represents GRPC server which bridge uses to handle rpc calls
type GRPCServer struct {
	lis  net.Listener
	serv *grpc.Server
}

// NewGRPCServer creates a new GRPC server
func NewGRPCServer(lis net.Listener, serv *grpc.Server) *GRPCServer {
	return &GRPCServer{
		lis:  lis,
		serv: serv,
	}
}

// Serve handles serving the server
func (s *GRPCServer) Serve() error {
	return s.serv.Serve(s.lis)
}
