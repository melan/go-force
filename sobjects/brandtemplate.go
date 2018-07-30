// This file was generated for SObject BrandTemplate, API Version v43.0 at 2018-07-30 03:48:07.269430022 -0400 EDT m=+53.614048198

package sobjects

import (
	"fmt"
	"strings"
)

type BrandTemplate struct {
	BaseSObject
	CreatedById      string `force:",omitempty"`
	CreatedDate      string `force:",omitempty"`
	Description      string `force:",omitempty"`
	DeveloperName    string `force:",omitempty"`
	Id               string `force:",omitempty"`
	IsActive         bool   `force:",omitempty"`
	LastModifiedById string `force:",omitempty"`
	LastModifiedDate string `force:",omitempty"`
	Name             string `force:",omitempty"`
	NamespacePrefix  string `force:",omitempty"`
	SystemModstamp   string `force:",omitempty"`
	Value            string `force:",omitempty"`
}

func (t *BrandTemplate) ApiName() string {
	return "BrandTemplate"
}

func (t *BrandTemplate) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("BrandTemplate #%s - %s\n", t.Id, t.Name))
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
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))
	builder.WriteString(fmt.Sprintf("\tValue: %v\n", t.Value))

	return builder.String()
}

type BrandTemplateQueryResponse struct {
	BaseQuery
	Records []BrandTemplate `json:"Records" force:"records"`
}
