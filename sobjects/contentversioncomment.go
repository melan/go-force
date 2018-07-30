// This file was generated for SObject ContentVersionComment, API Version v43.0 at 2018-07-30 03:47:50.575198719 -0400 EDT m=+36.919190462

package sobjects

import (
	"fmt"
	"strings"
)

type ContentVersionComment struct {
	BaseSObject
	ContentDocumentId string `force:",omitempty"`
	ContentVersionId  string `force:",omitempty"`
	CreatedDate       string `force:",omitempty"`
	Id                string `force:",omitempty"`
	UserComment       string `force:",omitempty"`
}

func (t *ContentVersionComment) ApiName() string {
	return "ContentVersionComment"
}

func (t *ContentVersionComment) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("ContentVersionComment #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tContentDocumentId: %v\n", t.ContentDocumentId))
	builder.WriteString(fmt.Sprintf("\tContentVersionId: %v\n", t.ContentVersionId))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tUserComment: %v\n", t.UserComment))

	return builder.String()
}

type ContentVersionCommentQueryResponse struct {
	BaseQuery
	Records []ContentVersionComment `json:"Records" force:"records"`
}
