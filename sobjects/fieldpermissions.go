// This file was generated for SObject FieldPermissions, API Version v43.0 at 2018-07-30 03:48:04.702666353 -0400 EDT m=+51.047188215

package sobjects

import (
	"fmt"
	"strings"
)

type FieldPermissions struct {
	BaseSObject
	Field           string `force:",omitempty"`
	Id              string `force:",omitempty"`
	ParentId        string `force:",omitempty"`
	PermissionsEdit bool   `force:",omitempty"`
	PermissionsRead bool   `force:",omitempty"`
	SobjectType     string `force:",omitempty"`
	SystemModstamp  string `force:",omitempty"`
}

func (t *FieldPermissions) ApiName() string {
	return "FieldPermissions"
}

func (t *FieldPermissions) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("FieldPermissions #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tField: %v\n", t.Field))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tParentId: %v\n", t.ParentId))
	builder.WriteString(fmt.Sprintf("\tPermissionsEdit: %v\n", t.PermissionsEdit))
	builder.WriteString(fmt.Sprintf("\tPermissionsRead: %v\n", t.PermissionsRead))
	builder.WriteString(fmt.Sprintf("\tSobjectType: %v\n", t.SobjectType))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))

	return builder.String()
}

type FieldPermissionsQueryResponse struct {
	BaseQuery
	Records []FieldPermissions `json:"Records" force:"records"`
}
