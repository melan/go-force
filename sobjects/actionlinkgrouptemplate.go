// This file was generated for SObject ActionLinkGroupTemplate, API Version v43.0 at 2018-07-30 03:47:28.282386531 -0400 EDT m=+14.625541760

package sobjects

import (
	"fmt"
	"strings"
)

type ActionLinkGroupTemplate struct {
	BaseSObject
	Category             string `force:",omitempty"`
	CreatedById          string `force:",omitempty"`
	CreatedDate          string `force:",omitempty"`
	DeveloperName        string `force:",omitempty"`
	ExecutionsAllowed    string `force:",omitempty"`
	HoursUntilExpiration int    `force:",omitempty"`
	Id                   string `force:",omitempty"`
	IsDeleted            bool   `force:",omitempty"`
	IsPublished          bool   `force:",omitempty"`
	Language             string `force:",omitempty"`
	LastModifiedById     string `force:",omitempty"`
	LastModifiedDate     string `force:",omitempty"`
	MasterLabel          string `force:",omitempty"`
	NamespacePrefix      string `force:",omitempty"`
	SystemModstamp       string `force:",omitempty"`
}

func (t *ActionLinkGroupTemplate) ApiName() string {
	return "ActionLinkGroupTemplate"
}

func (t *ActionLinkGroupTemplate) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("ActionLinkGroupTemplate #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCategory: %v\n", t.Category))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tDeveloperName: %v\n", t.DeveloperName))
	builder.WriteString(fmt.Sprintf("\tExecutionsAllowed: %v\n", t.ExecutionsAllowed))
	builder.WriteString(fmt.Sprintf("\tHoursUntilExpiration: %v\n", t.HoursUntilExpiration))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tIsPublished: %v\n", t.IsPublished))
	builder.WriteString(fmt.Sprintf("\tLanguage: %v\n", t.Language))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tMasterLabel: %v\n", t.MasterLabel))
	builder.WriteString(fmt.Sprintf("\tNamespacePrefix: %v\n", t.NamespacePrefix))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))

	return builder.String()
}

type ActionLinkGroupTemplateQueryResponse struct {
	BaseQuery
	Records []ActionLinkGroupTemplate `json:"Records" force:"records"`
}
