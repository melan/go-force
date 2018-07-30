// This file was generated for SObject CollaborationGroupMemberRequest, API Version v43.0 at 2018-07-30 03:47:37.763690275 -0400 EDT m=+24.107201280

package sobjects

import (
	"fmt"
	"strings"
)

type CollaborationGroupMemberRequest struct {
	BaseSObject
	CollaborationGroupId string `force:",omitempty"`
	CreatedById          string `force:",omitempty"`
	CreatedDate          string `force:",omitempty"`
	Id                   string `force:",omitempty"`
	LastModifiedById     string `force:",omitempty"`
	LastModifiedDate     string `force:",omitempty"`
	RequesterId          string `force:",omitempty"`
	ResponseMessage      string `force:",omitempty"`
	Status               string `force:",omitempty"`
	SystemModstamp       string `force:",omitempty"`
}

func (t *CollaborationGroupMemberRequest) ApiName() string {
	return "CollaborationGroupMemberRequest"
}

func (t *CollaborationGroupMemberRequest) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("CollaborationGroupMemberRequest #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCollaborationGroupId: %v\n", t.CollaborationGroupId))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tRequesterId: %v\n", t.RequesterId))
	builder.WriteString(fmt.Sprintf("\tResponseMessage: %v\n", t.ResponseMessage))
	builder.WriteString(fmt.Sprintf("\tStatus: %v\n", t.Status))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))

	return builder.String()
}

type CollaborationGroupMemberRequestQueryResponse struct {
	BaseQuery
	Records []CollaborationGroupMemberRequest `json:"Records" force:"records"`
}
