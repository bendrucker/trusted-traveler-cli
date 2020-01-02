package ttapi

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLocationList(t *testing.T) {
	mux, server, client := setup()
	defer server.Close()

	mux.HandleFunc("/locations/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		_, _ = fmt.Fprint(w, `[{"id":1, "name": "Test Airport"}]`)
	})

	locations, err := client.Locations.List(LocationParameters{})

	if !assert.NoError(t, err) {
		return
	}

	if assert.Len(t, locations, 1) {
		assert.Equal(t, Location{ID: 1, Name: "Test Airport"}, locations[0])
	}
}

func TestLocationListWithParams(t *testing.T) {
	mux, server, client := setup()
	defer server.Close()

	mux.HandleFunc("/locations/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)

		assert.Equal(t, r.URL.Query().Get("inviteOnly"), "true")
		assert.Equal(t, r.URL.RequestURI(), "/locations/?inviteOnly=true")

		_, _ = fmt.Fprint(w, `[{"id":1, "name": "Secret Airport", "inviteOnly": true}]`)
	})

	locations, err := client.Locations.List(LocationParameters{
		InviteOnly: Bool(true),
	})

	if !assert.NoError(t, err) {
		return
	}

	assert.Len(t, locations, 1)
}
