// This file was generated for SObject NamedCredential, API Version v43.0 at 2018-07-30 03:47:20.867565454 -0400 EDT m=+7.210442449

package sobjects

import (
	"fmt"
	"strings"
)

type NamedCredential struct {
	BaseSObject
	AuthProviderId                            string `force:",omitempty"`
	CalloutOptionsAllowMergeFieldsInBody      bool   `force:",omitempty"`
	CalloutOptionsAllowMergeFieldsInHeader    bool   `force:",omitempty"`
	CalloutOptionsGenerateAuthorizationHeader bool   `force:",omitempty"`
	CreatedById                               string `force:",omitempty"`
	CreatedDate                               string `force:",omitempty"`
	DeveloperName                             string `force:",omitempty"`
	Endpoint                                  string `force:",omitempty"`
	Id                                        string `force:",omitempty"`
	IsDeleted                                 bool   `force:",omitempty"`
	Language                                  string `force:",omitempty"`
	LastModifiedById                          string `force:",omitempty"`
	LastModifiedDate                          string `force:",omitempty"`
	MasterLabel                               string `force:",omitempty"`
	NamespacePrefix                           string `force:",omitempty"`
	PrincipalType                             string `force:",omitempty"`
	SystemModstamp                            string `force:",omitempty"`
}

func (t *NamedCredential) ApiName() string {
	return "NamedCredential"
}

func (t *NamedCredential) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("NamedCredential #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tAuthProviderId: %v\n", t.AuthProviderId))
	builder.WriteString(fmt.Sprintf("\tCalloutOptionsAllowMergeFieldsInBody: %v\n", t.CalloutOptionsAllowMergeFieldsInBody))
	builder.WriteString(fmt.Sprintf("\tCalloutOptionsAllowMergeFieldsInHeader: %v\n", t.CalloutOptionsAllowMergeFieldsInHeader))
	builder.WriteString(fmt.Sprintf("\tCalloutOptionsGenerateAuthorizationHeader: %v\n", t.CalloutOptionsGenerateAuthorizationHeader))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tDeveloperName: %v\n", t.DeveloperName))
	builder.WriteString(fmt.Sprintf("\tEndpoint: %v\n", t.Endpoint))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tLanguage: %v\n", t.Language))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tMasterLabel: %v\n", t.MasterLabel))
	builder.WriteString(fmt.Sprintf("\tNamespacePrefix: %v\n", t.NamespacePrefix))
	builder.WriteString(fmt.Sprintf("\tPrincipalType: %v\n", t.PrincipalType))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))

	return builder.String()
}

type NamedCredentialQueryResponse struct {
	BaseQuery
	Records []NamedCredential `json:"Records" force:"records"`
}
