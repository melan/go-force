// This file was generated for SObject UserProvisioningRequest, API Version v43.0 at 2018-07-30 03:47:16.474077394 -0400 EDT m=+2.816789526

package sobjects

import (
	"fmt"
	"strings"
)

type UserProvisioningRequest struct {
	BaseSObject
	AppName           string `force:",omitempty"`
	ApprovalStatus    string `force:",omitempty"`
	ConnectedAppId    string `force:",omitempty"`
	CreatedById       string `force:",omitempty"`
	CreatedDate       string `force:",omitempty"`
	ExternalUserId    string `force:",omitempty"`
	Id                string `force:",omitempty"`
	IsDeleted         bool   `force:",omitempty"`
	LastModifiedById  string `force:",omitempty"`
	LastModifiedDate  string `force:",omitempty"`
	ManagerId         string `force:",omitempty"`
	Name              string `force:",omitempty"`
	Operation         string `force:",omitempty"`
	OwnerId           string `force:",omitempty"`
	ParentId          string `force:",omitempty"`
	RetryCount        int    `force:",omitempty"`
	SalesforceUserId  string `force:",omitempty"`
	ScheduleDate      string `force:",omitempty"`
	State             string `force:",omitempty"`
	SystemModstamp    string `force:",omitempty"`
	UserProvAccountId string `force:",omitempty"`
	UserProvConfigId  string `force:",omitempty"`
}

func (t *UserProvisioningRequest) ApiName() string {
	return "UserProvisioningRequest"
}

func (t *UserProvisioningRequest) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("UserProvisioningRequest #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tAppName: %v\n", t.AppName))
	builder.WriteString(fmt.Sprintf("\tApprovalStatus: %v\n", t.ApprovalStatus))
	builder.WriteString(fmt.Sprintf("\tConnectedAppId: %v\n", t.ConnectedAppId))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tExternalUserId: %v\n", t.ExternalUserId))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tManagerId: %v\n", t.ManagerId))
	builder.WriteString(fmt.Sprintf("\tName: %v\n", t.Name))
	builder.WriteString(fmt.Sprintf("\tOperation: %v\n", t.Operation))
	builder.WriteString(fmt.Sprintf("\tOwnerId: %v\n", t.OwnerId))
	builder.WriteString(fmt.Sprintf("\tParentId: %v\n", t.ParentId))
	builder.WriteString(fmt.Sprintf("\tRetryCount: %v\n", t.RetryCount))
	builder.WriteString(fmt.Sprintf("\tSalesforceUserId: %v\n", t.SalesforceUserId))
	builder.WriteString(fmt.Sprintf("\tScheduleDate: %v\n", t.ScheduleDate))
	builder.WriteString(fmt.Sprintf("\tState: %v\n", t.State))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))
	builder.WriteString(fmt.Sprintf("\tUserProvAccountId: %v\n", t.UserProvAccountId))
	builder.WriteString(fmt.Sprintf("\tUserProvConfigId: %v\n", t.UserProvConfigId))

	return builder.String()
}

type UserProvisioningRequestQueryResponse struct {
	BaseQuery
	Records []UserProvisioningRequest `json:"Records" force:"records"`
}
