// This file was generated for SObject ContentDocumentLink, API Version v43.0 at 2018-07-30 03:47:56.377179234 -0400 EDT m=+42.721388690

package sobjects

import (
	"fmt"
	"strings"
)

type ContentDocumentLink struct {
	BaseSObject
	ContentDocumentId string `force:",omitempty"`
	Id                string `force:",omitempty"`
	IsDeleted         bool   `force:",omitempty"`
	LinkedEntityId    string `force:",omitempty"`
	ShareType         string `force:",omitempty"`
	SystemModstamp    string `force:",omitempty"`
	Visibility        string `force:",omitempty"`
}

func (t *ContentDocumentLink) ApiName() string {
	return "ContentDocumentLink"
}

func (t *ContentDocumentLink) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("ContentDocumentLink #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tContentDocumentId: %v\n", t.ContentDocumentId))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tLinkedEntityId: %v\n", t.LinkedEntityId))
	builder.WriteString(fmt.Sprintf("\tShareType: %v\n", t.ShareType))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))
	builder.WriteString(fmt.Sprintf("\tVisibility: %v\n", t.Visibility))

	return builder.String()
}

type ContentDocumentLinkQueryResponse struct {
	BaseQuery
	Records []ContentDocumentLink `json:"Records" force:"records"`
}
