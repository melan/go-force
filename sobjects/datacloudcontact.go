// This file was generated for SObject DatacloudContact, API Version v43.0 at 2018-07-30 03:48:11.604019131 -0400 EDT m=+57.948799959

package sobjects

import (
	"fmt"
	"strings"
)

type DatacloudContact struct {
	BaseSObject
	City        string `force:",omitempty"`
	CompanyId   string `force:",omitempty"`
	CompanyName string `force:",omitempty"`
	ContactId   string `force:",omitempty"`
	Country     string `force:",omitempty"`
	Department  string `force:",omitempty"`
	Email       string `force:",omitempty"`
	ExternalId  string `force:",omitempty"`
	FirstName   string `force:",omitempty"`
	Id          string `force:",omitempty"`
	IsInCrm     bool   `force:",omitempty"`
	IsInactive  bool   `force:",omitempty"`
	IsOwned     bool   `force:",omitempty"`
	LastName    string `force:",omitempty"`
	Level       string `force:",omitempty"`
	Phone       string `force:",omitempty"`
	State       string `force:",omitempty"`
	Street      string `force:",omitempty"`
	Title       string `force:",omitempty"`
	UpdatedDate string `force:",omitempty"`
	Zip         string `force:",omitempty"`
}

func (t *DatacloudContact) ApiName() string {
	return "DatacloudContact"
}

func (t *DatacloudContact) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("DatacloudContact #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCity: %v\n", t.City))
	builder.WriteString(fmt.Sprintf("\tCompanyId: %v\n", t.CompanyId))
	builder.WriteString(fmt.Sprintf("\tCompanyName: %v\n", t.CompanyName))
	builder.WriteString(fmt.Sprintf("\tContactId: %v\n", t.ContactId))
	builder.WriteString(fmt.Sprintf("\tCountry: %v\n", t.Country))
	builder.WriteString(fmt.Sprintf("\tDepartment: %v\n", t.Department))
	builder.WriteString(fmt.Sprintf("\tEmail: %v\n", t.Email))
	builder.WriteString(fmt.Sprintf("\tExternalId: %v\n", t.ExternalId))
	builder.WriteString(fmt.Sprintf("\tFirstName: %v\n", t.FirstName))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsInCrm: %v\n", t.IsInCrm))
	builder.WriteString(fmt.Sprintf("\tIsInactive: %v\n", t.IsInactive))
	builder.WriteString(fmt.Sprintf("\tIsOwned: %v\n", t.IsOwned))
	builder.WriteString(fmt.Sprintf("\tLastName: %v\n", t.LastName))
	builder.WriteString(fmt.Sprintf("\tLevel: %v\n", t.Level))
	builder.WriteString(fmt.Sprintf("\tPhone: %v\n", t.Phone))
	builder.WriteString(fmt.Sprintf("\tState: %v\n", t.State))
	builder.WriteString(fmt.Sprintf("\tStreet: %v\n", t.Street))
	builder.WriteString(fmt.Sprintf("\tTitle: %v\n", t.Title))
	builder.WriteString(fmt.Sprintf("\tUpdatedDate: %v\n", t.UpdatedDate))
	builder.WriteString(fmt.Sprintf("\tZip: %v\n", t.Zip))

	return builder.String()
}

type DatacloudContactQueryResponse struct {
	BaseQuery
	Records []DatacloudContact `json:"Records" force:"records"`
}
