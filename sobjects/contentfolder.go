// This file was generated for SObject ContentFolder, API Version v43.0 at 2018-07-30 03:47:16.803499314 -0400 EDT m=+3.146223807

package sobjects

import (
	"fmt"
	"strings"
)

type ContentFolder struct {
	BaseSObject
	CreatedById           string `force:",omitempty"`
	CreatedDate           string `force:",omitempty"`
	Id                    string `force:",omitempty"`
	IsDeleted             bool   `force:",omitempty"`
	LastModifiedById      string `force:",omitempty"`
	LastModifiedDate      string `force:",omitempty"`
	Name                  string `force:",omitempty"`
	ParentContentFolderId string `force:",omitempty"`
	SystemModstamp        string `force:",omitempty"`
}

func (t *ContentFolder) ApiName() string {
	return "ContentFolder"
}

func (t *ContentFolder) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("ContentFolder #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tName: %v\n", t.Name))
	builder.WriteString(fmt.Sprintf("\tParentContentFolderId: %v\n", t.ParentContentFolderId))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))

	return builder.String()
}

type ContentFolderQueryResponse struct {
	BaseQuery
	Records []ContentFolder `json:"Records" force:"records"`
}
