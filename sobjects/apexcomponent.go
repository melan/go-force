// This file was generated for SObject ApexComponent, API Version v43.0 at 2018-07-30 03:47:57.703399514 -0400 EDT m=+44.047658735

package sobjects

import (
	"fmt"
	"strings"
)

type ApexComponent struct {
	BaseSObject
	ApiVersion       float64 `force:",omitempty"`
	ControllerKey    string  `force:",omitempty"`
	ControllerType   string  `force:",omitempty"`
	CreatedById      string  `force:",omitempty"`
	CreatedDate      string  `force:",omitempty"`
	Description      string  `force:",omitempty"`
	Id               string  `force:",omitempty"`
	LastModifiedById string  `force:",omitempty"`
	LastModifiedDate string  `force:",omitempty"`
	Markup           string  `force:",omitempty"`
	MasterLabel      string  `force:",omitempty"`
	Name             string  `force:",omitempty"`
	NamespacePrefix  string  `force:",omitempty"`
	SystemModstamp   string  `force:",omitempty"`
}

func (t *ApexComponent) ApiName() string {
	return "ApexComponent"
}

func (t *ApexComponent) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("ApexComponent #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tApiVersion: %v\n", t.ApiVersion))
	builder.WriteString(fmt.Sprintf("\tControllerKey: %v\n", t.ControllerKey))
	builder.WriteString(fmt.Sprintf("\tControllerType: %v\n", t.ControllerType))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tDescription: %v\n", t.Description))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tMarkup: %v\n", t.Markup))
	builder.WriteString(fmt.Sprintf("\tMasterLabel: %v\n", t.MasterLabel))
	builder.WriteString(fmt.Sprintf("\tName: %v\n", t.Name))
	builder.WriteString(fmt.Sprintf("\tNamespacePrefix: %v\n", t.NamespacePrefix))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))

	return builder.String()
}

type ApexComponentQueryResponse struct {
	BaseQuery
	Records []ApexComponent `json:"Records" force:"records"`
}
