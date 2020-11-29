package server

import (
	"context"

	rs3pb "github.com/alog-rs/proto/rs3"
)

type RunescapeThreeServer struct {
	rs3pb.UnimplementedRunescapeServer
}

func NewRunescapeThreeServer() RunescapeThreeServer {
	return RunescapeThreeServer{}
}

func (r RunescapeThreeServer) GetPlayerProfile(ctx context.Context, in *rs3pb.GetPlayerProfileRequest) (*rs3pb.PlayerProfile, error) {
	return &rs3pb.PlayerProfile{}, nil
}
