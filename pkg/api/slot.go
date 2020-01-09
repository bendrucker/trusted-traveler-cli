package ttapi

import (
	"fmt"
	"net/http"

	"github.com/dghubble/sling"
)

// SlotService can list appointment slots from the schedule API
type SlotService struct {
	client *sling.Sling
}

// Slot is an available interview appointment
type Slot struct {
	LocationID int    `json:"locationId" header:"Location,text"`
	Start      string `json:"startTimestamp" header:"Start"`
	End        string `json:"endTimestamp" header:"End"`
	Active     bool   `json:"active"`
	Duration   int    `json:"duration" header:"Duration"`
}

func newSlotService(client *sling.Sling) *SlotService {
	return &SlotService{
		client: client,
	}
}

// SlotParameters represent filters on a list of slots. All fields are optional and will be omitted if nil.
type SlotParameters struct {
	OrderBy    *string `url:"orderBy,omitempty"`
	Limit      *int    `url:"limit,omitempty"`
	LocationID *string `url:"locationId,omitempty"`
	Minimum    *int    `url:"minimum,omitempty"`
}

// List fetches the list of slots based on the supplied parameters
func (s *SlotService) List(params SlotParameters) ([]Slot, error) {
	slots := new([]Slot)

	resp, err := s.client.
		New().
		Get("").
		QueryStruct(params).
		Receive(slots, nil)

	if err != nil && resp.StatusCode != http.StatusOK {
		return *slots, fmt.Errorf("unexpected status: %s", resp.Status)
	}

	return *slots, err
}
