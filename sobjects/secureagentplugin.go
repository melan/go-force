// This file was generated for SObject SecureAgentPlugin, API Version v43.0 at 2018-07-30 03:47:32.623052705 -0400 EDT m=+18.966370813

package sobjects

import (
	"fmt"
	"strings"
)

type SecureAgentPlugin struct {
	BaseSObject
	CreatedById       string `force:",omitempty"`
	CreatedDate       string `force:",omitempty"`
	Id                string `force:",omitempty"`
	IsDeleted         bool   `force:",omitempty"`
	LastModifiedById  string `force:",omitempty"`
	LastModifiedDate  string `force:",omitempty"`
	PluginName        string `force:",omitempty"`
	PluginType        string `force:",omitempty"`
	RequestedVersion  string `force:",omitempty"`
	SecureAgentId     string `force:",omitempty"`
	SystemModstamp    string `force:",omitempty"`
	UpdateWindowEnd   string `force:",omitempty"`
	UpdateWindowStart string `force:",omitempty"`
}

func (t *SecureAgentPlugin) ApiName() string {
	return "SecureAgentPlugin"
}

func (t *SecureAgentPlugin) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("SecureAgentPlugin #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tPluginName: %v\n", t.PluginName))
	builder.WriteString(fmt.Sprintf("\tPluginType: %v\n", t.PluginType))
	builder.WriteString(fmt.Sprintf("\tRequestedVersion: %v\n", t.RequestedVersion))
	builder.WriteString(fmt.Sprintf("\tSecureAgentId: %v\n", t.SecureAgentId))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))
	builder.WriteString(fmt.Sprintf("\tUpdateWindowEnd: %v\n", t.UpdateWindowEnd))
	builder.WriteString(fmt.Sprintf("\tUpdateWindowStart: %v\n", t.UpdateWindowStart))

	return builder.String()
}

type SecureAgentPluginQueryResponse struct {
	BaseQuery
	Records []SecureAgentPlugin `json:"Records" force:"records"`
}
