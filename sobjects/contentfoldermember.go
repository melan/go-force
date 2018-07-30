// This file was generated for SObject ContentFolderMember, API Version v43.0 at 2018-07-30 03:48:08.707127664 -0400 EDT m=+55.051799789

package sobjects

import (
	"fmt"
	"strings"
)

type ContentFolderMember struct {
	BaseSObject
	ChildRecordId         string `force:",omitempty"`
	CreatedById           string `force:",omitempty"`
	CreatedDate           string `force:",omitempty"`
	Id                    string `force:",omitempty"`
	IsDeleted             bool   `force:",omitempty"`
	LastModifiedById      string `force:",omitempty"`
	LastModifiedDate      string `force:",omitempty"`
	ParentContentFolderId string `force:",omitempty"`
	SystemModstamp        string `force:",omitempty"`
}

func (t *ContentFolderMember) ApiName() string {
	return "ContentFolderMember"
}

func (t *ContentFolderMember) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("ContentFolderMember #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tChildRecordId: %v\n", t.ChildRecordId))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tParentContentFolderId: %v\n", t.ParentContentFolderId))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))

	return builder.String()
}

type ContentFolderMemberQueryResponse struct {
	BaseQuery
	Records []ContentFolderMember `json:"Records" force:"records"`
}
