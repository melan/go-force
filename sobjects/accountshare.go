// This file was generated for SObject AccountShare, API Version v43.0 at 2018-07-30 03:47:41.163590917 -0400 EDT m=+27.507229499

package sobjects

import (
	"fmt"
	"strings"
)

type AccountShare struct {
	BaseSObject
	AccountAccessLevel     string `force:",omitempty"`
	AccountId              string `force:",omitempty"`
	CaseAccessLevel        string `force:",omitempty"`
	ContactAccessLevel     string `force:",omitempty"`
	Id                     string `force:",omitempty"`
	IsDeleted              bool   `force:",omitempty"`
	LastModifiedById       string `force:",omitempty"`
	LastModifiedDate       string `force:",omitempty"`
	OpportunityAccessLevel string `force:",omitempty"`
	RowCause               string `force:",omitempty"`
	UserOrGroupId          string `force:",omitempty"`
}

func (t *AccountShare) ApiName() string {
	return "AccountShare"
}

func (t *AccountShare) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("AccountShare #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tAccountAccessLevel: %v\n", t.AccountAccessLevel))
	builder.WriteString(fmt.Sprintf("\tAccountId: %v\n", t.AccountId))
	builder.WriteString(fmt.Sprintf("\tCaseAccessLevel: %v\n", t.CaseAccessLevel))
	builder.WriteString(fmt.Sprintf("\tContactAccessLevel: %v\n", t.ContactAccessLevel))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tOpportunityAccessLevel: %v\n", t.OpportunityAccessLevel))
	builder.WriteString(fmt.Sprintf("\tRowCause: %v\n", t.RowCause))
	builder.WriteString(fmt.Sprintf("\tUserOrGroupId: %v\n", t.UserOrGroupId))

	return builder.String()
}

type AccountShareQueryResponse struct {
	BaseQuery
	Records []AccountShare `json:"Records" force:"records"`
}
