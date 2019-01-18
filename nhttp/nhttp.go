package nhttp

import "net/http"

// Client is an interface that http.Client implements
type Client interface {
	Do(*http.Request) (*http.Response, error)
}
