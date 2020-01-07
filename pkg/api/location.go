package ttapi

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dghubble/sling"
)

// LocationService can list locations from the schedule API
type LocationService struct {
	client *sling.Sling
}

// Location is a place where Trusted Traveler services are offered
type Location struct {
	ID                  int       `json:"id" header:"ID,text"`
	Name                string    `json:"name" header:"Name"`
	ShortName           string    `json:"shortName"`
	LocationType        string    `json:"locationType"`
	LocationCode        string    `json:"locationCode"`
	Address             string    `json:"address"`
	AddressAdditional   string    `json:"addressAdditional"`
	City                string    `json:"city" header:"City"`
	State               string    `json:"state" header:"State"`
	PostalCode          string    `json:"postalCode"`
	CountryCode         string    `json:"countryCode"`
	TzData              string    `json:"tzData"`
	PhoneNumber         string    `json:"phoneNumber"`
	PhoneAreaCode       string    `json:"phoneAreaCode"`
	PhoneCountryCode    string    `json:"phoneCountryCode"`
	PhoneExtension      string    `json:"phoneExtension"`
	PhoneAltNumber      string    `json:"phoneAltNumber"`
	PhoneAltAreaCode    string    `json:"phoneAltAreaCode"`
	PhoneAltCountryCode string    `json:"phoneAltCountryCode"`
	PhoneAltExtension   string    `json:"phoneAltExtension"`
	FaxNumber           string    `json:"faxNumber"`
	FaxAreaCode         string    `json:"faxAreaCode"`
	FaxCountryCode      string    `json:"faxCountryCode"`
	FaxExtension        string    `json:"faxExtension"`
	EffectiveDate       string    `json:"effectiveDate"`
	Temporary           bool      `json:"temporary"`
	InviteOnly          bool      `json:"inviteOnly"`
	Operational         bool      `json:"operational"`
	Directions          string    `json:"directions"`
	Notes               string    `json:"notes"`
	MapFileName         string    `json:"mapFileName"`
	LastUpdatedBy       string    `json:"lastUpdatedBy"`
	LastUpdatedDate     string    `json:"lastUpdatedDate"`
	CreatedDate         string    `json:"createdDate"`
	Services            []Service `json:"services" header:"Services,text"`
}

// Zone returns a time.Location describing the time zone for the Trusted Traveler location
func (l *Location) Zone() (*time.Location, error) {
	return time.LoadLocation(l.TzData)
}

// Service is a Trusted Traveler service offered at a Location
type Service struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (s Service) String() string {
	return s.Name
}

func newLocationService(client *sling.Sling) *LocationService {
	return &LocationService{
		client: client,
	}
}

// LocationParameters represent filters on a list of locations. All fields are optional and will be omitted if nil.
type LocationParameters struct {
	InviteOnly  *bool   `url:"inviteOnly,omitempty"`
	Operational *bool   `url:"operational,omitempty"`
	ServiceName *string `url:"serviceName,omitempty"`
}

// List fetches the list of locations based on the supplied parameters
func (s *LocationService) List(params LocationParameters) ([]Location, error) {
	locations := new([]Location)

	resp, err := s.client.
		New().
		Get("").
		QueryStruct(params).
		Receive(locations, nil)

	if err != nil && resp.StatusCode != http.StatusOK {
		return *locations, fmt.Errorf("unexpected status: %s", resp.Status)
	}

	return *locations, err
}
