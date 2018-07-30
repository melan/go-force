// This file was generated for SObject CaseTeamTemplateMember, API Version v43.0 at 2018-07-30 03:47:59.69302055 -0400 EDT m=+46.037354430

package sobjects

import (
	"fmt"
	"strings"
)

type CaseTeamTemplateMember struct {
	BaseSObject
	CreatedById      string `force:",omitempty"`
	CreatedDate      string `force:",omitempty"`
	Id               string `force:",omitempty"`
	LastModifiedById string `force:",omitempty"`
	LastModifiedDate string `force:",omitempty"`
	MemberId         string `force:",omitempty"`
	SystemModstamp   string `force:",omitempty"`
	TeamRoleId       string `force:",omitempty"`
	TeamTemplateId   string `force:",omitempty"`
}

func (t *CaseTeamTemplateMember) ApiName() string {
	return "CaseTeamTemplateMember"
}

func (t *CaseTeamTemplateMember) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("CaseTeamTemplateMember #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tMemberId: %v\n", t.MemberId))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))
	builder.WriteString(fmt.Sprintf("\tTeamRoleId: %v\n", t.TeamRoleId))
	builder.WriteString(fmt.Sprintf("\tTeamTemplateId: %v\n", t.TeamTemplateId))

	return builder.String()
}

type CaseTeamTemplateMemberQueryResponse struct {
	BaseQuery
	Records []CaseTeamTemplateMember `json:"Records" force:"records"`
}
