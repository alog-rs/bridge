package server

import (
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/alog-rs/bridge/internal/helpers"
	rs3pb "github.com/alog-rs/proto"
	"github.com/alog-rs/shared-packages/pkg/utilities"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

func handleSignals(errc chan<- error) {
	c := make(chan os.Signal, 1)

	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	err := fmt.Errorf("Encountered signal: %s", <-c)

	errc <- err
}

func startGRPCServer(errc chan<- error) {
	gRPCPort, ok := os.LookupEnv("GRPC_PORT")

	if !ok {
		errc <- errors.New("Please specify a valid GRPC_PORT")

		return
	}

	lis, err := net.Listen("tcp", ":"+gRPCPort)

	if err != nil {
		errc <- err

		return
	}

	s := grpc.NewServer()
	grpc_health_v1.RegisterHealthServer(s, NewHealthServer())
	rs3pb.RegisterRunescapeServer(s, NewRunescapeThreeServer())

	if utilities.IsDev() {
		reflection.Register(s)
	}

	server := NewGRPCServer(lis, s)

	log.Printf("Serving GRPC server from port %s\n", gRPCPort)

	errc <- server.Serve()
}

func gracefulShutdown(err error) {
	log.Println("Shutting down...", err)
}

// Initialize initializes the server and associated services
func Initialize(cmd *cobra.Command, args []string) {
	log.Printf("Starting Bridge: %s...\n", helpers.BuildVersion())

	errc := make(chan error)

	go handleSignals(errc)
	go startGRPCServer(errc)

	gracefulShutdown(<-errc)
}
