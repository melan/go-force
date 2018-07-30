// This file was generated for SObject UserProvisioningConfig, API Version v43.0 at 2018-07-30 03:48:07.502685175 -0400 EDT m=+53.847312104

package sobjects

import (
	"fmt"
	"strings"
)

type UserProvisioningConfig struct {
	BaseSObject
	ApprovalRequired   string `force:",omitempty"`
	ConnectedAppId     string `force:",omitempty"`
	CreatedById        string `force:",omitempty"`
	CreatedDate        string `force:",omitempty"`
	DeveloperName      string `force:",omitempty"`
	Enabled            bool   `force:",omitempty"`
	EnabledOperations  string `force:",omitempty"`
	Id                 string `force:",omitempty"`
	IsDeleted          bool   `force:",omitempty"`
	Language           string `force:",omitempty"`
	LastModifiedById   string `force:",omitempty"`
	LastModifiedDate   string `force:",omitempty"`
	LastReconDateTime  string `force:",omitempty"`
	MasterLabel        string `force:",omitempty"`
	NamedCredentialId  string `force:",omitempty"`
	NamespacePrefix    string `force:",omitempty"`
	Notes              string `force:",omitempty"`
	OnUpdateAttributes string `force:",omitempty"`
	ReconFilter        string `force:",omitempty"`
	SystemModstamp     string `force:",omitempty"`
	UserAccountMapping string `force:",omitempty"`
}

func (t *UserProvisioningConfig) ApiName() string {
	return "UserProvisioningConfig"
}

func (t *UserProvisioningConfig) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("UserProvisioningConfig #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tApprovalRequired: %v\n", t.ApprovalRequired))
	builder.WriteString(fmt.Sprintf("\tConnectedAppId: %v\n", t.ConnectedAppId))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tDeveloperName: %v\n", t.DeveloperName))
	builder.WriteString(fmt.Sprintf("\tEnabled: %v\n", t.Enabled))
	builder.WriteString(fmt.Sprintf("\tEnabledOperations: %v\n", t.EnabledOperations))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tLanguage: %v\n", t.Language))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tLastReconDateTime: %v\n", t.LastReconDateTime))
	builder.WriteString(fmt.Sprintf("\tMasterLabel: %v\n", t.MasterLabel))
	builder.WriteString(fmt.Sprintf("\tNamedCredentialId: %v\n", t.NamedCredentialId))
	builder.WriteString(fmt.Sprintf("\tNamespacePrefix: %v\n", t.NamespacePrefix))
	builder.WriteString(fmt.Sprintf("\tNotes: %v\n", t.Notes))
	builder.WriteString(fmt.Sprintf("\tOnUpdateAttributes: %v\n", t.OnUpdateAttributes))
	builder.WriteString(fmt.Sprintf("\tReconFilter: %v\n", t.ReconFilter))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))
	builder.WriteString(fmt.Sprintf("\tUserAccountMapping: %v\n", t.UserAccountMapping))

	return builder.String()
}

type UserProvisioningConfigQueryResponse struct {
	BaseQuery
	Records []UserProvisioningConfig `json:"Records" force:"records"`
}
