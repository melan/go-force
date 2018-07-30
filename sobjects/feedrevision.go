// This file was generated for SObject FeedRevision, API Version v43.0 at 2018-07-30 03:47:37.645128955 -0400 EDT m=+23.988635511

package sobjects

import (
	"fmt"
	"strings"
)

type FeedRevision struct {
	BaseSObject
	Action          string `force:",omitempty"`
	CreatedById     string `force:",omitempty"`
	CreatedDate     string `force:",omitempty"`
	EditedAttribute string `force:",omitempty"`
	FeedEntityId    string `force:",omitempty"`
	Id              string `force:",omitempty"`
	IsDeleted       bool   `force:",omitempty"`
	IsValueRichText bool   `force:",omitempty"`
	Revision        int    `force:",omitempty"`
	SystemModstamp  string `force:",omitempty"`
	Value           string `force:",omitempty"`
}

func (t *FeedRevision) ApiName() string {
	return "FeedRevision"
}

func (t *FeedRevision) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("FeedRevision #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tAction: %v\n", t.Action))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tEditedAttribute: %v\n", t.EditedAttribute))
	builder.WriteString(fmt.Sprintf("\tFeedEntityId: %v\n", t.FeedEntityId))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tIsValueRichText: %v\n", t.IsValueRichText))
	builder.WriteString(fmt.Sprintf("\tRevision: %v\n", t.Revision))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))
	builder.WriteString(fmt.Sprintf("\tValue: %v\n", t.Value))

	return builder.String()
}

type FeedRevisionQueryResponse struct {
	BaseQuery
	Records []FeedRevision `json:"Records" force:"records"`
}
