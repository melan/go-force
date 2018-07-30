// This file was generated for SObject UserEntityAccess, API Version v43.0 at 2018-07-30 03:47:27.443627178 -0400 EDT m=+13.786750933

package sobjects

import (
	"fmt"
	"strings"
)

type UserEntityAccess struct {
	BaseSObject
	DurableId          string `force:",omitempty"`
	EntityDefinitionId string `force:",omitempty"`
	Id                 string `force:",omitempty"`
	IsActivateable     bool   `force:",omitempty"`
	IsCreatable        bool   `force:",omitempty"`
	IsDeletable        bool   `force:",omitempty"`
	IsEditable         bool   `force:",omitempty"`
	IsFlsUpdatable     bool   `force:",omitempty"`
	IsMergeable        bool   `force:",omitempty"`
	IsReadable         bool   `force:",omitempty"`
	IsUndeletable      bool   `force:",omitempty"`
	IsUpdatable        bool   `force:",omitempty"`
	UserId             string `force:",omitempty"`
}

func (t *UserEntityAccess) ApiName() string {
	return "UserEntityAccess"
}

func (t *UserEntityAccess) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("UserEntityAccess #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tDurableId: %v\n", t.DurableId))
	builder.WriteString(fmt.Sprintf("\tEntityDefinitionId: %v\n", t.EntityDefinitionId))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsActivateable: %v\n", t.IsActivateable))
	builder.WriteString(fmt.Sprintf("\tIsCreatable: %v\n", t.IsCreatable))
	builder.WriteString(fmt.Sprintf("\tIsDeletable: %v\n", t.IsDeletable))
	builder.WriteString(fmt.Sprintf("\tIsEditable: %v\n", t.IsEditable))
	builder.WriteString(fmt.Sprintf("\tIsFlsUpdatable: %v\n", t.IsFlsUpdatable))
	builder.WriteString(fmt.Sprintf("\tIsMergeable: %v\n", t.IsMergeable))
	builder.WriteString(fmt.Sprintf("\tIsReadable: %v\n", t.IsReadable))
	builder.WriteString(fmt.Sprintf("\tIsUndeletable: %v\n", t.IsUndeletable))
	builder.WriteString(fmt.Sprintf("\tIsUpdatable: %v\n", t.IsUpdatable))
	builder.WriteString(fmt.Sprintf("\tUserId: %v\n", t.UserId))

	return builder.String()
}

type UserEntityAccessQueryResponse struct {
	BaseQuery
	Records []UserEntityAccess `json:"Records" force:"records"`
}
