package service

import (
	"fmt"
	"log"
	"strconv"

	"github.com/alog-rs/bridge/internal/helpers"

	rs3pb "github.com/alog-rs/proto/rs3"
)

// IRS3Svc provides methods to fetch Runescape 3 data for the server
type IRS3Svc interface {
	GetPlayerProfile(name string, activityCount uint32) (*rs3pb.PlayerProfile, error)
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
// 3) Highscores
//
// If one of the above fails we will attempt the rest in the order. They are ordered by
// amount of data they return, the ease of use, and how "supported" they are for external
// use
func (svc *RS3Svc) GetPlayerProfile(name string, activityCount uint32) (*rs3pb.PlayerProfile, error) {
	endpoint := fmt.Sprintf(helpers.RunemetricsProfileEndpoint, name, strconv.FormatUint(uint64(activityCount), 10))

	log.Print(string(endpoint))

	res, err := svc.Req.Get(endpoint)

	if err != nil {
		log.Fatal("Failed to make request")
	}

	log.Print(string(res))

	return &rs3pb.PlayerProfile{}, nil
}
