// This file was generated for SObject FeedLike, API Version v43.0 at 2018-07-30 03:47:18.743501415 -0400 EDT m=+5.086298705

package sobjects

import (
	"fmt"
	"strings"
)

type FeedLike struct {
	BaseSObject
	CreatedById  string `force:",omitempty"`
	CreatedDate  string `force:",omitempty"`
	FeedEntityId string `force:",omitempty"`
	FeedItemId   string `force:",omitempty"`
	Id           string `force:",omitempty"`
	InsertedById string `force:",omitempty"`
	IsDeleted    bool   `force:",omitempty"`
}

func (t *FeedLike) ApiName() string {
	return "FeedLike"
}

func (t *FeedLike) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("FeedLike #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tFeedEntityId: %v\n", t.FeedEntityId))
	builder.WriteString(fmt.Sprintf("\tFeedItemId: %v\n", t.FeedItemId))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tInsertedById: %v\n", t.InsertedById))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))

	return builder.String()
}

type FeedLikeQueryResponse struct {
	BaseQuery
	Records []FeedLike `json:"Records" force:"records"`
}
