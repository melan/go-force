// This file was generated for SObject CaseTeamMember, API Version v43.0 at 2018-07-30 03:47:38.475714471 -0400 EDT m=+24.819252194

package sobjects

import (
	"fmt"
	"strings"
)

type CaseTeamMember struct {
	BaseSObject
	CreatedById          string `force:",omitempty"`
	CreatedDate          string `force:",omitempty"`
	Id                   string `force:",omitempty"`
	LastModifiedById     string `force:",omitempty"`
	LastModifiedDate     string `force:",omitempty"`
	MemberId             string `force:",omitempty"`
	ParentId             string `force:",omitempty"`
	SystemModstamp       string `force:",omitempty"`
	TeamRoleId           string `force:",omitempty"`
	TeamTemplateId       string `force:",omitempty"`
	TeamTemplateMemberId string `force:",omitempty"`
}

func (t *CaseTeamMember) ApiName() string {
	return "CaseTeamMember"
}

func (t *CaseTeamMember) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("CaseTeamMember #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tMemberId: %v\n", t.MemberId))
	builder.WriteString(fmt.Sprintf("\tParentId: %v\n", t.ParentId))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))
	builder.WriteString(fmt.Sprintf("\tTeamRoleId: %v\n", t.TeamRoleId))
	builder.WriteString(fmt.Sprintf("\tTeamTemplateId: %v\n", t.TeamTemplateId))
	builder.WriteString(fmt.Sprintf("\tTeamTemplateMemberId: %v\n", t.TeamTemplateMemberId))

	return builder.String()
}

type CaseTeamMemberQueryResponse struct {
	BaseQuery
	Records []CaseTeamMember `json:"Records" force:"records"`
}
