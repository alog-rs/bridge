package server

import (
	"context"

	"github.com/alog-rs/bridge/service"
	rs3pb "github.com/alog-rs/proto/rs3"
)

// RunescapeThreeServer implements methods required by rs3pb
type RunescapeThreeServer struct {
	Svc service.IRS3Svc
	rs3pb.UnimplementedRunescapeServer
}

// NewRunescapeThreeServer creates a new server backed by the provided service
func NewRunescapeThreeServer(svc service.IRS3Svc) RunescapeThreeServer {
	return RunescapeThreeServer{
		Svc: svc,
	}
}

// GetPlayerProfile returns a players profile from the backing service
func (r RunescapeThreeServer) GetPlayerProfile(ctx context.Context, in *rs3pb.GetPlayerProfileRequest) (*rs3pb.PlayerProfile, error) {
	profile, err := r.Svc.GetPlayerProfile(in.GetName(), in.GetActivityCount())

	if err != nil {
		return nil, err
	}

	return profile, nil
}
