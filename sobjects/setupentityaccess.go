// This file was generated for SObject SetupEntityAccess, API Version v43.0 at 2018-07-30 03:47:54.085654366 -0400 EDT m=+40.429777836

package sobjects

import (
	"fmt"
	"strings"
)

type SetupEntityAccess struct {
	BaseSObject
	Id              string `force:",omitempty"`
	ParentId        string `force:",omitempty"`
	SetupEntityId   string `force:",omitempty"`
	SetupEntityType string `force:",omitempty"`
	SystemModstamp  string `force:",omitempty"`
}

func (t *SetupEntityAccess) ApiName() string {
	return "SetupEntityAccess"
}

func (t *SetupEntityAccess) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("SetupEntityAccess #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tParentId: %v\n", t.ParentId))
	builder.WriteString(fmt.Sprintf("\tSetupEntityId: %v\n", t.SetupEntityId))
	builder.WriteString(fmt.Sprintf("\tSetupEntityType: %v\n", t.SetupEntityType))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))

	return builder.String()
}

type SetupEntityAccessQueryResponse struct {
	BaseQuery
	Records []SetupEntityAccess `json:"Records" force:"records"`
}
