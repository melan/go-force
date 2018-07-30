// This file was generated for SObject CustomPermission, API Version v43.0 at 2018-07-30 03:47:41.28762622 -0400 EDT m=+27.631269456

package sobjects

import (
	"fmt"
	"strings"
)

type CustomPermission struct {
	BaseSObject
	CreatedById      string `force:",omitempty"`
	CreatedDate      string `force:",omitempty"`
	Description      string `force:",omitempty"`
	DeveloperName    string `force:",omitempty"`
	Id               string `force:",omitempty"`
	IsDeleted        bool   `force:",omitempty"`
	IsProtected      bool   `force:",omitempty"`
	Language         string `force:",omitempty"`
	LastModifiedById string `force:",omitempty"`
	LastModifiedDate string `force:",omitempty"`
	MasterLabel      string `force:",omitempty"`
	NamespacePrefix  string `force:",omitempty"`
	SystemModstamp   string `force:",omitempty"`
}

func (t *CustomPermission) ApiName() string {
	return "CustomPermission"
}

func (t *CustomPermission) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("CustomPermission #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tDescription: %v\n", t.Description))
	builder.WriteString(fmt.Sprintf("\tDeveloperName: %v\n", t.DeveloperName))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tIsProtected: %v\n", t.IsProtected))
	builder.WriteString(fmt.Sprintf("\tLanguage: %v\n", t.Language))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tMasterLabel: %v\n", t.MasterLabel))
	builder.WriteString(fmt.Sprintf("\tNamespacePrefix: %v\n", t.NamespacePrefix))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))

	return builder.String()
}

type CustomPermissionQueryResponse struct {
	BaseQuery
	Records []CustomPermission `json:"Records" force:"records"`
}
