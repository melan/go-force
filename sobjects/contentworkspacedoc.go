// This file was generated for SObject ContentWorkspaceDoc, API Version v43.0 at 2018-07-30 03:47:53.704672384 -0400 EDT m=+40.048781557

package sobjects

import (
	"fmt"
	"strings"
)

type ContentWorkspaceDoc struct {
	BaseSObject
	ContentDocumentId  string `force:",omitempty"`
	ContentWorkspaceId string `force:",omitempty"`
	CreatedDate        string `force:",omitempty"`
	Id                 string `force:",omitempty"`
	IsDeleted          bool   `force:",omitempty"`
	IsOwner            bool   `force:",omitempty"`
	SystemModstamp     string `force:",omitempty"`
}

func (t *ContentWorkspaceDoc) ApiName() string {
	return "ContentWorkspaceDoc"
}

func (t *ContentWorkspaceDoc) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("ContentWorkspaceDoc #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tContentDocumentId: %v\n", t.ContentDocumentId))
	builder.WriteString(fmt.Sprintf("\tContentWorkspaceId: %v\n", t.ContentWorkspaceId))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tIsOwner: %v\n", t.IsOwner))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))

	return builder.String()
}

type ContentWorkspaceDocQueryResponse struct {
	BaseQuery
	Records []ContentWorkspaceDoc `json:"Records" force:"records"`
}
