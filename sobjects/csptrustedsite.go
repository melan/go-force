// This file was generated for SObject CspTrustedSite, API Version v43.0 at 2018-07-30 03:48:05.253162813 -0400 EDT m=+51.597705331

package sobjects

import (
	"fmt"
	"strings"
)

type CspTrustedSite struct {
	BaseSObject
	CreatedById      string `force:",omitempty"`
	CreatedDate      string `force:",omitempty"`
	Description      string `force:",omitempty"`
	DeveloperName    string `force:",omitempty"`
	EndpointUrl      string `force:",omitempty"`
	Id               string `force:",omitempty"`
	IsActive         bool   `force:",omitempty"`
	IsDeleted        bool   `force:",omitempty"`
	Language         string `force:",omitempty"`
	LastModifiedById string `force:",omitempty"`
	LastModifiedDate string `force:",omitempty"`
	MasterLabel      string `force:",omitempty"`
	NamespacePrefix  string `force:",omitempty"`
	SystemModstamp   string `force:",omitempty"`
}

func (t *CspTrustedSite) ApiName() string {
	return "CspTrustedSite"
}

func (t *CspTrustedSite) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("CspTrustedSite #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tDescription: %v\n", t.Description))
	builder.WriteString(fmt.Sprintf("\tDeveloperName: %v\n", t.DeveloperName))
	builder.WriteString(fmt.Sprintf("\tEndpointUrl: %v\n", t.EndpointUrl))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsActive: %v\n", t.IsActive))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tLanguage: %v\n", t.Language))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tMasterLabel: %v\n", t.MasterLabel))
	builder.WriteString(fmt.Sprintf("\tNamespacePrefix: %v\n", t.NamespacePrefix))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))

	return builder.String()
}

type CspTrustedSiteQueryResponse struct {
	BaseQuery
	Records []CspTrustedSite `json:"Records" force:"records"`
}
