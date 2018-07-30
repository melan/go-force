// This file was generated for SObject ChatterActivity, API Version v43.0 at 2018-07-30 03:47:26.314334058 -0400 EDT m=+12.657415437

package sobjects

import (
	"fmt"
	"strings"
)

type ChatterActivity struct {
	BaseSObject
	CommentCount         int    `force:",omitempty"`
	CommentReceivedCount int    `force:",omitempty"`
	Id                   string `force:",omitempty"`
	InfluenceRawRank     int    `force:",omitempty"`
	LikeReceivedCount    int    `force:",omitempty"`
	ParentId             string `force:",omitempty"`
	PostCount            int    `force:",omitempty"`
	SystemModstamp       string `force:",omitempty"`
}

func (t *ChatterActivity) ApiName() string {
	return "ChatterActivity"
}

func (t *ChatterActivity) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("ChatterActivity #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCommentCount: %v\n", t.CommentCount))
	builder.WriteString(fmt.Sprintf("\tCommentReceivedCount: %v\n", t.CommentReceivedCount))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tInfluenceRawRank: %v\n", t.InfluenceRawRank))
	builder.WriteString(fmt.Sprintf("\tLikeReceivedCount: %v\n", t.LikeReceivedCount))
	builder.WriteString(fmt.Sprintf("\tParentId: %v\n", t.ParentId))
	builder.WriteString(fmt.Sprintf("\tPostCount: %v\n", t.PostCount))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))

	return builder.String()
}

type ChatterActivityQueryResponse struct {
	BaseQuery
	Records []ChatterActivity `json:"Records" force:"records"`
}
