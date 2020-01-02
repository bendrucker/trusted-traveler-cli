package ttapi

import "github.com/dghubble/sling"

// NewClient initializes a new Client object
func NewClient(options Options) *Client {
	client := sling.
		New().
		Base(options.URL)

	return &Client{
		Locations: newLocationService(client.New().Path("locations/")),
	}
}

// Client is an HTTP client for the Trusted Traveler API
type Client struct {
	Locations *LocationService
}

// Options holds configurable parameters for a Client
type Options struct {
	URL string
}
