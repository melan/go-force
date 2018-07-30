package sobjects

import (
	"strings"
)

type Address struct {
	Accuracy    string  `json:"geocodeAccuracy,omitempty"`
	City        string  `json:"city,omitempty"`
	Country     string  `json:"country,omitempty"`
	CountryCode string  `json:"countryCode,omitempty"`
	Latitude    float64 `json:"latitude,omitempty"`
	Longitude   float64 `json:"longitude,omitempty"`
	PostalCode  string  `json:"postalCode,omitempty"`
	State       string  `json:"state,omitempty"`
	StateCode   string  `json:"stateCode,omitempty"`
	Street      string  `json:"street,omitempty"`
}

func (a *Address) String() string {
	var address = make([]string, 0, 10)
	if a.Street != "" {
		address = append(address, a.Street)
	}

	if a.City != "" {
		address = append(address, a.City)
	}

	if a.State != "" {
		address = append(address, a.State)
	} else if a.StateCode != "" {
		address = append(address, a.StateCode)
	}

	if a.Country != "" {
		address = append(address, a.Country)
	} else if a.CountryCode != "" {
		address = append(address, a.CountryCode)
	}

	if a.PostalCode != "" {
		address = append(address, a.PostalCode)
	}

	return strings.Join(address, ", ")
}
