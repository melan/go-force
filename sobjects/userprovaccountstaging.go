// This file was generated for SObject UserProvAccountStaging, API Version v43.0 at 2018-07-30 03:47:54.639201989 -0400 EDT m=+40.983346230

package sobjects

import (
	"fmt"
	"strings"
)

type UserProvAccountStaging struct {
	BaseSObject
	ConnectedAppId    string `force:",omitempty"`
	CreatedById       string `force:",omitempty"`
	CreatedDate       string `force:",omitempty"`
	ExternalEmail     string `force:",omitempty"`
	ExternalFirstName string `force:",omitempty"`
	ExternalLastName  string `force:",omitempty"`
	ExternalUserId    string `force:",omitempty"`
	ExternalUsername  string `force:",omitempty"`
	Id                string `force:",omitempty"`
	IsDeleted         bool   `force:",omitempty"`
	LastModifiedById  string `force:",omitempty"`
	LastModifiedDate  string `force:",omitempty"`
	LinkState         string `force:",omitempty"`
	Name              string `force:",omitempty"`
	SalesforceUserId  string `force:",omitempty"`
	Status            string `force:",omitempty"`
	SystemModstamp    string `force:",omitempty"`
}

func (t *UserProvAccountStaging) ApiName() string {
	return "UserProvAccountStaging"
}

func (t *UserProvAccountStaging) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("UserProvAccountStaging #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tConnectedAppId: %v\n", t.ConnectedAppId))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tExternalEmail: %v\n", t.ExternalEmail))
	builder.WriteString(fmt.Sprintf("\tExternalFirstName: %v\n", t.ExternalFirstName))
	builder.WriteString(fmt.Sprintf("\tExternalLastName: %v\n", t.ExternalLastName))
	builder.WriteString(fmt.Sprintf("\tExternalUserId: %v\n", t.ExternalUserId))
	builder.WriteString(fmt.Sprintf("\tExternalUsername: %v\n", t.ExternalUsername))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tLinkState: %v\n", t.LinkState))
	builder.WriteString(fmt.Sprintf("\tName: %v\n", t.Name))
	builder.WriteString(fmt.Sprintf("\tSalesforceUserId: %v\n", t.SalesforceUserId))
	builder.WriteString(fmt.Sprintf("\tStatus: %v\n", t.Status))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))

	return builder.String()
}

type UserProvAccountStagingQueryResponse struct {
	BaseQuery
	Records []UserProvAccountStaging `json:"Records" force:"records"`
}
