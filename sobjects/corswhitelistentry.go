// This file was generated for SObject CorsWhitelistEntry, API Version v43.0 at 2018-07-30 03:47:20.993007571 -0400 EDT m=+7.335889273

package sobjects

import (
	"fmt"
	"strings"
)

type CorsWhitelistEntry struct {
	BaseSObject
	CreatedById      string `force:",omitempty"`
	CreatedDate      string `force:",omitempty"`
	DeveloperName    string `force:",omitempty"`
	Id               string `force:",omitempty"`
	IsDeleted        bool   `force:",omitempty"`
	Language         string `force:",omitempty"`
	LastModifiedById string `force:",omitempty"`
	LastModifiedDate string `force:",omitempty"`
	MasterLabel      string `force:",omitempty"`
	NamespacePrefix  string `force:",omitempty"`
	SystemModstamp   string `force:",omitempty"`
	UrlPattern       string `force:",omitempty"`
}

func (t *CorsWhitelistEntry) ApiName() string {
	return "CorsWhitelistEntry"
}

func (t *CorsWhitelistEntry) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("CorsWhitelistEntry #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tDeveloperName: %v\n", t.DeveloperName))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tLanguage: %v\n", t.Language))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tMasterLabel: %v\n", t.MasterLabel))
	builder.WriteString(fmt.Sprintf("\tNamespacePrefix: %v\n", t.NamespacePrefix))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))
	builder.WriteString(fmt.Sprintf("\tUrlPattern: %v\n", t.UrlPattern))

	return builder.String()
}

type CorsWhitelistEntryQueryResponse struct {
	BaseQuery
	Records []CorsWhitelistEntry `json:"Records" force:"records"`
}
