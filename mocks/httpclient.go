package mocks

import "net/http"

// MockHTTPClient is a mock of an HTTP client.
type MockHTTPClient struct {
	// GetCount is a custom function to be used in the Get method.
	GetFunc func(url string) (*http.Response, error)
	// GetCount is the number of times the Get method was called.
	GetCount int
}

// Get issues a GET to the specified URL.
// This method calls the `GetFunc`, if it is defined. Otherwise, it returns an empty `http.Response`.
func (m *MockHTTPClient) Get(url string) (*http.Response, error) {
	m.GetCount++

	if m.GetFunc != nil {
		return m.GetFunc(url)
	}
	return &http.Response{}, nil
}
