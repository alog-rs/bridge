package service_test

import (
	"errors"
	"testing"

	"github.com/alog-rs/bridge/internal/mocks"
	"github.com/alog-rs/bridge/service"
	rs3pb "github.com/alog-rs/proto/rs3"
	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

type Mock struct {
	req  *mocks.MockHTTPRequest
	ctrl *gomock.Controller
}

func buildMock(t *testing.T) (*service.RS3Svc, *Mock) {
	ctrl := gomock.NewController(t)
	req := mocks.NewMockHTTPRequest(ctrl)

	return service.NewRS3Svc(req), &Mock{
		req:  req,
		ctrl: ctrl,
	}
}

func TestGetPlayerProfile(t *testing.T) {
	cases := []struct {
		name       string
		mockResErr error
		mockRes    []byte
		expectErr  bool
		expected   *rs3pb.PlayerProfile
	}{
		{
			name:       "Request failed",
			mockResErr: errors.New("Mock request error"),
			mockRes:    nil,
			expectErr:  true,
		},
		{
			name:       "Malformed response",
			mockResErr: nil,
			mockRes:    []byte("err"),
			expectErr:  true,
		},
		{
			name:       "Profile not found",
			mockResErr: nil,
			mockRes:    mocks.RuneMetricsProfileResponseNotFound,
			expectErr:  true,
		},
		{
			name:       "Profile private",
			mockResErr: nil,
			mockRes:    mocks.RuneMetricsProfileResponsePrivate,
			expectErr:  true,
		},
		{
			name:       "Failed to convert to PB",
			mockResErr: nil,
			mockRes:    mocks.RuneMetricsProfileResponseSuccess,
			expectErr:  false,
			expected:   mocks.RuneMetricsPlayerProfile,
		},
		{
			name:       "Successful request",
			mockResErr: nil,
			mockRes:    mocks.RuneMetricsProfileResponseSuccess,
			expectErr:  false,
			expected:   mocks.RuneMetricsPlayerProfile,
		},
	}

	svc, mock := buildMock(t)
	for _, tc := range cases {
		mockUser := "mockuser"
		mockActivityCount := 20
		t.Run(tc.name, func(t *testing.T) {
			mock.req.EXPECT().GetRuneMetricsProfile(mockUser, mockActivityCount).Return(tc.mockRes, tc.mockResErr)

			result, err := svc.GetPlayerProfile(mockUser, mockActivityCount)

			if tc.expectErr && err == nil {
				t.Errorf("Expected error but received nil")
			}

			diff := cmp.Diff(
				tc.expected,
				result,
				cmpopts.IgnoreUnexported(rs3pb.PlayerProfile{}),
				cmpopts.IgnoreUnexported(rs3pb.QuestData{}),
				cmpopts.IgnoreUnexported(rs3pb.SkillData{}),
				cmpopts.IgnoreUnexported(rs3pb.PlayerActivityItem{}),
			)

			if !tc.expectErr && diff != "" {
				t.Errorf("GetPlayerProfile() mismatch (-want, +got):\n%s", diff)
			}
		})
	}
}
