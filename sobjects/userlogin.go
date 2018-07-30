// This file was generated for SObject UserLogin, API Version v43.0 at 2018-07-30 03:48:05.949060868 -0400 EDT m=+52.293629499

package sobjects

import (
	"fmt"
	"strings"
)

type UserLogin struct {
	BaseSObject
	Id               string `force:",omitempty"`
	IsFrozen         bool   `force:",omitempty"`
	IsPasswordLocked bool   `force:",omitempty"`
	LastModifiedById string `force:",omitempty"`
	LastModifiedDate string `force:",omitempty"`
	UserId           string `force:",omitempty"`
}

func (t *UserLogin) ApiName() string {
	return "UserLogin"
}

func (t *UserLogin) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("UserLogin #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsFrozen: %v\n", t.IsFrozen))
	builder.WriteString(fmt.Sprintf("\tIsPasswordLocked: %v\n", t.IsPasswordLocked))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tUserId: %v\n", t.UserId))

	return builder.String()
}

type UserLoginQueryResponse struct {
	BaseQuery
	Records []UserLogin `json:"Records" force:"records"`
}
