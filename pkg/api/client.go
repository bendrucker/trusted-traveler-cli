package ttapi

import "github.com/dghubble/sling"

// NewClient initializes a new Client object
func NewClient(options Options) *Client {
	client := sling.
		New().
		Base(options.URL)

	return &Client{
		Locations: newLocationService(client.New().Path("locations/")),
		Slots:     newSlotService(client.New().Path("slots/")),
	}
}

// Client is an HTTP client for the Trusted Traveler API
type Client struct {
	Locations *LocationService
	Slots     *SlotService
}

// Options holds configurable parameters for a Client
type Options struct {
	URL string
}
