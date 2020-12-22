package helpers

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// HTTPRequest interface implements methods which make requests to external endpoints
type HTTPRequest interface {
	GetRuneMetricsProfile(string, int) ([]byte, error)
}

// JAGEXRequest allows for making requests to JAGEX endpoints
type JAGEXRequest struct{}

// NewJAGEXRequest creates a new JAGEXRequest
func NewJAGEXRequest() *JAGEXRequest {
	return &JAGEXRequest{}
}

// GetRuneMetricsProfile is responsible for calling the /profile endpoint from RuneMetrics
func (r *JAGEXRequest) GetRuneMetricsProfile(user string, activityCount int) ([]byte, error) {
	endpoint := CreateRuneMetricsProfileEndpoint(user, activityCount)
	req, err := http.NewRequest(http.MethodGet, endpoint, nil)

	if err != nil {
		return nil, err
	}

	client := http.Client{
		Timeout: time.Second * 10,
	}

	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("Failed to call GetProfile: Receieved %d status code", res.StatusCode)
	}

	if res.Body == nil {
		return nil, errors.New("Failed to call GetProfile: Received nil body")
	}

	defer res.Body.Close()

	return ioutil.ReadAll(res.Body)
}
