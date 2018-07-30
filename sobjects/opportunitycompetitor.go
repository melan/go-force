// This file was generated for SObject OpportunityCompetitor, API Version v43.0 at 2018-07-30 03:47:47.5274252 -0400 EDT m=+33.871302579

package sobjects

import (
	"fmt"
	"strings"
)

type OpportunityCompetitor struct {
	BaseSObject
	CompetitorName   string `force:",omitempty"`
	CreatedById      string `force:",omitempty"`
	CreatedDate      string `force:",omitempty"`
	Id               string `force:",omitempty"`
	IsDeleted        bool   `force:",omitempty"`
	LastModifiedById string `force:",omitempty"`
	LastModifiedDate string `force:",omitempty"`
	OpportunityId    string `force:",omitempty"`
	Strengths        string `force:",omitempty"`
	SystemModstamp   string `force:",omitempty"`
	Weaknesses       string `force:",omitempty"`
}

func (t *OpportunityCompetitor) ApiName() string {
	return "OpportunityCompetitor"
}

func (t *OpportunityCompetitor) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("OpportunityCompetitor #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCompetitorName: %v\n", t.CompetitorName))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tOpportunityId: %v\n", t.OpportunityId))
	builder.WriteString(fmt.Sprintf("\tStrengths: %v\n", t.Strengths))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))
	builder.WriteString(fmt.Sprintf("\tWeaknesses: %v\n", t.Weaknesses))

	return builder.String()
}

type OpportunityCompetitorQueryResponse struct {
	BaseQuery
	Records []OpportunityCompetitor `json:"Records" force:"records"`
}
