// This file was generated for SObject IdeaComment, API Version v43.0 at 2018-07-30 03:47:52.134088995 -0400 EDT m=+38.478139234

package sobjects

import (
	"fmt"
	"strings"
)

type IdeaComment struct {
	BaseSObject
	CommentBody          string `force:",omitempty"`
	CommunityId          string `force:",omitempty"`
	CreatedById          string `force:",omitempty"`
	CreatedDate          string `force:",omitempty"`
	CreatorFullPhotoUrl  string `force:",omitempty"`
	CreatorName          string `force:",omitempty"`
	CreatorSmallPhotoUrl string `force:",omitempty"`
	Id                   string `force:",omitempty"`
	IdeaId               string `force:",omitempty"`
	IsDeleted            bool   `force:",omitempty"`
	IsHtml               bool   `force:",omitempty"`
	SystemModstamp       string `force:",omitempty"`
	UpVotes              int    `force:",omitempty"`
}

func (t *IdeaComment) ApiName() string {
	return "IdeaComment"
}

func (t *IdeaComment) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("IdeaComment #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCommentBody: %v\n", t.CommentBody))
	builder.WriteString(fmt.Sprintf("\tCommunityId: %v\n", t.CommunityId))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tCreatorFullPhotoUrl: %v\n", t.CreatorFullPhotoUrl))
	builder.WriteString(fmt.Sprintf("\tCreatorName: %v\n", t.CreatorName))
	builder.WriteString(fmt.Sprintf("\tCreatorSmallPhotoUrl: %v\n", t.CreatorSmallPhotoUrl))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIdeaId: %v\n", t.IdeaId))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tIsHtml: %v\n", t.IsHtml))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))
	builder.WriteString(fmt.Sprintf("\tUpVotes: %v\n", t.UpVotes))

	return builder.String()
}

type IdeaCommentQueryResponse struct {
	BaseQuery
	Records []IdeaComment `json:"Records" force:"records"`
}
