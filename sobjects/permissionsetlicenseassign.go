// This file was generated for SObject PermissionSetLicenseAssign, API Version v43.0 at 2018-07-30 03:48:10.027010672 -0400 EDT m=+56.371732324

package sobjects

import (
	"fmt"
	"strings"
)

type PermissionSetLicenseAssign struct {
	BaseSObject
	AssigneeId             string `force:",omitempty"`
	CreatedById            string `force:",omitempty"`
	CreatedDate            string `force:",omitempty"`
	Id                     string `force:",omitempty"`
	IsDeleted              bool   `force:",omitempty"`
	LastModifiedById       string `force:",omitempty"`
	LastModifiedDate       string `force:",omitempty"`
	PermissionSetLicenseId string `force:",omitempty"`
	SystemModstamp         string `force:",omitempty"`
}

func (t *PermissionSetLicenseAssign) ApiName() string {
	return "PermissionSetLicenseAssign"
}

func (t *PermissionSetLicenseAssign) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("PermissionSetLicenseAssign #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tAssigneeId: %v\n", t.AssigneeId))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tPermissionSetLicenseId: %v\n", t.PermissionSetLicenseId))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))

	return builder.String()
}

type PermissionSetLicenseAssignQueryResponse struct {
	BaseQuery
	Records []PermissionSetLicenseAssign `json:"Records" force:"records"`
}
