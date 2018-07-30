// This file was generated for SObject CustomHttpHeader, API Version v43.0 at 2018-07-30 03:48:08.587294141 -0400 EDT m=+54.931961769

package sobjects

import (
	"fmt"
	"strings"
)

type CustomHttpHeader struct {
	BaseSObject
	CreatedById      string `force:",omitempty"`
	CreatedDate      string `force:",omitempty"`
	Description      string `force:",omitempty"`
	HeaderFieldName  string `force:",omitempty"`
	HeaderFieldValue string `force:",omitempty"`
	Id               string `force:",omitempty"`
	IsActive         bool   `force:",omitempty"`
	LastModifiedById string `force:",omitempty"`
	LastModifiedDate string `force:",omitempty"`
	ParentId         string `force:",omitempty"`
	SystemModstamp   string `force:",omitempty"`
}

func (t *CustomHttpHeader) ApiName() string {
	return "CustomHttpHeader"
}

func (t *CustomHttpHeader) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("CustomHttpHeader #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tDescription: %v\n", t.Description))
	builder.WriteString(fmt.Sprintf("\tHeaderFieldName: %v\n", t.HeaderFieldName))
	builder.WriteString(fmt.Sprintf("\tHeaderFieldValue: %v\n", t.HeaderFieldValue))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsActive: %v\n", t.IsActive))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tParentId: %v\n", t.ParentId))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))

	return builder.String()
}

type CustomHttpHeaderQueryResponse struct {
	BaseQuery
	Records []CustomHttpHeader `json:"Records" force:"records"`
}
