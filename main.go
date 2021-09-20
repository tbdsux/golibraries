package golibraries

import (
	"fmt"
	"net/http"

	"github.com/TheBoringDude/scuffed-go/requester"
	"github.com/google/go-querystring/query"
)

// API_URL is the default api server address.
const API_URL = "https://libraries.io/api/"

type LibrariesIO struct {
	ApiKey string `url:"api_key"`
	client *requester.Requester
}

type PaginateOptions struct {
	Page    int
	PerPage int
}

type QueryOptions struct {
	ApiKey  string `url:"api_key,omitempty"`
	Page    int    `url:"page,omitempty"`
	PerPage int    `url:"per_page,omitempty"`
}

// Create a new LibrariesIO instance.
func New(apiKey string) *LibrariesIO {
	return &LibrariesIO{
		ApiKey: apiKey,
		client: requester.NewRequester(&http.Client{}),
	}
}

// parses url query options
func parseQuery(v interface{}) string {
	x, err := query.Values(v)
	if err != nil {
		fmt.Println(err)
	}

	return x.Encode()
}
