// This file was generated for SObject ApexClass, API Version v43.0 at 2018-07-30 03:47:23.946271246 -0400 EDT m=+10.289263767

package sobjects

import (
	"fmt"
	"strings"
)

type ApexClass struct {
	BaseSObject
	ApiVersion            float64 `force:",omitempty"`
	Body                  string  `force:",omitempty"`
	BodyCrc               float64 `force:",omitempty"`
	CreatedById           string  `force:",omitempty"`
	CreatedDate           string  `force:",omitempty"`
	Id                    string  `force:",omitempty"`
	IsValid               bool    `force:",omitempty"`
	LastModifiedById      string  `force:",omitempty"`
	LastModifiedDate      string  `force:",omitempty"`
	LengthWithoutComments int     `force:",omitempty"`
	Name                  string  `force:",omitempty"`
	NamespacePrefix       string  `force:",omitempty"`
	Status                string  `force:",omitempty"`
	SystemModstamp        string  `force:",omitempty"`
}

func (t *ApexClass) ApiName() string {
	return "ApexClass"
}

func (t *ApexClass) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("ApexClass #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tApiVersion: %v\n", t.ApiVersion))
	builder.WriteString(fmt.Sprintf("\tBody: %v\n", t.Body))
	builder.WriteString(fmt.Sprintf("\tBodyCrc: %v\n", t.BodyCrc))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsValid: %v\n", t.IsValid))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tLengthWithoutComments: %v\n", t.LengthWithoutComments))
	builder.WriteString(fmt.Sprintf("\tName: %v\n", t.Name))
	builder.WriteString(fmt.Sprintf("\tNamespacePrefix: %v\n", t.NamespacePrefix))
	builder.WriteString(fmt.Sprintf("\tStatus: %v\n", t.Status))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))

	return builder.String()
}

type ApexClassQueryResponse struct {
	BaseQuery
	Records []ApexClass `json:"Records" force:"records"`
}
