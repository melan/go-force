// This file was generated for SObject ObjectPermissions, API Version v43.0 at 2018-07-30 03:47:45.016576458 -0400 EDT m=+31.360359620

package sobjects

import (
	"fmt"
	"strings"
)

type ObjectPermissions struct {
	BaseSObject
	CreatedById                 string `force:",omitempty"`
	CreatedDate                 string `force:",omitempty"`
	Id                          string `force:",omitempty"`
	LastModifiedById            string `force:",omitempty"`
	LastModifiedDate            string `force:",omitempty"`
	ParentId                    string `force:",omitempty"`
	PermissionsCreate           bool   `force:",omitempty"`
	PermissionsDelete           bool   `force:",omitempty"`
	PermissionsEdit             bool   `force:",omitempty"`
	PermissionsModifyAllRecords bool   `force:",omitempty"`
	PermissionsRead             bool   `force:",omitempty"`
	PermissionsViewAllRecords   bool   `force:",omitempty"`
	SobjectType                 string `force:",omitempty"`
	SystemModstamp              string `force:",omitempty"`
}

func (t *ObjectPermissions) ApiName() string {
	return "ObjectPermissions"
}

func (t *ObjectPermissions) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("ObjectPermissions #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tParentId: %v\n", t.ParentId))
	builder.WriteString(fmt.Sprintf("\tPermissionsCreate: %v\n", t.PermissionsCreate))
	builder.WriteString(fmt.Sprintf("\tPermissionsDelete: %v\n", t.PermissionsDelete))
	builder.WriteString(fmt.Sprintf("\tPermissionsEdit: %v\n", t.PermissionsEdit))
	builder.WriteString(fmt.Sprintf("\tPermissionsModifyAllRecords: %v\n", t.PermissionsModifyAllRecords))
	builder.WriteString(fmt.Sprintf("\tPermissionsRead: %v\n", t.PermissionsRead))
	builder.WriteString(fmt.Sprintf("\tPermissionsViewAllRecords: %v\n", t.PermissionsViewAllRecords))
	builder.WriteString(fmt.Sprintf("\tSobjectType: %v\n", t.SobjectType))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))

	return builder.String()
}

type ObjectPermissionsQueryResponse struct {
	BaseQuery
	Records []ObjectPermissions `json:"Records" force:"records"`
}
