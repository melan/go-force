// This file was generated for SObject Folder, API Version v43.0 at 2018-07-30 03:47:53.97127505 -0400 EDT m=+40.315394228

package sobjects

import (
	"fmt"
	"strings"
)

type Folder struct {
	BaseSObject
	AccessType       string `force:",omitempty"`
	CreatedById      string `force:",omitempty"`
	CreatedDate      string `force:",omitempty"`
	DeveloperName    string `force:",omitempty"`
	Id               string `force:",omitempty"`
	IsReadonly       bool   `force:",omitempty"`
	LastModifiedById string `force:",omitempty"`
	LastModifiedDate string `force:",omitempty"`
	Name             string `force:",omitempty"`
	NamespacePrefix  string `force:",omitempty"`
	ParentId         string `force:",omitempty"`
	SystemModstamp   string `force:",omitempty"`
	Type             string `force:",omitempty"`
}

func (t *Folder) ApiName() string {
	return "Folder"
}

func (t *Folder) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("Folder #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tAccessType: %v\n", t.AccessType))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tDeveloperName: %v\n", t.DeveloperName))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsReadonly: %v\n", t.IsReadonly))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tName: %v\n", t.Name))
	builder.WriteString(fmt.Sprintf("\tNamespacePrefix: %v\n", t.NamespacePrefix))
	builder.WriteString(fmt.Sprintf("\tParentId: %v\n", t.ParentId))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))
	builder.WriteString(fmt.Sprintf("\tType: %v\n", t.Type))

	return builder.String()
}

type FolderQueryResponse struct {
	BaseQuery
	Records []Folder `json:"Records" force:"records"`
}
