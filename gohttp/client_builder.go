package gohttp

import (
	"net/http"
	"time"
)

type ClientBuilder interface {
	SetHeaders(http.Header) ClientBuilder
	SetConnectionTimeout(time.Duration) ClientBuilder
	SetResponseTimeout(time.Duration) ClientBuilder
	SetMaxIdleConnections(int) ClientBuilder
	DisableTimeouts(bool) ClientBuilder
	SetHttpClient(c *http.Client) ClientBuilder
	SetUserAgent(userAgent string) ClientBuilder

	Build() Client
}

type clientBuilder struct {
	headers           http.Header
	maxIdleConns      int
	connectionTimeout time.Duration
	responseTimeout   time.Duration
	disableTimeouts   bool
	client            *http.Client
	userAgent         string
}

func NewBuilder() ClientBuilder {
	builder := &clientBuilder{}
	return builder
}

func (c *clientBuilder) Build() Client {
	client := httpClient{
		builder: c,
	}
	return &client
}

func (c *clientBuilder) SetHeaders(headers http.Header) ClientBuilder {
	c.headers = headers
	return c
}

func (c *clientBuilder) SetConnectionTimeout(timeout time.Duration) ClientBuilder {
	c.connectionTimeout = timeout
	return c
}

func (c *clientBuilder) SetResponseTimeout(timeout time.Duration) ClientBuilder {
	c.responseTimeout = timeout
	return c
}

func (c *clientBuilder) SetMaxIdleConnections(connections int) ClientBuilder {
	c.maxIdleConns = connections
	return c
}

func (c *clientBuilder) DisableTimeouts(disable bool) ClientBuilder {
	c.disableTimeouts = disable
	return c
}

func (c *clientBuilder) SetHttpClient(client *http.Client) ClientBuilder {
	c.client = client
	return c
}

func (c *clientBuilder) SetUserAgent(userAgent string) ClientBuilder {
	c.userAgent = userAgent
	return c
}
