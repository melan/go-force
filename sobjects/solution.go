// This file was generated for SObject Solution, API Version v43.0 at 2018-07-30 03:47:27.693866269 -0400 EDT m=+14.036999414

package sobjects

import (
	"fmt"
	"strings"
)

type Solution struct {
	BaseSObject
	CreatedById           string `force:",omitempty"`
	CreatedDate           string `force:",omitempty"`
	Id                    string `force:",omitempty"`
	IsDeleted             bool   `force:",omitempty"`
	IsHtml                bool   `force:",omitempty"`
	IsPublished           bool   `force:",omitempty"`
	IsPublishedInPublicKb bool   `force:",omitempty"`
	IsReviewed            bool   `force:",omitempty"`
	LastModifiedById      string `force:",omitempty"`
	LastModifiedDate      string `force:",omitempty"`
	LastReferencedDate    string `force:",omitempty"`
	LastViewedDate        string `force:",omitempty"`
	OwnerId               string `force:",omitempty"`
	SolutionName          string `force:",omitempty"`
	SolutionNote          string `force:",omitempty"`
	SolutionNumber        string `force:",omitempty"`
	Status                string `force:",omitempty"`
	SystemModstamp        string `force:",omitempty"`
	TimesUsed             int    `force:",omitempty"`
}

func (t *Solution) ApiName() string {
	return "Solution"
}

func (t *Solution) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("Solution #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tIsHtml: %v\n", t.IsHtml))
	builder.WriteString(fmt.Sprintf("\tIsPublished: %v\n", t.IsPublished))
	builder.WriteString(fmt.Sprintf("\tIsPublishedInPublicKb: %v\n", t.IsPublishedInPublicKb))
	builder.WriteString(fmt.Sprintf("\tIsReviewed: %v\n", t.IsReviewed))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tLastReferencedDate: %v\n", t.LastReferencedDate))
	builder.WriteString(fmt.Sprintf("\tLastViewedDate: %v\n", t.LastViewedDate))
	builder.WriteString(fmt.Sprintf("\tOwnerId: %v\n", t.OwnerId))
	builder.WriteString(fmt.Sprintf("\tSolutionName: %v\n", t.SolutionName))
	builder.WriteString(fmt.Sprintf("\tSolutionNote: %v\n", t.SolutionNote))
	builder.WriteString(fmt.Sprintf("\tSolutionNumber: %v\n", t.SolutionNumber))
	builder.WriteString(fmt.Sprintf("\tStatus: %v\n", t.Status))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))
	builder.WriteString(fmt.Sprintf("\tTimesUsed: %v\n", t.TimesUsed))

	return builder.String()
}

type SolutionQueryResponse struct {
	BaseQuery
	Records []Solution `json:"Records" force:"records"`
}
