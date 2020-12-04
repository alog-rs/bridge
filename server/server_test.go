package server_test

import (
	"context"
	"errors"
	"io"
	"log"
	"net"
	"testing"

	"github.com/alog-rs/bridge/internal/mocks"
	"github.com/alog-rs/bridge/server"
	rs3pb "github.com/alog-rs/proto/rs3"
	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/test/bufconn"
)

type MockRS3Server struct {
	srv  *grpc.Server
	svc  *mocks.MockIRS3Svc
	ctrl *gomock.Controller
}

const bufSize = 1024 * 1024
const mockUser = "Uss"
const mockActivityCount = 20

func runHealthServer() (*bufconn.Listener, *grpc.Server) {
	lis := bufconn.Listen(bufSize)
	srv := grpc.NewServer()

	grpc_health_v1.RegisterHealthServer(srv, server.NewHealthServer())

	go func() {
		if err := srv.Serve(lis); err != nil {
			log.Fatalf("Failed to start GRPC server %v\n", err)
		}
	}()

	return lis, srv
}

func runRS3Server(t *testing.T) (*bufconn.Listener, *MockRS3Server) {
	lis := bufconn.Listen(bufSize)
	srv := grpc.NewServer()
	ctrl := gomock.NewController(t)
	svc := mocks.NewMockIRS3Svc(ctrl)

	rs3pb.RegisterRunescapeServer(srv, server.NewRunescapeThreeServer(svc))

	go func() {
		if err := srv.Serve(lis); err != nil {
			log.Fatalf("Failed to start GRPC server %v\n", err)
		}
	}()

	return lis, &MockRS3Server{
		srv,
		svc,
		ctrl,
	}
}

func createHealthClient(ctx context.Context, lis *bufconn.Listener) (grpc_health_v1.HealthClient, error) {
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) {
		return lis.Dial()
	}), grpc.WithInsecure())

	if err != nil {
		return nil, err
	}

	return grpc_health_v1.NewHealthClient(conn), nil
}

func createRS3Client(ctx context.Context, lis *bufconn.Listener) (rs3pb.RunescapeClient, error) {
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) {
		return lis.Dial()
	}), grpc.WithInsecure())

	if err != nil {
		return nil, err
	}

	return rs3pb.NewRunescapeClient(conn), nil
}

func TestHealthServerCheck(t *testing.T) {
	lis, s := runHealthServer()
	ctx := context.Background()
	client, err := createHealthClient(ctx, lis)

	defer s.Stop()

	if err != nil {
		t.Fatalf("Failed to created client %v", err)
	}

	res, err := client.Check(ctx, &grpc_health_v1.HealthCheckRequest{})

	if err != nil {
		t.Fatalf("grpc_health_v1.Check failed %v", err)
	}

	expected := grpc_health_v1.HealthCheckResponse_SERVING

	if res.Status != expected {
		log.Fatalf("expected %d got %d", expected, res.Status)
	}
}

func TestHealthServerWatch(t *testing.T) {
	lis, s := runHealthServer()
	ctx := context.Background()
	client, err := createHealthClient(ctx, lis)

	defer s.Stop()

	if err != nil {
		t.Fatalf("Failed to create client %v", err)
	}

	stream, err := client.Watch(ctx, &grpc_health_v1.HealthCheckRequest{})

	if err != nil {
		t.Fatalf("grpc_health_v1.Watch failed %v", err)
	}

	_, err = stream.Recv()

	if err == nil {
		log.Fatalf("expected an error instead received nil")
	}

	if err == io.EOF {
		log.Fatalf("expected an error instead received empty stream")
	}

	errStatus, _ := status.FromError(err)

	if errStatus.Code() != codes.Unimplemented {
		log.Fatalf("expected %s got %s", codes.Unimplemented.String(), errStatus.Code().String())
	}
}

func TestRunescapeThreeServerGetPlayerProfile(t *testing.T) {
	lis, mock := runRS3Server(t)
	ctx := context.Background()
	client, err := createRS3Client(ctx, lis)

	if err != nil {
		t.Errorf("Failed to create client %v", err)
	}

	defer mock.srv.Stop()

	t.Run("Handles errors", func(t *testing.T) {
		mock.svc.EXPECT().GetPlayerProfile(mockUser, mockActivityCount).Times(1).Return(nil, errors.New(
			"Mock error",
		))

		_, err := client.GetPlayerProfile(ctx, &rs3pb.GetPlayerProfileRequest{
			Name:          mockUser,
			ActivityCount: mockActivityCount,
		})

		if err == nil {
			t.Error("Expected error but err is nil")
		}
	})

	t.Run("Handles success", func(t *testing.T) {
		mock.svc.EXPECT().GetPlayerProfile(mockUser, mockActivityCount).Times(1).Return(mocks.RuneMetricsPlayerProfile, nil)

		res, err := client.GetPlayerProfile(ctx, &rs3pb.GetPlayerProfileRequest{
			Name:          mockUser,
			ActivityCount: mockActivityCount,
		})

		if err != nil {
			t.Fatalf("rs3.GetPlayerProfile failed %v", err)
		}

		diff := cmp.Diff(
			mocks.RuneMetricsPlayerProfile,
			res,
			cmpopts.IgnoreUnexported(rs3pb.PlayerProfile{}),
			cmpopts.IgnoreUnexported(rs3pb.QuestData{}),
			cmpopts.IgnoreUnexported(rs3pb.SkillData{}),
			cmpopts.IgnoreUnexported(rs3pb.PlayerActivityItem{}),
		)

		if diff != "" {
			t.Errorf("GetPlayerProfile() mismatch (-want, +got):\n%s", diff)
		}
	})
}
