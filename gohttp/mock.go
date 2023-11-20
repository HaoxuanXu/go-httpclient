package gohttp

import (
	"fmt"
	"net/http"
)

// The Mock struct provides a clean way to configure HTTP mocks based on
// the combination between request methods, URL, and request body
type Mock struct {
	Method             string
	Url                string
	RequestBody        string
	ResponseBody       string
	ResponseStatusCode int
	Error              error
}

// GetResponse returns a Response object based on the mock configuration
func (m *Mock) GetResponse() (*Response, error) {
	if m.Error != nil {
		return nil, m.Error
	}
	response := Response{
		status:     fmt.Sprintf("%d %s", m.ResponseStatusCode, http.StatusText(m.ResponseStatusCode)),
		statusCode: m.ResponseStatusCode,
		body:       []byte(m.ResponseBody),
	}
	return &response, nil
}
