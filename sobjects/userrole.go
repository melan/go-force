// This file was generated for SObject UserRole, API Version v43.0 at 2018-07-30 03:47:28.779279982 -0400 EDT m=+15.122453856

package sobjects

import (
	"fmt"
	"strings"
)

type UserRole struct {
	BaseSObject
	CaseAccessForAccountOwner        string `force:",omitempty"`
	ContactAccessForAccountOwner     string `force:",omitempty"`
	DeveloperName                    string `force:",omitempty"`
	ForecastUserId                   string `force:",omitempty"`
	Id                               string `force:",omitempty"`
	LastModifiedById                 string `force:",omitempty"`
	LastModifiedDate                 string `force:",omitempty"`
	MayForecastManagerShare          bool   `force:",omitempty"`
	Name                             string `force:",omitempty"`
	OpportunityAccessForAccountOwner string `force:",omitempty"`
	ParentRoleId                     string `force:",omitempty"`
	PortalAccountId                  string `force:",omitempty"`
	PortalAccountOwnerId             string `force:",omitempty"`
	PortalType                       string `force:",omitempty"`
	RollupDescription                string `force:",omitempty"`
	SystemModstamp                   string `force:",omitempty"`
}

func (t *UserRole) ApiName() string {
	return "UserRole"
}

func (t *UserRole) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("UserRole #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCaseAccessForAccountOwner: %v\n", t.CaseAccessForAccountOwner))
	builder.WriteString(fmt.Sprintf("\tContactAccessForAccountOwner: %v\n", t.ContactAccessForAccountOwner))
	builder.WriteString(fmt.Sprintf("\tDeveloperName: %v\n", t.DeveloperName))
	builder.WriteString(fmt.Sprintf("\tForecastUserId: %v\n", t.ForecastUserId))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tMayForecastManagerShare: %v\n", t.MayForecastManagerShare))
	builder.WriteString(fmt.Sprintf("\tName: %v\n", t.Name))
	builder.WriteString(fmt.Sprintf("\tOpportunityAccessForAccountOwner: %v\n", t.OpportunityAccessForAccountOwner))
	builder.WriteString(fmt.Sprintf("\tParentRoleId: %v\n", t.ParentRoleId))
	builder.WriteString(fmt.Sprintf("\tPortalAccountId: %v\n", t.PortalAccountId))
	builder.WriteString(fmt.Sprintf("\tPortalAccountOwnerId: %v\n", t.PortalAccountOwnerId))
	builder.WriteString(fmt.Sprintf("\tPortalType: %v\n", t.PortalType))
	builder.WriteString(fmt.Sprintf("\tRollupDescription: %v\n", t.RollupDescription))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))

	return builder.String()
}

type UserRoleQueryResponse struct {
	BaseQuery
	Records []UserRole `json:"Records" force:"records"`
}
