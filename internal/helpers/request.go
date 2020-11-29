package helpers

import (
	"io/ioutil"
	"net/http"
	"time"
)

// HTTPRequest interface implements methods which make requests to external endpoints
type HTTPRequest interface {
	Get(string) ([]byte, error)
}

// JAGEXRequest holds information used to make requests to JAGEX endpoints
type JAGEXRequest struct {
	Client  http.Client
	Headers http.Header
}

// NewJAGEXRequest creates a new JAGEXRequest
func NewJAGEXRequest(client http.Client, headers http.Header) *JAGEXRequest {
	return &JAGEXRequest{
		Client:  client,
		Headers: headers,
	}
}

// DefaultClient returns a default client to use in requests
func DefaultClient() http.Client {
	return http.Client{
		// TODO: Test this
		Timeout: time.Second * 10,
	}
}

// Get make a new GET request to the provided endpoint
func (r *JAGEXRequest) Get(endpoint string) ([]byte, error) {
	var err error = nil
	req, err := http.NewRequest(http.MethodGet, endpoint, nil)

	if err != nil {
		return nil, err
	}

	req.Header = r.Headers

	res, err := r.Client.Do(req)

	if err != nil {
		return nil, err
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	return ioutil.ReadAll(res.Body)
}
