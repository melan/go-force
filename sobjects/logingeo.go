// This file was generated for SObject LoginGeo, API Version v43.0 at 2018-07-30 03:47:32.50817275 -0400 EDT m=+18.851486547

package sobjects

import (
	"fmt"
	"strings"
)

type LoginGeo struct {
	BaseSObject
	City             string  `force:",omitempty"`
	Country          string  `force:",omitempty"`
	CountryIso       string  `force:",omitempty"`
	CreatedById      string  `force:",omitempty"`
	CreatedDate      string  `force:",omitempty"`
	Id               string  `force:",omitempty"`
	IsDeleted        bool    `force:",omitempty"`
	LastModifiedById string  `force:",omitempty"`
	LastModifiedDate string  `force:",omitempty"`
	Latitude         float64 `force:",omitempty"`
	LoginTime        string  `force:",omitempty"`
	Longitude        float64 `force:",omitempty"`
	PostalCode       string  `force:",omitempty"`
	Subdivision      string  `force:",omitempty"`
	SystemModstamp   string  `force:",omitempty"`
}

func (t *LoginGeo) ApiName() string {
	return "LoginGeo"
}

func (t *LoginGeo) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("LoginGeo #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCity: %v\n", t.City))
	builder.WriteString(fmt.Sprintf("\tCountry: %v\n", t.Country))
	builder.WriteString(fmt.Sprintf("\tCountryIso: %v\n", t.CountryIso))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tLatitude: %v\n", t.Latitude))
	builder.WriteString(fmt.Sprintf("\tLoginTime: %v\n", t.LoginTime))
	builder.WriteString(fmt.Sprintf("\tLongitude: %v\n", t.Longitude))
	builder.WriteString(fmt.Sprintf("\tPostalCode: %v\n", t.PostalCode))
	builder.WriteString(fmt.Sprintf("\tSubdivision: %v\n", t.Subdivision))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))

	return builder.String()
}

type LoginGeoQueryResponse struct {
	BaseQuery
	Records []LoginGeo `json:"Records" force:"records"`
}
