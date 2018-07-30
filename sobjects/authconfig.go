// This file was generated for SObject AuthConfig, API Version v43.0 at 2018-07-30 03:47:40.746743146 -0400 EDT m=+27.090366087

package sobjects

import (
	"fmt"
	"strings"
)

type AuthConfig struct {
	BaseSObject
	AuthOptionsAuthProvider     bool   `force:",omitempty"`
	AuthOptionsSaml             bool   `force:",omitempty"`
	AuthOptionsUsernamePassword bool   `force:",omitempty"`
	CreatedById                 string `force:",omitempty"`
	CreatedDate                 string `force:",omitempty"`
	DeveloperName               string `force:",omitempty"`
	Id                          string `force:",omitempty"`
	IsActive                    bool   `force:",omitempty"`
	IsDeleted                   bool   `force:",omitempty"`
	Language                    string `force:",omitempty"`
	LastModifiedById            string `force:",omitempty"`
	LastModifiedDate            string `force:",omitempty"`
	MasterLabel                 string `force:",omitempty"`
	NamespacePrefix             string `force:",omitempty"`
	SystemModstamp              string `force:",omitempty"`
	Type                        string `force:",omitempty"`
	Url                         string `force:",omitempty"`
}

func (t *AuthConfig) ApiName() string {
	return "AuthConfig"
}

func (t *AuthConfig) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("AuthConfig #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tAuthOptionsAuthProvider: %v\n", t.AuthOptionsAuthProvider))
	builder.WriteString(fmt.Sprintf("\tAuthOptionsSaml: %v\n", t.AuthOptionsSaml))
	builder.WriteString(fmt.Sprintf("\tAuthOptionsUsernamePassword: %v\n", t.AuthOptionsUsernamePassword))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tDeveloperName: %v\n", t.DeveloperName))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsActive: %v\n", t.IsActive))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tLanguage: %v\n", t.Language))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tMasterLabel: %v\n", t.MasterLabel))
	builder.WriteString(fmt.Sprintf("\tNamespacePrefix: %v\n", t.NamespacePrefix))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))
	builder.WriteString(fmt.Sprintf("\tType: %v\n", t.Type))
	builder.WriteString(fmt.Sprintf("\tUrl: %v\n", t.Url))

	return builder.String()
}

type AuthConfigQueryResponse struct {
	BaseQuery
	Records []AuthConfig `json:"Records" force:"records"`
}
