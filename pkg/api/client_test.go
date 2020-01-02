package ttapi

import (
	"net/http"
	"net/http/httptest"
)

func setup() (*http.ServeMux, *httptest.Server, *Client) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	client := NewClient(Options{URL: server.URL})

	return mux, server, client
}
