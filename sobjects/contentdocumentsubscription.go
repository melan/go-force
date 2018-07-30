// This file was generated for SObject ContentDocumentSubscription, API Version v43.0 at 2018-07-30 03:47:25.568469276 -0400 EDT m=+11.911522668

package sobjects

import (
	"fmt"
	"strings"
)

type ContentDocumentSubscription struct {
	BaseSObject
	ContentDocumentId string `force:",omitempty"`
	Id                string `force:",omitempty"`
	IsCommentSub      bool   `force:",omitempty"`
	IsDocumentSub     bool   `force:",omitempty"`
	UserId            string `force:",omitempty"`
}

func (t *ContentDocumentSubscription) ApiName() string {
	return "ContentDocumentSubscription"
}

func (t *ContentDocumentSubscription) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("ContentDocumentSubscription #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tContentDocumentId: %v\n", t.ContentDocumentId))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsCommentSub: %v\n", t.IsCommentSub))
	builder.WriteString(fmt.Sprintf("\tIsDocumentSub: %v\n", t.IsDocumentSub))
	builder.WriteString(fmt.Sprintf("\tUserId: %v\n", t.UserId))

	return builder.String()
}

type ContentDocumentSubscriptionQueryResponse struct {
	BaseQuery
	Records []ContentDocumentSubscription `json:"Records" force:"records"`
}
