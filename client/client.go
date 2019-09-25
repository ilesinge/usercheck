package client

import "net/http"

// HTTPClient interface
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// ErrNetworkFailure error
type ErrNetworkFailure struct {
	Cause error
}

func (e *ErrNetworkFailure) Error() string {
	return "An error occured during Twitter username availability check"
}

// Unwrap returns the underlying error
func (e *ErrNetworkFailure) Unwrap() error {
	return e.Cause
}
