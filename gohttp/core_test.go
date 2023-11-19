package gohttp

import (
	"net/http"
	"testing"
)

func TestGetRequestHeader(t *testing.T) {

	// Initialization
	client := httpClient{builder: &clientBuilder{}}
	commonHeaders := make(http.Header)
	commonHeaders.Set("Content-Type", "application/json")
	commonHeaders.Set("User-Agent", "cool-http-client")
	client.builder.headers = commonHeaders

	// Execution
	requestHeaders := make(http.Header)
	requestHeaders.Set("X-Request-Id", "ABC-123")
	finalHeaders := client.getRequestHeaders(requestHeaders)

	// Validation
	if len(finalHeaders) != 3 {
		t.Errorf("we expected 3 headers, got %d", len(finalHeaders))
	}
	if finalHeaders.Get("Content-Type") != "application/json" {
		t.Error("invalid content type received")
	}
	if finalHeaders.Get("User-Agent") != "cool-http-client" {
		t.Error("invalid user agent received")
	}
	if finalHeaders.Get("X-Request-Id") != "ABC-123" {
		t.Error("invalid request id received")
	}

}

func TestGetRequestBody(t *testing.T) {
	client := httpClient{}
	t.Run("noBodyNilResponse", func(t *testing.T) {
		body, err := client.getRequestBody("", nil)

		if err != nil {
			t.Errorf("error not expected but received: %v", err)
		}
		if body != nil {
			t.Errorf("nil body expected but received: %v", body)
		}
	})
	t.Run("jsonBodyResponse", func(t *testing.T) {
		requestBody := []string{"one", "two"}
		body, err := client.getRequestBody("application/json", requestBody)
		if err != nil {
			t.Errorf("error not expected when marshalling slice to json, but received: %v", err)
		}
		if string(body) != `["one","two"]` {
			t.Error("invalid json body obtained")
		}

	})
	t.Run("jsonBodyResponseDefault", func(t *testing.T) {
		requestBody := []string{"one", "two"}
		body, err := client.getRequestBody("", requestBody)
		if err != nil {
			t.Errorf("error not expected when marshalling slice to json, but received: %v", err)
		}
		if string(body) != `["one","two"]` {
			t.Errorf("invalid json body obtained, expected [\"one\", \"two\"], got %s", string(body))
		}
	})
	t.Run("xmlBodyResponse", func(t *testing.T) {
		requestBody := []string{"one", "two"}
		body, err := client.getRequestBody("application/xml", requestBody)
		if err != nil {
			t.Errorf("error not expected when marshalling slice to xml, but received: %v", err)
		}
		if string(body) != `<string>one</string><string>two</string>` {
			t.Errorf("invalid xml body obtained, expected [\"one\", \"two\"], got %s", string(body))
		}
	})
}
