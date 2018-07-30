// This file was generated for SObject FeedPollChoice, API Version v43.0 at 2018-07-30 03:47:30.655707301 -0400 EDT m=+16.998951586

package sobjects

import (
	"fmt"
	"strings"
)

type FeedPollChoice struct {
	BaseSObject
	ChoiceBody  string `force:",omitempty"`
	CreatedById string `force:",omitempty"`
	CreatedDate string `force:",omitempty"`
	FeedItemId  string `force:",omitempty"`
	Id          string `force:",omitempty"`
	IsDeleted   bool   `force:",omitempty"`
	Position    int    `force:",omitempty"`
}

func (t *FeedPollChoice) ApiName() string {
	return "FeedPollChoice"
}

func (t *FeedPollChoice) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("FeedPollChoice #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tChoiceBody: %v\n", t.ChoiceBody))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tFeedItemId: %v\n", t.FeedItemId))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tPosition: %v\n", t.Position))

	return builder.String()
}

type FeedPollChoiceQueryResponse struct {
	BaseQuery
	Records []FeedPollChoice `json:"Records" force:"records"`
}
