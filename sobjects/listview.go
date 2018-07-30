// This file was generated for SObject ListView, API Version v43.0 at 2018-07-30 03:47:29.474005346 -0400 EDT m=+15.817205289

package sobjects

import (
	"fmt"
	"strings"
)

type ListView struct {
	BaseSObject
	CreatedById        string `force:",omitempty"`
	CreatedDate        string `force:",omitempty"`
	DeveloperName      string `force:",omitempty"`
	Id                 string `force:",omitempty"`
	IsSoqlCompatible   bool   `force:",omitempty"`
	LastModifiedById   string `force:",omitempty"`
	LastModifiedDate   string `force:",omitempty"`
	LastReferencedDate string `force:",omitempty"`
	LastViewedDate     string `force:",omitempty"`
	Name               string `force:",omitempty"`
	NamespacePrefix    string `force:",omitempty"`
	SobjectType        string `force:",omitempty"`
	SystemModstamp     string `force:",omitempty"`
}

func (t *ListView) ApiName() string {
	return "ListView"
}

func (t *ListView) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("ListView #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tDeveloperName: %v\n", t.DeveloperName))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsSoqlCompatible: %v\n", t.IsSoqlCompatible))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tLastReferencedDate: %v\n", t.LastReferencedDate))
	builder.WriteString(fmt.Sprintf("\tLastViewedDate: %v\n", t.LastViewedDate))
	builder.WriteString(fmt.Sprintf("\tName: %v\n", t.Name))
	builder.WriteString(fmt.Sprintf("\tNamespacePrefix: %v\n", t.NamespacePrefix))
	builder.WriteString(fmt.Sprintf("\tSobjectType: %v\n", t.SobjectType))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))

	return builder.String()
}

type ListViewQueryResponse struct {
	BaseQuery
	Records []ListView `json:"Records" force:"records"`
}
