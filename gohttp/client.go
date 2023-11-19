package gohttp

import (
	"net/http"
	"sync"
)

type Client interface {
	Get(string, http.Header) (*Response, error)
	Post(string, http.Header, interface{}) (*Response, error)
	Put(string, http.Header, interface{}) (*Response, error)
	Patch(string, http.Header, interface{}) (*Response, error)
	Delete(string, http.Header) (*Response, error)
}

type httpClient struct {
	builder    *clientBuilder
	client     *http.Client
	clientOnce sync.Once
}

func (h *httpClient) Get(url string, headers http.Header) (*Response, error) {
	return h.do(http.MethodGet, url, headers, nil)
}
func (h *httpClient) Post(url string, headers http.Header, body interface{}) (*Response, error) {
	return h.do(http.MethodPost, url, headers, body)
}

func (h *httpClient) Put(url string, headers http.Header, body interface{}) (*Response, error) {
	return h.do(http.MethodPut, url, headers, body)
}

func (h *httpClient) Patch(url string, headers http.Header, body interface{}) (*Response, error) {
	return h.do(http.MethodPatch, url, headers, body)
}

func (h *httpClient) Delete(url string, headers http.Header) (*Response, error) {
	return h.do(http.MethodDelete, url, headers, nil)
}
