// This file was generated for SObject DashboardFeed, API Version v43.0 at 2018-07-30 03:47:18.894209531 -0400 EDT m=+5.237012476

package sobjects

import (
	"fmt"
	"strings"
)

type DashboardFeed struct {
	BaseSObject
	Body             string `force:",omitempty"`
	CommentCount     int    `force:",omitempty"`
	CreatedById      string `force:",omitempty"`
	CreatedDate      string `force:",omitempty"`
	Id               string `force:",omitempty"`
	InsertedById     string `force:",omitempty"`
	IsDeleted        bool   `force:",omitempty"`
	IsRichText       bool   `force:",omitempty"`
	LastModifiedDate string `force:",omitempty"`
	LikeCount        int    `force:",omitempty"`
	LinkUrl          string `force:",omitempty"`
	ParentId         string `force:",omitempty"`
	RelatedRecordId  string `force:",omitempty"`
	SystemModstamp   string `force:",omitempty"`
	Title            string `force:",omitempty"`
	Type             string `force:",omitempty"`
}

func (t *DashboardFeed) ApiName() string {
	return "DashboardFeed"
}

func (t *DashboardFeed) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("DashboardFeed #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tBody: %v\n", t.Body))
	builder.WriteString(fmt.Sprintf("\tCommentCount: %v\n", t.CommentCount))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tInsertedById: %v\n", t.InsertedById))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tIsRichText: %v\n", t.IsRichText))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tLikeCount: %v\n", t.LikeCount))
	builder.WriteString(fmt.Sprintf("\tLinkUrl: %v\n", t.LinkUrl))
	builder.WriteString(fmt.Sprintf("\tParentId: %v\n", t.ParentId))
	builder.WriteString(fmt.Sprintf("\tRelatedRecordId: %v\n", t.RelatedRecordId))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))
	builder.WriteString(fmt.Sprintf("\tTitle: %v\n", t.Title))
	builder.WriteString(fmt.Sprintf("\tType: %v\n", t.Type))

	return builder.String()
}

type DashboardFeedQueryResponse struct {
	BaseQuery
	Records []DashboardFeed `json:"Records" force:"records"`
}
