// This file was generated for SObject DatacloudAddress, API Version v43.0 at 2018-07-30 03:48:06.878234807 -0400 EDT m=+53.222838305

package sobjects

import (
	"fmt"
	"strings"
)

type DatacloudAddress struct {
	BaseSObject
	AddressLine1    string `force:",omitempty"`
	AddressLine2    string `force:",omitempty"`
	City            string `force:",omitempty"`
	Country         string `force:",omitempty"`
	ExternalId      string `force:",omitempty"`
	GeoAccuracyCode string `force:",omitempty"`
	GeoAccuracyNum  string `force:",omitempty"`
	Id              string `force:",omitempty"`
	Latitude        string `force:",omitempty"`
	Longitude       string `force:",omitempty"`
	PostalCode      string `force:",omitempty"`
	State           string `force:",omitempty"`
}

func (t *DatacloudAddress) ApiName() string {
	return "DatacloudAddress"
}

func (t *DatacloudAddress) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("DatacloudAddress #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tAddressLine1: %v\n", t.AddressLine1))
	builder.WriteString(fmt.Sprintf("\tAddressLine2: %v\n", t.AddressLine2))
	builder.WriteString(fmt.Sprintf("\tCity: %v\n", t.City))
	builder.WriteString(fmt.Sprintf("\tCountry: %v\n", t.Country))
	builder.WriteString(fmt.Sprintf("\tExternalId: %v\n", t.ExternalId))
	builder.WriteString(fmt.Sprintf("\tGeoAccuracyCode: %v\n", t.GeoAccuracyCode))
	builder.WriteString(fmt.Sprintf("\tGeoAccuracyNum: %v\n", t.GeoAccuracyNum))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tLatitude: %v\n", t.Latitude))
	builder.WriteString(fmt.Sprintf("\tLongitude: %v\n", t.Longitude))
	builder.WriteString(fmt.Sprintf("\tPostalCode: %v\n", t.PostalCode))
	builder.WriteString(fmt.Sprintf("\tState: %v\n", t.State))

	return builder.String()
}

type DatacloudAddressQueryResponse struct {
	BaseQuery
	Records []DatacloudAddress `json:"Records" force:"records"`
}
