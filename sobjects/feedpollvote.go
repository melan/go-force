// This file was generated for SObject FeedPollVote, API Version v43.0 at 2018-07-30 03:47:41.573961273 -0400 EDT m=+27.917615254

package sobjects

import (
	"fmt"
	"strings"
)

type FeedPollVote struct {
	BaseSObject
	ChoiceId         string `force:",omitempty"`
	CreatedById      string `force:",omitempty"`
	CreatedDate      string `force:",omitempty"`
	FeedItemId       string `force:",omitempty"`
	Id               string `force:",omitempty"`
	IsDeleted        bool   `force:",omitempty"`
	LastModifiedDate string `force:",omitempty"`
}

func (t *FeedPollVote) ApiName() string {
	return "FeedPollVote"
}

func (t *FeedPollVote) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("FeedPollVote #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tChoiceId: %v\n", t.ChoiceId))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tFeedItemId: %v\n", t.FeedItemId))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))

	return builder.String()
}

type FeedPollVoteQueryResponse struct {
	BaseQuery
	Records []FeedPollVote `json:"Records" force:"records"`
}
