// This file was generated for SObject ContractContactRole, API Version v43.0 at 2018-07-30 03:48:03.875450895 -0400 EDT m=+50.219941716

package sobjects

import (
	"fmt"
	"strings"
)

type ContractContactRole struct {
	BaseSObject
	ContactId        string `force:",omitempty"`
	ContractId       string `force:",omitempty"`
	CreatedById      string `force:",omitempty"`
	CreatedDate      string `force:",omitempty"`
	Id               string `force:",omitempty"`
	IsDeleted        bool   `force:",omitempty"`
	IsPrimary        bool   `force:",omitempty"`
	LastModifiedById string `force:",omitempty"`
	LastModifiedDate string `force:",omitempty"`
	Role             string `force:",omitempty"`
	SystemModstamp   string `force:",omitempty"`
}

func (t *ContractContactRole) ApiName() string {
	return "ContractContactRole"
}

func (t *ContractContactRole) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("ContractContactRole #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tContactId: %v\n", t.ContactId))
	builder.WriteString(fmt.Sprintf("\tContractId: %v\n", t.ContractId))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tIsPrimary: %v\n", t.IsPrimary))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tRole: %v\n", t.Role))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))

	return builder.String()
}

type ContractContactRoleQueryResponse struct {
	BaseQuery
	Records []ContractContactRole `json:"Records" force:"records"`
}
