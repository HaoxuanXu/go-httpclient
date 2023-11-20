package gohttp

import (
	"net/http"
	"sync"
)

type Client interface {
	Get(url string, headers ...http.Header) (*Response, error)
	Post(url string, body interface{}, headers ...http.Header) (*Response, error)
	Put(url string, body interface{}, headers ...http.Header) (*Response, error)
	Patch(url string, body interface{}, headers ...http.Header) (*Response, error)
	Delete(url string, headers ...http.Header) (*Response, error)
	Options(url string, headers ...http.Header) (*Response, error)
}

type httpClient struct {
	builder    *clientBuilder
	client     *http.Client
	clientOnce sync.Once
}

func (h *httpClient) Get(url string, headers ...http.Header) (*Response, error) {
	requestHeaders := getHeaders(headers...)
	return h.do(http.MethodGet, url, requestHeaders, nil)
}
func (h *httpClient) Post(url string, body interface{}, headers ...http.Header) (*Response, error) {
	requestHeaders := getHeaders(headers...)
	return h.do(http.MethodPost, url, requestHeaders, body)
}

func (h *httpClient) Put(url string, body interface{}, headers ...http.Header) (*Response, error) {
	requestHeaders := getHeaders(headers...)
	return h.do(http.MethodPut, url, requestHeaders, body)
}

func (h *httpClient) Patch(url string, body interface{}, headers ...http.Header) (*Response, error) {
	requestHeaders := getHeaders(headers...)
	return h.do(http.MethodPatch, url, requestHeaders, body)
}

func (h *httpClient) Delete(url string, headers ...http.Header) (*Response, error) {
	requestHeaders := getHeaders(headers...)
	return h.do(http.MethodDelete, url, requestHeaders, nil)
}

func (h *httpClient) Options(url string, headers ...http.Header) (*Response, error) {
	requestHeaders := getHeaders(headers...)
	return h.do(http.MethodOptions, url, requestHeaders, nil)
}
