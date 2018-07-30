// This file was generated for SObject SecureAgent, API Version v43.0 at 2018-07-30 03:47:25.074501777 -0400 EDT m=+11.417536634

package sobjects

import (
	"fmt"
	"strings"
)

type SecureAgent struct {
	BaseSObject
	AgentKey              string `force:",omitempty"`
	CreatedById           string `force:",omitempty"`
	CreatedDate           string `force:",omitempty"`
	DeveloperName         string `force:",omitempty"`
	Id                    string `force:",omitempty"`
	IsDeleted             bool   `force:",omitempty"`
	Language              string `force:",omitempty"`
	LastModifiedById      string `force:",omitempty"`
	LastModifiedDate      string `force:",omitempty"`
	MasterLabel           string `force:",omitempty"`
	Priority              int    `force:",omitempty"`
	ProxyUserId           string `force:",omitempty"`
	SecureAgentsClusterId string `force:",omitempty"`
	SystemModstamp        string `force:",omitempty"`
}

func (t *SecureAgent) ApiName() string {
	return "SecureAgent"
}

func (t *SecureAgent) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("SecureAgent #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tAgentKey: %v\n", t.AgentKey))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tDeveloperName: %v\n", t.DeveloperName))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tLanguage: %v\n", t.Language))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tMasterLabel: %v\n", t.MasterLabel))
	builder.WriteString(fmt.Sprintf("\tPriority: %v\n", t.Priority))
	builder.WriteString(fmt.Sprintf("\tProxyUserId: %v\n", t.ProxyUserId))
	builder.WriteString(fmt.Sprintf("\tSecureAgentsClusterId: %v\n", t.SecureAgentsClusterId))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))

	return builder.String()
}

type SecureAgentQueryResponse struct {
	BaseQuery
	Records []SecureAgent `json:"Records" force:"records"`
}
