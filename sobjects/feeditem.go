// This file was generated for SObject FeedItem, API Version v43.0 at 2018-07-30 03:47:41.434037839 -0400 EDT m=+27.777686570

package sobjects

import (
	"fmt"
	"strings"
)

type FeedItem struct {
	BaseSObject
	BestCommentId      string `force:",omitempty"`
	Body               string `force:",omitempty"`
	CommentCount       int    `force:",omitempty"`
	CreatedById        string `force:",omitempty"`
	CreatedDate        string `force:",omitempty"`
	HasContent         bool   `force:",omitempty"`
	HasFeedEntity      bool   `force:",omitempty"`
	HasLink            bool   `force:",omitempty"`
	HasVerifiedComment bool   `force:",omitempty"`
	Id                 string `force:",omitempty"`
	InsertedById       string `force:",omitempty"`
	IsClosed           bool   `force:",omitempty"`
	IsDeleted          bool   `force:",omitempty"`
	IsRichText         bool   `force:",omitempty"`
	LastEditById       string `force:",omitempty"`
	LastEditDate       string `force:",omitempty"`
	LastModifiedDate   string `force:",omitempty"`
	LikeCount          int    `force:",omitempty"`
	LinkUrl            string `force:",omitempty"`
	ParentId           string `force:",omitempty"`
	RelatedRecordId    string `force:",omitempty"`
	Revision           int    `force:",omitempty"`
	Status             string `force:",omitempty"`
	SystemModstamp     string `force:",omitempty"`
	Title              string `force:",omitempty"`
	Type               string `force:",omitempty"`
}

func (t *FeedItem) ApiName() string {
	return "FeedItem"
}

func (t *FeedItem) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("FeedItem #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tBestCommentId: %v\n", t.BestCommentId))
	builder.WriteString(fmt.Sprintf("\tBody: %v\n", t.Body))
	builder.WriteString(fmt.Sprintf("\tCommentCount: %v\n", t.CommentCount))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tHasContent: %v\n", t.HasContent))
	builder.WriteString(fmt.Sprintf("\tHasFeedEntity: %v\n", t.HasFeedEntity))
	builder.WriteString(fmt.Sprintf("\tHasLink: %v\n", t.HasLink))
	builder.WriteString(fmt.Sprintf("\tHasVerifiedComment: %v\n", t.HasVerifiedComment))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tInsertedById: %v\n", t.InsertedById))
	builder.WriteString(fmt.Sprintf("\tIsClosed: %v\n", t.IsClosed))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tIsRichText: %v\n", t.IsRichText))
	builder.WriteString(fmt.Sprintf("\tLastEditById: %v\n", t.LastEditById))
	builder.WriteString(fmt.Sprintf("\tLastEditDate: %v\n", t.LastEditDate))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tLikeCount: %v\n", t.LikeCount))
	builder.WriteString(fmt.Sprintf("\tLinkUrl: %v\n", t.LinkUrl))
	builder.WriteString(fmt.Sprintf("\tParentId: %v\n", t.ParentId))
	builder.WriteString(fmt.Sprintf("\tRelatedRecordId: %v\n", t.RelatedRecordId))
	builder.WriteString(fmt.Sprintf("\tRevision: %v\n", t.Revision))
	builder.WriteString(fmt.Sprintf("\tStatus: %v\n", t.Status))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))
	builder.WriteString(fmt.Sprintf("\tTitle: %v\n", t.Title))
	builder.WriteString(fmt.Sprintf("\tType: %v\n", t.Type))

	return builder.String()
}

type FeedItemQueryResponse struct {
	BaseQuery
	Records []FeedItem `json:"Records" force:"records"`
}
