// This file was generated for SObject GroupMember, API Version v43.0 at 2018-07-30 03:47:29.5954862 -0400 EDT m=+15.938690701

package sobjects

import (
	"fmt"
	"strings"
)

type GroupMember struct {
	BaseSObject
	GroupId        string `force:",omitempty"`
	Id             string `force:",omitempty"`
	SystemModstamp string `force:",omitempty"`
	UserOrGroupId  string `force:",omitempty"`
}

func (t *GroupMember) ApiName() string {
	return "GroupMember"
}

func (t *GroupMember) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("GroupMember #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tGroupId: %v\n", t.GroupId))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))
	builder.WriteString(fmt.Sprintf("\tUserOrGroupId: %v\n", t.UserOrGroupId))

	return builder.String()
}

type GroupMemberQueryResponse struct {
	BaseQuery
	Records []GroupMember `json:"Records" force:"records"`
}
