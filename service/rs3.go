package service

import (
	"errors"

	"github.com/alog-rs/bridge/internal/helpers"
	"github.com/alog-rs/bridge/internal/types"

	rs3pb "github.com/alog-rs/proto/rs3"
)

// IRS3Svc provides methods to fetch Runescape 3 data for the server
type IRS3Svc interface {
	GetPlayerProfile(name string, activityCount int) (*rs3pb.PlayerProfile, error)
}

// RS3Svc handles fetching Runescape 3 data for the server
type RS3Svc struct {
	Req helpers.HTTPRequest
}

// NewRS3Svc creates a new RS3Svc
func NewRS3Svc(req helpers.HTTPRequest) *RS3Svc {
	return &RS3Svc{
		Req: req,
	}
}

// GetPlayerProfile fetching a players profile from a range of different JAGEX endpoints
//
// 1) RuneMetrics
// 2) Highscores lite
//
// If one of the above fails we will attempt to use the next one. They are ordered by ease-of-use
func (svc *RS3Svc) GetPlayerProfile(user string, activityCount int) (*rs3pb.PlayerProfile, error) {
	rm, rmErr := svc.fetchProfileFromRuneMetrics(user, activityCount)

	if rmErr.IsPresent() {
		return nil, errors.New("Failed to get player profile")
	}

	return rm, nil
}

func (svc *RS3Svc) fetchProfileFromRuneMetrics(user string, activityCount int) (*rs3pb.PlayerProfile, types.Error) {
	var err error

	res, err := svc.Req.GetRuneMetricsProfile(user, activityCount)

	if err != nil {
		return nil, types.ErrorRequestFailed
	}

	profile, err := types.NewRuneMetricsPlayerProfile(res)

	if err != nil {
		return nil, types.ErrorInternal
	}

	profileError := profile.GetError()

	if profileError.IsPresent() {
		return nil, profileError
	}

	pb, err := profile.ConvertToPB()

	if err != nil {
		return nil, types.ErrorInternal
	}

	return pb, types.ErrorNone
}
