// This file was generated for SObject CollaborationInvitation, API Version v43.0 at 2018-07-30 03:48:00.115930459 -0400 EDT m=+46.460280208

package sobjects

import (
	"fmt"
	"strings"
)

type CollaborationInvitation struct {
	BaseSObject
	CreatedById                string `force:",omitempty"`
	CreatedDate                string `force:",omitempty"`
	Id                         string `force:",omitempty"`
	InvitedUserEmail           string `force:",omitempty"`
	InvitedUserEmailNormalized string `force:",omitempty"`
	InviterId                  string `force:",omitempty"`
	LastModifiedById           string `force:",omitempty"`
	LastModifiedDate           string `force:",omitempty"`
	OptionalMessage            string `force:",omitempty"`
	ParentId                   string `force:",omitempty"`
	SharedEntityId             string `force:",omitempty"`
	Status                     string `force:",omitempty"`
	SystemModstamp             string `force:",omitempty"`
}

func (t *CollaborationInvitation) ApiName() string {
	return "CollaborationInvitation"
}

func (t *CollaborationInvitation) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("CollaborationInvitation #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tInvitedUserEmail: %v\n", t.InvitedUserEmail))
	builder.WriteString(fmt.Sprintf("\tInvitedUserEmailNormalized: %v\n", t.InvitedUserEmailNormalized))
	builder.WriteString(fmt.Sprintf("\tInviterId: %v\n", t.InviterId))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tOptionalMessage: %v\n", t.OptionalMessage))
	builder.WriteString(fmt.Sprintf("\tParentId: %v\n", t.ParentId))
	builder.WriteString(fmt.Sprintf("\tSharedEntityId: %v\n", t.SharedEntityId))
	builder.WriteString(fmt.Sprintf("\tStatus: %v\n", t.Status))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))

	return builder.String()
}

type CollaborationInvitationQueryResponse struct {
	BaseQuery
	Records []CollaborationInvitation `json:"Records" force:"records"`
}
