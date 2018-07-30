// This file was generated for SObject ContentWorkspaceMember, API Version v43.0 at 2018-07-30 03:47:23.111111129 -0400 EDT m=+9.454072311

package sobjects

import (
	"fmt"
	"strings"
)

type ContentWorkspaceMember struct {
	BaseSObject
	ContentWorkspaceId           string `force:",omitempty"`
	ContentWorkspacePermissionId string `force:",omitempty"`
	CreatedById                  string `force:",omitempty"`
	CreatedDate                  string `force:",omitempty"`
	Id                           string `force:",omitempty"`
	MemberId                     string `force:",omitempty"`
	MemberType                   string `force:",omitempty"`
}

func (t *ContentWorkspaceMember) ApiName() string {
	return "ContentWorkspaceMember"
}

func (t *ContentWorkspaceMember) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("ContentWorkspaceMember #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tContentWorkspaceId: %v\n", t.ContentWorkspaceId))
	builder.WriteString(fmt.Sprintf("\tContentWorkspacePermissionId: %v\n", t.ContentWorkspacePermissionId))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tMemberId: %v\n", t.MemberId))
	builder.WriteString(fmt.Sprintf("\tMemberType: %v\n", t.MemberType))

	return builder.String()
}

type ContentWorkspaceMemberQueryResponse struct {
	BaseQuery
	Records []ContentWorkspaceMember `json:"Records" force:"records"`
}
