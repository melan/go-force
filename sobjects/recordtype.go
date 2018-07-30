// This file was generated for SObject RecordType, API Version v43.0 at 2018-07-30 03:47:16.646162446 -0400 EDT m=+2.988881035

package sobjects

import (
	"fmt"
	"strings"
)

type RecordType struct {
	BaseSObject
	BusinessProcessId string `force:",omitempty"`
	CreatedById       string `force:",omitempty"`
	CreatedDate       string `force:",omitempty"`
	Description       string `force:",omitempty"`
	DeveloperName     string `force:",omitempty"`
	Id                string `force:",omitempty"`
	IsActive          bool   `force:",omitempty"`
	LastModifiedById  string `force:",omitempty"`
	LastModifiedDate  string `force:",omitempty"`
	Name              string `force:",omitempty"`
	NamespacePrefix   string `force:",omitempty"`
	SobjectType       string `force:",omitempty"`
	SystemModstamp    string `force:",omitempty"`
}

func (t *RecordType) ApiName() string {
	return "RecordType"
}

func (t *RecordType) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("RecordType #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tBusinessProcessId: %v\n", t.BusinessProcessId))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tDescription: %v\n", t.Description))
	builder.WriteString(fmt.Sprintf("\tDeveloperName: %v\n", t.DeveloperName))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsActive: %v\n", t.IsActive))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tName: %v\n", t.Name))
	builder.WriteString(fmt.Sprintf("\tNamespacePrefix: %v\n", t.NamespacePrefix))
	builder.WriteString(fmt.Sprintf("\tSobjectType: %v\n", t.SobjectType))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))

	return builder.String()
}

type RecordTypeQueryResponse struct {
	BaseQuery
	Records []RecordType `json:"Records" force:"records"`
}
