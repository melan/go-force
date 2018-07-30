// This file was generated for SObject TransactionSecurityPolicy, API Version v43.0 at 2018-07-30 03:47:29.71884162 -0400 EDT m=+16.062050750

package sobjects

import (
	"fmt"
	"strings"
)

type TransactionSecurityPolicy struct {
	BaseSObject
	ActionConfig     string `force:",omitempty"`
	ApexPolicyId     string `force:",omitempty"`
	CreatedById      string `force:",omitempty"`
	CreatedDate      string `force:",omitempty"`
	Description      string `force:",omitempty"`
	DeveloperName    string `force:",omitempty"`
	EventType        string `force:",omitempty"`
	ExecutionUserId  string `force:",omitempty"`
	Id               string `force:",omitempty"`
	IsDeleted        bool   `force:",omitempty"`
	Language         string `force:",omitempty"`
	LastModifiedById string `force:",omitempty"`
	LastModifiedDate string `force:",omitempty"`
	MasterLabel      string `force:",omitempty"`
	NamespacePrefix  string `force:",omitempty"`
	ResourceName     string `force:",omitempty"`
	State            string `force:",omitempty"`
	SystemModstamp   string `force:",omitempty"`
	Type             string `force:",omitempty"`
}

func (t *TransactionSecurityPolicy) ApiName() string {
	return "TransactionSecurityPolicy"
}

func (t *TransactionSecurityPolicy) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("TransactionSecurityPolicy #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tActionConfig: %v\n", t.ActionConfig))
	builder.WriteString(fmt.Sprintf("\tApexPolicyId: %v\n", t.ApexPolicyId))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tDescription: %v\n", t.Description))
	builder.WriteString(fmt.Sprintf("\tDeveloperName: %v\n", t.DeveloperName))
	builder.WriteString(fmt.Sprintf("\tEventType: %v\n", t.EventType))
	builder.WriteString(fmt.Sprintf("\tExecutionUserId: %v\n", t.ExecutionUserId))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tLanguage: %v\n", t.Language))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tMasterLabel: %v\n", t.MasterLabel))
	builder.WriteString(fmt.Sprintf("\tNamespacePrefix: %v\n", t.NamespacePrefix))
	builder.WriteString(fmt.Sprintf("\tResourceName: %v\n", t.ResourceName))
	builder.WriteString(fmt.Sprintf("\tState: %v\n", t.State))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))
	builder.WriteString(fmt.Sprintf("\tType: %v\n", t.Type))

	return builder.String()
}

type TransactionSecurityPolicyQueryResponse struct {
	BaseQuery
	Records []TransactionSecurityPolicy `json:"Records" force:"records"`
}
