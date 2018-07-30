// This file was generated for SObject GrantedByLicense, API Version v43.0 at 2018-07-30 03:48:08.305894045 -0400 EDT m=+54.650551114

package sobjects

import (
	"fmt"
	"strings"
)

type GrantedByLicense struct {
	BaseSObject
	CreatedById            string `force:",omitempty"`
	CreatedDate            string `force:",omitempty"`
	CustomPermissionId     string `force:",omitempty"`
	Id                     string `force:",omitempty"`
	IsDeleted              bool   `force:",omitempty"`
	LastModifiedById       string `force:",omitempty"`
	LastModifiedDate       string `force:",omitempty"`
	PermissionSetLicenseId string `force:",omitempty"`
	SystemModstamp         string `force:",omitempty"`
}

func (t *GrantedByLicense) ApiName() string {
	return "GrantedByLicense"
}

func (t *GrantedByLicense) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("GrantedByLicense #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tCustomPermissionId: %v\n", t.CustomPermissionId))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tPermissionSetLicenseId: %v\n", t.PermissionSetLicenseId))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))

	return builder.String()
}

type GrantedByLicenseQueryResponse struct {
	BaseQuery
	Records []GrantedByLicense `json:"Records" force:"records"`
}
