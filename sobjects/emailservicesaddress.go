// This file was generated for SObject EmailServicesAddress, API Version v43.0 at 2018-07-30 03:47:50.443265351 -0400 EDT m=+36.787252143

package sobjects

import (
	"fmt"
	"strings"
)

type EmailServicesAddress struct {
	BaseSObject
	AuthorizedSenders string `force:",omitempty"`
	CreatedById       string `force:",omitempty"`
	CreatedDate       string `force:",omitempty"`
	DeveloperName     string `force:",omitempty"`
	EmailDomainName   string `force:",omitempty"`
	FunctionId        string `force:",omitempty"`
	Id                string `force:",omitempty"`
	IsActive          bool   `force:",omitempty"`
	LastModifiedById  string `force:",omitempty"`
	LastModifiedDate  string `force:",omitempty"`
	LocalPart         string `force:",omitempty"`
	RunAsUserId       string `force:",omitempty"`
	SystemModstamp    string `force:",omitempty"`
}

func (t *EmailServicesAddress) ApiName() string {
	return "EmailServicesAddress"
}

func (t *EmailServicesAddress) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("EmailServicesAddress #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tAuthorizedSenders: %v\n", t.AuthorizedSenders))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tDeveloperName: %v\n", t.DeveloperName))
	builder.WriteString(fmt.Sprintf("\tEmailDomainName: %v\n", t.EmailDomainName))
	builder.WriteString(fmt.Sprintf("\tFunctionId: %v\n", t.FunctionId))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsActive: %v\n", t.IsActive))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tLocalPart: %v\n", t.LocalPart))
	builder.WriteString(fmt.Sprintf("\tRunAsUserId: %v\n", t.RunAsUserId))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))

	return builder.String()
}

type EmailServicesAddressQueryResponse struct {
	BaseQuery
	Records []EmailServicesAddress `json:"Records" force:"records"`
}
