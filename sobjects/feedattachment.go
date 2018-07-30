// This file was generated for SObject FeedAttachment, API Version v43.0 at 2018-07-30 03:48:11.864160076 -0400 EDT m=+58.208950665

package sobjects

import (
	"fmt"
	"strings"
)

type FeedAttachment struct {
	BaseSObject
	FeedEntityId string `force:",omitempty"`
	Id           string `force:",omitempty"`
	IsDeleted    bool   `force:",omitempty"`
	RecordId     string `force:",omitempty"`
	Title        string `force:",omitempty"`
	Type         string `force:",omitempty"`
	Value        string `force:",omitempty"`
}

func (t *FeedAttachment) ApiName() string {
	return "FeedAttachment"
}

func (t *FeedAttachment) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("FeedAttachment #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tFeedEntityId: %v\n", t.FeedEntityId))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tRecordId: %v\n", t.RecordId))
	builder.WriteString(fmt.Sprintf("\tTitle: %v\n", t.Title))
	builder.WriteString(fmt.Sprintf("\tType: %v\n", t.Type))
	builder.WriteString(fmt.Sprintf("\tValue: %v\n", t.Value))

	return builder.String()
}

type FeedAttachmentQueryResponse struct {
	BaseQuery
	Records []FeedAttachment `json:"Records" force:"records"`
}
