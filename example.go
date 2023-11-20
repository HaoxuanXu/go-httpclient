package main

import (
	"fmt"

	"github.com/HaoxuanXu/go-httpclient/gohttp"
)

var (
	githubHttpClient = getClient()
)

func getClient() gohttp.Client {
	client := gohttp.NewBuilder().
		DisableTimeouts(true).
		SetMaxIdleConnections(5).
		Build()

	return client
}

func main() {
	getUrls()
}

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func getUrls() {
	response, err := githubHttpClient.Get("https://api.github.com", nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(response.Status())
	fmt.Println(response.StatusCode())
	fmt.Println(response.Headers())
	fmt.Println(response.String())
}

// func createUser(user User) {
// 	requestBody, _ := json.Marshal(user)
// 	response, err := githubHttpClient.Post("https://api.github.com", nil, requestBody)
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println(response.StatusCode())
// 	fmt.Println(string(response.Bytes()))
// }
