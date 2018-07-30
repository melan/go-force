// This file was generated for SObject SamlSsoConfig, API Version v43.0 at 2018-07-30 03:47:58.43630138 -0400 EDT m=+44.780588103

package sobjects

import (
	"fmt"
	"strings"
)

type SamlSsoConfig struct {
	BaseSObject
	AttributeFormat         string `force:",omitempty"`
	AttributeName           string `force:",omitempty"`
	Audience                string `force:",omitempty"`
	CreatedById             string `force:",omitempty"`
	CreatedDate             string `force:",omitempty"`
	DeveloperName           string `force:",omitempty"`
	ErrorUrl                string `force:",omitempty"`
	ExecutionUserId         string `force:",omitempty"`
	Id                      string `force:",omitempty"`
	IdentityLocation        string `force:",omitempty"`
	IdentityMapping         string `force:",omitempty"`
	IsDeleted               bool   `force:",omitempty"`
	Issuer                  string `force:",omitempty"`
	Language                string `force:",omitempty"`
	LastModifiedById        string `force:",omitempty"`
	LastModifiedDate        string `force:",omitempty"`
	LoginUrl                string `force:",omitempty"`
	LogoutUrl               string `force:",omitempty"`
	MasterLabel             string `force:",omitempty"`
	NamespacePrefix         string `force:",omitempty"`
	OptionsSpInitBinding    bool   `force:",omitempty"`
	OptionsUserProvisioning bool   `force:",omitempty"`
	RequestSignatureMethod  string `force:",omitempty"`
	SamlJitHandlerId        string `force:",omitempty"`
	SingleLogoutBinding     string `force:",omitempty"`
	SingleLogoutUrl         string `force:",omitempty"`
	SystemModstamp          string `force:",omitempty"`
	ValidationCert          string `force:",omitempty"`
	Version                 string `force:",omitempty"`
}

func (t *SamlSsoConfig) ApiName() string {
	return "SamlSsoConfig"
}

func (t *SamlSsoConfig) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("SamlSsoConfig #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tAttributeFormat: %v\n", t.AttributeFormat))
	builder.WriteString(fmt.Sprintf("\tAttributeName: %v\n", t.AttributeName))
	builder.WriteString(fmt.Sprintf("\tAudience: %v\n", t.Audience))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tDeveloperName: %v\n", t.DeveloperName))
	builder.WriteString(fmt.Sprintf("\tErrorUrl: %v\n", t.ErrorUrl))
	builder.WriteString(fmt.Sprintf("\tExecutionUserId: %v\n", t.ExecutionUserId))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIdentityLocation: %v\n", t.IdentityLocation))
	builder.WriteString(fmt.Sprintf("\tIdentityMapping: %v\n", t.IdentityMapping))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tIssuer: %v\n", t.Issuer))
	builder.WriteString(fmt.Sprintf("\tLanguage: %v\n", t.Language))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tLoginUrl: %v\n", t.LoginUrl))
	builder.WriteString(fmt.Sprintf("\tLogoutUrl: %v\n", t.LogoutUrl))
	builder.WriteString(fmt.Sprintf("\tMasterLabel: %v\n", t.MasterLabel))
	builder.WriteString(fmt.Sprintf("\tNamespacePrefix: %v\n", t.NamespacePrefix))
	builder.WriteString(fmt.Sprintf("\tOptionsSpInitBinding: %v\n", t.OptionsSpInitBinding))
	builder.WriteString(fmt.Sprintf("\tOptionsUserProvisioning: %v\n", t.OptionsUserProvisioning))
	builder.WriteString(fmt.Sprintf("\tRequestSignatureMethod: %v\n", t.RequestSignatureMethod))
	builder.WriteString(fmt.Sprintf("\tSamlJitHandlerId: %v\n", t.SamlJitHandlerId))
	builder.WriteString(fmt.Sprintf("\tSingleLogoutBinding: %v\n", t.SingleLogoutBinding))
	builder.WriteString(fmt.Sprintf("\tSingleLogoutUrl: %v\n", t.SingleLogoutUrl))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))
	builder.WriteString(fmt.Sprintf("\tValidationCert: %v\n", t.ValidationCert))
	builder.WriteString(fmt.Sprintf("\tVersion: %v\n", t.Version))

	return builder.String()
}

type SamlSsoConfigQueryResponse struct {
	BaseQuery
	Records []SamlSsoConfig `json:"Records" force:"records"`
}
