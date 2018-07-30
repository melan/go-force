// This file was generated for SObject UserProvisioningLog, API Version v43.0 at 2018-07-30 03:47:45.761880142 -0400 EDT m=+32.105691271

package sobjects

import (
	"fmt"
	"strings"
)

type UserProvisioningLog struct {
	BaseSObject
	CreatedById               string `force:",omitempty"`
	CreatedDate               string `force:",omitempty"`
	Details                   string `force:",omitempty"`
	ExternalUserId            string `force:",omitempty"`
	ExternalUsername          string `force:",omitempty"`
	Id                        string `force:",omitempty"`
	IsDeleted                 bool   `force:",omitempty"`
	LastModifiedById          string `force:",omitempty"`
	LastModifiedDate          string `force:",omitempty"`
	Name                      string `force:",omitempty"`
	Status                    string `force:",omitempty"`
	SystemModstamp            string `force:",omitempty"`
	UserId                    string `force:",omitempty"`
	UserProvisioningRequestId string `force:",omitempty"`
}

func (t *UserProvisioningLog) ApiName() string {
	return "UserProvisioningLog"
}

func (t *UserProvisioningLog) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("UserProvisioningLog #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tDetails: %v\n", t.Details))
	builder.WriteString(fmt.Sprintf("\tExternalUserId: %v\n", t.ExternalUserId))
	builder.WriteString(fmt.Sprintf("\tExternalUsername: %v\n", t.ExternalUsername))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tName: %v\n", t.Name))
	builder.WriteString(fmt.Sprintf("\tStatus: %v\n", t.Status))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))
	builder.WriteString(fmt.Sprintf("\tUserId: %v\n", t.UserId))
	builder.WriteString(fmt.Sprintf("\tUserProvisioningRequestId: %v\n", t.UserProvisioningRequestId))

	return builder.String()
}

type UserProvisioningLogQueryResponse struct {
	BaseQuery
	Records []UserProvisioningLog `json:"Records" force:"records"`
}
