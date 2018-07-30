// This file was generated for SObject FeedSignal, API Version v43.0 at 2018-07-30 03:47:57.334650169 -0400 EDT m=+43.678895553

package sobjects

import (
	"fmt"
	"strings"
)

type FeedSignal struct {
	BaseSObject
	CreatedById  string `force:",omitempty"`
	CreatedDate  string `force:",omitempty"`
	FeedEntityId string `force:",omitempty"`
	FeedItemId   string `force:",omitempty"`
	Id           string `force:",omitempty"`
	InsertedById string `force:",omitempty"`
	IsDeleted    bool   `force:",omitempty"`
	SignalType   string `force:",omitempty"`
	SignalValue  int    `force:",omitempty"`
}

func (t *FeedSignal) ApiName() string {
	return "FeedSignal"
}

func (t *FeedSignal) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("FeedSignal #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tFeedEntityId: %v\n", t.FeedEntityId))
	builder.WriteString(fmt.Sprintf("\tFeedItemId: %v\n", t.FeedItemId))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tInsertedById: %v\n", t.InsertedById))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tSignalType: %v\n", t.SignalType))
	builder.WriteString(fmt.Sprintf("\tSignalValue: %v\n", t.SignalValue))

	return builder.String()
}

type FeedSignalQueryResponse struct {
	BaseQuery
	Records []FeedSignal `json:"Records" force:"records"`
}
