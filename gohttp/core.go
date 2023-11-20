package gohttp

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"io"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/HaoxuanXu/go-httpclient/gomime"
)

const (
	DEFAULT_MAX_IDLE_CONNECTIONS = 5
	DEFAULT_RESPONSE_TIMEOUT     = 5 * time.Second
	DEFAULT_CONNECTION_TIMEOUT   = time.Second
)

func (c *httpClient) getRequestBody(contentType string, body interface{}) ([]byte, error) {
	if body == nil {
		return nil, nil
	}

	switch strings.ToLower(contentType) {
	case gomime.ContentTypeJson:
		return json.Marshal(body)
	case gomime.ContentTypeXml:
		return xml.Marshal(body)
	default:
		return json.Marshal(body)
	}
}

func (c *httpClient) getHttpClient() *http.Client {
	c.clientOnce.Do(func() {
		if c.builder.client != nil {
			c.client = c.builder.client
			return
		}
		c.client = &http.Client{
			Timeout: c.getConnectionTimeout() + c.getResponseTimeout(),
			Transport: &http.Transport{
				MaxIdleConns:          c.getMaxIdleConnections(),
				ResponseHeaderTimeout: c.getResponseTimeout(),
				DialContext: (&net.Dialer{
					Timeout: c.getConnectionTimeout(),
				}).DialContext,
			},
		}
	})

	return c.client
}

func (c *httpClient) getMaxIdleConnections() int {
	if c.builder.maxIdleConns > 0 {
		return c.builder.maxIdleConns
	}
	return DEFAULT_MAX_IDLE_CONNECTIONS
}

func (c *httpClient) getResponseTimeout() time.Duration {
	if c.builder.responseTimeout > 0 {
		return c.builder.responseTimeout
	}
	if c.builder.disableTimeouts {
		return 0
	}
	return DEFAULT_RESPONSE_TIMEOUT
}

func (c *httpClient) getConnectionTimeout() time.Duration {
	if c.builder.connectionTimeout > 0 {
		return c.builder.connectionTimeout
	}
	if c.builder.disableTimeouts {
		return 0
	}
	return DEFAULT_CONNECTION_TIMEOUT
}

func (c *httpClient) do(method, url string, headers http.Header, body interface{}) (*Response, error) {
	fullHeaders := c.getRequestHeaders(headers)

	requestBody, err := c.getRequestBody(fullHeaders.Get("Content-Type"), body)
	if err != nil {
		return nil, err
	}

	if mock := mockupServer.getMock(method, url, string(requestBody)); mock != nil {
		return mock.GetResponse()
	}

	request, err := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, errors.New("unable to create a new request")
	}

	request.Header = fullHeaders

	client := c.getHttpClient()

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	respBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	finalResponse := Response{
		status:     response.Status,
		statusCode: response.StatusCode,
		headers:    response.Header,
		body:       respBody,
	}
	return &finalResponse, nil

}
