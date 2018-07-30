// This file was generated for SObject SessionPermSetActivation, API Version v43.0 at 2018-07-30 03:47:52.258200777 -0400 EDT m=+38.602255673

package sobjects

import (
	"fmt"
	"strings"
)

type SessionPermSetActivation struct {
	BaseSObject
	AuthSessionId    string `force:",omitempty"`
	CreatedById      string `force:",omitempty"`
	CreatedDate      string `force:",omitempty"`
	Description      string `force:",omitempty"`
	Id               string `force:",omitempty"`
	IsDeleted        bool   `force:",omitempty"`
	LastModifiedById string `force:",omitempty"`
	LastModifiedDate string `force:",omitempty"`
	PermissionSetId  string `force:",omitempty"`
	SystemModstamp   string `force:",omitempty"`
	UserId           string `force:",omitempty"`
}

func (t *SessionPermSetActivation) ApiName() string {
	return "SessionPermSetActivation"
}

func (t *SessionPermSetActivation) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("SessionPermSetActivation #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tAuthSessionId: %v\n", t.AuthSessionId))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tDescription: %v\n", t.Description))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tPermissionSetId: %v\n", t.PermissionSetId))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))
	builder.WriteString(fmt.Sprintf("\tUserId: %v\n", t.UserId))

	return builder.String()
}

type SessionPermSetActivationQueryResponse struct {
	BaseQuery
	Records []SessionPermSetActivation `json:"Records" force:"records"`
}
