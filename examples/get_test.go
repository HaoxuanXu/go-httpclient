package examples

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/HaoxuanXu/go-httpclient/gohttp"
)

func TestMain(m *testing.M) {
	fmt.Println("About to start test cases for package 'examples'")

	// tell the HTTP library to mock any further request from here
	gohttp.StartMockServer()

	os.Exit(m.Run())
}

func TestGet(t *testing.T) {

	defer gohttp.StopMockServer()

	t.Run("TestErrorFetchingFromGithub", func(t *testing.T) {
		// initialization
		gohttp.FlushMocks()
		gohttp.AddMock(gohttp.Mock{
			Method: http.MethodGet,
			Url:    "https://api.github.com",
			Error:  errors.New("timeout getting github endpoints"),
		})

		// execution
		endpoints, err := GetEndpoints()

		// validation
		if endpoints != nil {
			t.Error("no endpoints expected")
		}
		if err == nil {
			t.Error("an error was expected")
		}
		if err.Error() != "timeout getting github endpoints" {
			t.Error("invalid error message received")
		}

	})

	t.Run("TestErrorUnmarshalResponseBody", func(t *testing.T) {
		// initialization
		gohttp.FlushMocks()
		gohttp.AddMock(gohttp.Mock{
			Method:             http.MethodGet,
			Url:                "https://api.github.com",
			ResponseStatusCode: http.StatusOK,
			ResponseBody:       `{"current_user_url": 123}`,
		})
		// execution
		endpoints, err := GetEndpoints()

		// validation
		if endpoints != nil {
			t.Error("no endpoints expected")
		}
		if err == nil {
			t.Error("an error was expected")
		}
		if !strings.Contains(err.Error(), "cannot unmarshal number into Go struct field") {
			t.Error("invalid error message received")
		}
	})

	t.Run("TestNoError", func(t *testing.T) {
		// initialization
		gohttp.FlushMocks()
		gohttp.AddMock(gohttp.Mock{
			Method:             http.MethodGet,
			Url:                "https://api.github.com",
			ResponseStatusCode: http.StatusOK,
			ResponseBody:       `{"current_user_url": "https://api.github.com/user"}`,
		})
		// execution
		endpoints, err := GetEndpoints()

		// validation
		if err != nil {
			t.Errorf("no error was expected, but got %v", err)
		}
		if endpoints == nil {
			t.Error("endpoint was expected but not received")
		}
	})
}
