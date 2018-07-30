// This file was generated for SObject CaseComment, API Version v43.0 at 2018-07-30 03:47:15.882139002 -0400 EDT m=+2.224828922

package sobjects

import (
	"fmt"
	"strings"
)

type CaseComment struct {
	BaseSObject
	CommentBody      string `force:",omitempty"`
	CreatedById      string `force:",omitempty"`
	CreatedDate      string `force:",omitempty"`
	Id               string `force:",omitempty"`
	IsDeleted        bool   `force:",omitempty"`
	IsPublished      bool   `force:",omitempty"`
	LastModifiedById string `force:",omitempty"`
	LastModifiedDate string `force:",omitempty"`
	ParentId         string `force:",omitempty"`
	SystemModstamp   string `force:",omitempty"`
}

func (t *CaseComment) ApiName() string {
	return "CaseComment"
}

func (t *CaseComment) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("CaseComment #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCommentBody: %v\n", t.CommentBody))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tIsPublished: %v\n", t.IsPublished))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tParentId: %v\n", t.ParentId))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))

	return builder.String()
}

type CaseCommentQueryResponse struct {
	BaseQuery
	Records []CaseComment `json:"Records" force:"records"`
}
