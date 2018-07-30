// This file was generated for SObject UserFieldAccess, API Version v43.0 at 2018-07-30 03:47:38.732260418 -0400 EDT m=+25.075807767

package sobjects

import (
	"fmt"
	"strings"
)

type UserFieldAccess struct {
	BaseSObject
	DurableId          string `force:",omitempty"`
	EntityDefinitionId string `force:",omitempty"`
	FieldDefinitionId  string `force:",omitempty"`
	Id                 string `force:",omitempty"`
	IsAccessible       bool   `force:",omitempty"`
	IsCreatable        bool   `force:",omitempty"`
	IsUpdatable        bool   `force:",omitempty"`
	UserId             string `force:",omitempty"`
}

func (t *UserFieldAccess) ApiName() string {
	return "UserFieldAccess"
}

func (t *UserFieldAccess) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("UserFieldAccess #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tDurableId: %v\n", t.DurableId))
	builder.WriteString(fmt.Sprintf("\tEntityDefinitionId: %v\n", t.EntityDefinitionId))
	builder.WriteString(fmt.Sprintf("\tFieldDefinitionId: %v\n", t.FieldDefinitionId))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsAccessible: %v\n", t.IsAccessible))
	builder.WriteString(fmt.Sprintf("\tIsCreatable: %v\n", t.IsCreatable))
	builder.WriteString(fmt.Sprintf("\tIsUpdatable: %v\n", t.IsUpdatable))
	builder.WriteString(fmt.Sprintf("\tUserId: %v\n", t.UserId))

	return builder.String()
}

type UserFieldAccessQueryResponse struct {
	BaseQuery
	Records []UserFieldAccess `json:"Records" force:"records"`
}
