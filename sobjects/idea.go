// This file was generated for SObject Idea, API Version v43.0 at 2018-07-30 03:47:20.482032335 -0400 EDT m=+6.824894863

package sobjects

import (
	"fmt"
	"strings"
)

type Idea struct {
	BaseSObject
	Body                 string  `force:",omitempty"`
	Categories           string  `force:",omitempty"`
	CommunityId          string  `force:",omitempty"`
	CreatedById          string  `force:",omitempty"`
	CreatedDate          string  `force:",omitempty"`
	CreatorFullPhotoUrl  string  `force:",omitempty"`
	CreatorName          string  `force:",omitempty"`
	CreatorSmallPhotoUrl string  `force:",omitempty"`
	Id                   string  `force:",omitempty"`
	IsDeleted            bool    `force:",omitempty"`
	IsHtml               bool    `force:",omitempty"`
	IsMerged             bool    `force:",omitempty"`
	LastCommentDate      string  `force:",omitempty"`
	LastCommentId        string  `force:",omitempty"`
	LastModifiedById     string  `force:",omitempty"`
	LastModifiedDate     string  `force:",omitempty"`
	LastReferencedDate   string  `force:",omitempty"`
	LastViewedDate       string  `force:",omitempty"`
	NumComments          int     `force:",omitempty"`
	ParentIdeaId         string  `force:",omitempty"`
	RecordTypeId         string  `force:",omitempty"`
	Status               string  `force:",omitempty"`
	SystemModstamp       string  `force:",omitempty"`
	Title                string  `force:",omitempty"`
	VoteScore            float64 `force:",omitempty"`
	VoteTotal            float64 `force:",omitempty"`
}

func (t *Idea) ApiName() string {
	return "Idea"
}

func (t *Idea) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("Idea #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tBody: %v\n", t.Body))
	builder.WriteString(fmt.Sprintf("\tCategories: %v\n", t.Categories))
	builder.WriteString(fmt.Sprintf("\tCommunityId: %v\n", t.CommunityId))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tCreatorFullPhotoUrl: %v\n", t.CreatorFullPhotoUrl))
	builder.WriteString(fmt.Sprintf("\tCreatorName: %v\n", t.CreatorName))
	builder.WriteString(fmt.Sprintf("\tCreatorSmallPhotoUrl: %v\n", t.CreatorSmallPhotoUrl))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tIsHtml: %v\n", t.IsHtml))
	builder.WriteString(fmt.Sprintf("\tIsMerged: %v\n", t.IsMerged))
	builder.WriteString(fmt.Sprintf("\tLastCommentDate: %v\n", t.LastCommentDate))
	builder.WriteString(fmt.Sprintf("\tLastCommentId: %v\n", t.LastCommentId))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tLastReferencedDate: %v\n", t.LastReferencedDate))
	builder.WriteString(fmt.Sprintf("\tLastViewedDate: %v\n", t.LastViewedDate))
	builder.WriteString(fmt.Sprintf("\tNumComments: %v\n", t.NumComments))
	builder.WriteString(fmt.Sprintf("\tParentIdeaId: %v\n", t.ParentIdeaId))
	builder.WriteString(fmt.Sprintf("\tRecordTypeId: %v\n", t.RecordTypeId))
	builder.WriteString(fmt.Sprintf("\tStatus: %v\n", t.Status))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))
	builder.WriteString(fmt.Sprintf("\tTitle: %v\n", t.Title))
	builder.WriteString(fmt.Sprintf("\tVoteScore: %v\n", t.VoteScore))
	builder.WriteString(fmt.Sprintf("\tVoteTotal: %v\n", t.VoteTotal))

	return builder.String()
}

type IdeaQueryResponse struct {
	BaseQuery
	Records []Idea `json:"Records" force:"records"`
}
