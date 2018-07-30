// This file was generated for SObject FeedComment, API Version v43.0 at 2018-07-30 03:47:32.987134792 -0400 EDT m=+19.330466562

package sobjects

import (
	"fmt"
	"strings"
)

type FeedComment struct {
	BaseSObject
	CommentBody           string `force:",omitempty"`
	CommentType           string `force:",omitempty"`
	CreatedById           string `force:",omitempty"`
	CreatedDate           string `force:",omitempty"`
	FeedItemId            string `force:",omitempty"`
	HasEntityLinks        bool   `force:",omitempty"`
	Id                    string `force:",omitempty"`
	InsertedById          string `force:",omitempty"`
	IsDeleted             bool   `force:",omitempty"`
	IsRichText            bool   `force:",omitempty"`
	IsVerified            bool   `force:",omitempty"`
	LastEditById          string `force:",omitempty"`
	LastEditDate          string `force:",omitempty"`
	ParentId              string `force:",omitempty"`
	RelatedRecordId       string `force:",omitempty"`
	Revision              int    `force:",omitempty"`
	Status                string `force:",omitempty"`
	SystemModstamp        string `force:",omitempty"`
	ThreadChildrenCount   int    `force:",omitempty"`
	ThreadLastUpdatedDate string `force:",omitempty"`
	ThreadLevel           int    `force:",omitempty"`
	ThreadParentId        string `force:",omitempty"`
}

func (t *FeedComment) ApiName() string {
	return "FeedComment"
}

func (t *FeedComment) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("FeedComment #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCommentBody: %v\n", t.CommentBody))
	builder.WriteString(fmt.Sprintf("\tCommentType: %v\n", t.CommentType))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tFeedItemId: %v\n", t.FeedItemId))
	builder.WriteString(fmt.Sprintf("\tHasEntityLinks: %v\n", t.HasEntityLinks))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tInsertedById: %v\n", t.InsertedById))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tIsRichText: %v\n", t.IsRichText))
	builder.WriteString(fmt.Sprintf("\tIsVerified: %v\n", t.IsVerified))
	builder.WriteString(fmt.Sprintf("\tLastEditById: %v\n", t.LastEditById))
	builder.WriteString(fmt.Sprintf("\tLastEditDate: %v\n", t.LastEditDate))
	builder.WriteString(fmt.Sprintf("\tParentId: %v\n", t.ParentId))
	builder.WriteString(fmt.Sprintf("\tRelatedRecordId: %v\n", t.RelatedRecordId))
	builder.WriteString(fmt.Sprintf("\tRevision: %v\n", t.Revision))
	builder.WriteString(fmt.Sprintf("\tStatus: %v\n", t.Status))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))
	builder.WriteString(fmt.Sprintf("\tThreadChildrenCount: %v\n", t.ThreadChildrenCount))
	builder.WriteString(fmt.Sprintf("\tThreadLastUpdatedDate: %v\n", t.ThreadLastUpdatedDate))
	builder.WriteString(fmt.Sprintf("\tThreadLevel: %v\n", t.ThreadLevel))
	builder.WriteString(fmt.Sprintf("\tThreadParentId: %v\n", t.ThreadParentId))

	return builder.String()
}

type FeedCommentQueryResponse struct {
	BaseQuery
	Records []FeedComment `json:"Records" force:"records"`
}
