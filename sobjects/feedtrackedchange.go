// This file was generated for SObject FeedTrackedChange, API Version v43.0 at 2018-07-30 03:47:33.895560215 -0400 EDT m=+20.238926072

package sobjects

import (
	"fmt"
	"strings"
)

type FeedTrackedChange struct {
	BaseSObject
	FeedItemId string `force:",omitempty"`
	FieldName  string `force:",omitempty"`
	Id         string `force:",omitempty"`
	NewValue   string `force:",omitempty"`
	OldValue   string `force:",omitempty"`
}

func (t *FeedTrackedChange) ApiName() string {
	return "FeedTrackedChange"
}

func (t *FeedTrackedChange) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("FeedTrackedChange #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tFeedItemId: %v\n", t.FeedItemId))
	builder.WriteString(fmt.Sprintf("\tFieldName: %v\n", t.FieldName))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tNewValue: %v\n", t.NewValue))
	builder.WriteString(fmt.Sprintf("\tOldValue: %v\n", t.OldValue))

	return builder.String()
}

type FeedTrackedChangeQueryResponse struct {
	BaseQuery
	Records []FeedTrackedChange `json:"Records" force:"records"`
}
