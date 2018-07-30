// This file was generated for SObject SecureAgentPluginProperty, API Version v43.0 at 2018-07-30 03:47:46.026506574 -0400 EDT m=+32.370327633

package sobjects

import (
	"fmt"
	"strings"
)

type SecureAgentPluginProperty struct {
	BaseSObject
	CreatedById         string `force:",omitempty"`
	CreatedDate         string `force:",omitempty"`
	Id                  string `force:",omitempty"`
	IsDeleted           bool   `force:",omitempty"`
	LastModifiedById    string `force:",omitempty"`
	LastModifiedDate    string `force:",omitempty"`
	PropertyName        string `force:",omitempty"`
	PropertyValue       string `force:",omitempty"`
	SecureAgentPluginId string `force:",omitempty"`
	SystemModstamp      string `force:",omitempty"`
}

func (t *SecureAgentPluginProperty) ApiName() string {
	return "SecureAgentPluginProperty"
}

func (t *SecureAgentPluginProperty) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("SecureAgentPluginProperty #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tPropertyName: %v\n", t.PropertyName))
	builder.WriteString(fmt.Sprintf("\tPropertyValue: %v\n", t.PropertyValue))
	builder.WriteString(fmt.Sprintf("\tSecureAgentPluginId: %v\n", t.SecureAgentPluginId))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))

	return builder.String()
}

type SecureAgentPluginPropertyQueryResponse struct {
	BaseQuery
	Records []SecureAgentPluginProperty `json:"Records" force:"records"`
}
