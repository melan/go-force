// This file was generated for SObject ContentFolderLink, API Version v43.0 at 2018-07-30 03:47:35.98721158 -0400 EDT m=+22.330655924

package sobjects

import (
	"fmt"
	"strings"
)

type ContentFolderLink struct {
	BaseSObject
	ContentFolderId    string `force:",omitempty"`
	EnableFolderStatus string `force:",omitempty"`
	Id                 string `force:",omitempty"`
	IsDeleted          bool   `force:",omitempty"`
	ParentEntityId     string `force:",omitempty"`
}

func (t *ContentFolderLink) ApiName() string {
	return "ContentFolderLink"
}

func (t *ContentFolderLink) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("ContentFolderLink #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tContentFolderId: %v\n", t.ContentFolderId))
	builder.WriteString(fmt.Sprintf("\tEnableFolderStatus: %v\n", t.EnableFolderStatus))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tParentEntityId: %v\n", t.ParentEntityId))

	return builder.String()
}

type ContentFolderLinkQueryResponse struct {
	BaseQuery
	Records []ContentFolderLink `json:"Records" force:"records"`
}
