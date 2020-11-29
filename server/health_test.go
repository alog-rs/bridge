package server_test

import (
	"context"
	"io"
	"log"
	"net"
	"testing"

	"github.com/alog-rs/bridge/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func init() {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()

	grpc_health_v1.RegisterHealthServer(s, server.NewHealthServer())

	serv := server.NewGRPCServer(lis, s)

	go func() {
		if err := serv.Serve(); err != nil {
			log.Fatalf("Failed to start GRPC server %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func createClient(ctx context.Context) (grpc_health_v1.HealthClient, error) {
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())

	if err != nil {
		return nil, err
	}

	return grpc_health_v1.NewHealthClient(conn), nil
}

func TestHealthServerCheck(t *testing.T) {
	ctx := context.Background()
	client, clientErr := createClient(ctx)

	if clientErr != nil {
		t.Fatalf("Failed to created client %v", clientErr)
	}

	res, checkErr := client.Check(ctx, &grpc_health_v1.HealthCheckRequest{})

	if checkErr != nil {
		t.Fatalf("grpc_health_v1.Check failed %v", checkErr)
	}

	expected := grpc_health_v1.HealthCheckResponse_SERVING

	if res.Status != expected {
		log.Fatalf("expected %d got %d", expected, res.Status)
	}
}

func TestHealthServerWatch(t *testing.T) {
	ctx := context.Background()
	client, clientErr := createClient(ctx)

	if clientErr != nil {
		t.Fatalf("Failed to create client %v", clientErr)
	}

	stream, watchErr := client.Watch(ctx, &grpc_health_v1.HealthCheckRequest{})

	if watchErr != nil {
		t.Fatalf("grpc_health_v1.Watch failed %v", watchErr)
	}

	_, resError := stream.Recv()

	if resError == nil {
		log.Fatalf("expected an error instead received nil")
	}

	if resError == io.EOF {
		log.Fatalf("expected an error instead received empty stream")
	}

	errStatus, _ := status.FromError(resError)

	if errStatus.Code() != codes.Unimplemented {
		log.Fatalf("expected %s got %s", codes.Unimplemented.String(), errStatus.Code().String())
	}
}
