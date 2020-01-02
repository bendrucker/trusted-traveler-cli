package ttapi

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSlotList(t *testing.T) {
	mux, server, client := setup()
	defer server.Close()

	mux.HandleFunc("/slots/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		_, _ = fmt.Fprint(w, `[{"locationId":1, "startTimestamp" : "2020-01-01T00:00"}]`)
	})

	slots, err := client.Slots.List(SlotParameters{})

	if !assert.NoError(t, err) {
		return
	}

	assert.Len(t, slots, 1)
}

func TestSlotListWithParams(t *testing.T) {
	mux, server, client := setup()
	defer server.Close()

	mux.HandleFunc("/slots/", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)

		assert.Equal(t, r.URL.Query().Get("locationId"), "123")
		assert.Equal(t, r.URL.RequestURI(), "/slots/?locationId=123")

		_, _ = fmt.Fprint(w, `[{"locationId":123, "startTimestamp" : "2020-01-01T00:00"}]`)
	})

	locations, err := client.Slots.List(SlotParameters{
		LocationID: Int(123),
	})

	if !assert.NoError(t, err) {
		return
	}

	assert.Len(t, locations, 1)
}
