// This file was generated for SObject EmailServicesFunction, API Version v43.0 at 2018-07-30 03:47:49.931568044 -0400 EDT m=+36.275535636

package sobjects

import (
	"fmt"
	"strings"
)

type EmailServicesFunction struct {
	BaseSObject
	AddressInactiveAction       string `force:",omitempty"`
	ApexClassId                 string `force:",omitempty"`
	AttachmentOption            string `force:",omitempty"`
	AuthenticationFailureAction string `force:",omitempty"`
	AuthorizationFailureAction  string `force:",omitempty"`
	AuthorizedSenders           string `force:",omitempty"`
	CreatedById                 string `force:",omitempty"`
	CreatedDate                 string `force:",omitempty"`
	ErrorRoutingAddress         string `force:",omitempty"`
	FunctionInactiveAction      string `force:",omitempty"`
	FunctionName                string `force:",omitempty"`
	Id                          string `force:",omitempty"`
	IsActive                    bool   `force:",omitempty"`
	IsAuthenticationRequired    bool   `force:",omitempty"`
	IsErrorRoutingEnabled       bool   `force:",omitempty"`
	IsTextAttachmentsAsBinary   bool   `force:",omitempty"`
	IsTlsRequired               bool   `force:",omitempty"`
	LastModifiedById            string `force:",omitempty"`
	LastModifiedDate            string `force:",omitempty"`
	OverLimitAction             string `force:",omitempty"`
	SystemModstamp              string `force:",omitempty"`
}

func (t *EmailServicesFunction) ApiName() string {
	return "EmailServicesFunction"
}

func (t *EmailServicesFunction) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("EmailServicesFunction #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tAddressInactiveAction: %v\n", t.AddressInactiveAction))
	builder.WriteString(fmt.Sprintf("\tApexClassId: %v\n", t.ApexClassId))
	builder.WriteString(fmt.Sprintf("\tAttachmentOption: %v\n", t.AttachmentOption))
	builder.WriteString(fmt.Sprintf("\tAuthenticationFailureAction: %v\n", t.AuthenticationFailureAction))
	builder.WriteString(fmt.Sprintf("\tAuthorizationFailureAction: %v\n", t.AuthorizationFailureAction))
	builder.WriteString(fmt.Sprintf("\tAuthorizedSenders: %v\n", t.AuthorizedSenders))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tErrorRoutingAddress: %v\n", t.ErrorRoutingAddress))
	builder.WriteString(fmt.Sprintf("\tFunctionInactiveAction: %v\n", t.FunctionInactiveAction))
	builder.WriteString(fmt.Sprintf("\tFunctionName: %v\n", t.FunctionName))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsActive: %v\n", t.IsActive))
	builder.WriteString(fmt.Sprintf("\tIsAuthenticationRequired: %v\n", t.IsAuthenticationRequired))
	builder.WriteString(fmt.Sprintf("\tIsErrorRoutingEnabled: %v\n", t.IsErrorRoutingEnabled))
	builder.WriteString(fmt.Sprintf("\tIsTextAttachmentsAsBinary: %v\n", t.IsTextAttachmentsAsBinary))
	builder.WriteString(fmt.Sprintf("\tIsTlsRequired: %v\n", t.IsTlsRequired))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tOverLimitAction: %v\n", t.OverLimitAction))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))

	return builder.String()
}

type EmailServicesFunctionQueryResponse struct {
	BaseQuery
	Records []EmailServicesFunction `json:"Records" force:"records"`
}
