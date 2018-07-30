// This file was generated for SObject EmailTemplate, API Version v43.0 at 2018-07-30 03:47:54.2332133 -0400 EDT m=+40.577342306

package sobjects

import (
	"fmt"
	"strings"
)

type EmailTemplate struct {
	BaseSObject
	ApiVersion        float64 `force:",omitempty"`
	Body              string  `force:",omitempty"`
	BrandTemplateId   string  `force:",omitempty"`
	CreatedById       string  `force:",omitempty"`
	CreatedDate       string  `force:",omitempty"`
	Description       string  `force:",omitempty"`
	DeveloperName     string  `force:",omitempty"`
	Encoding          string  `force:",omitempty"`
	FolderId          string  `force:",omitempty"`
	HtmlValue         string  `force:",omitempty"`
	Id                string  `force:",omitempty"`
	IsActive          bool    `force:",omitempty"`
	LastModifiedById  string  `force:",omitempty"`
	LastModifiedDate  string  `force:",omitempty"`
	LastUsedDate      string  `force:",omitempty"`
	Markup            string  `force:",omitempty"`
	Name              string  `force:",omitempty"`
	NamespacePrefix   string  `force:",omitempty"`
	OwnerId           string  `force:",omitempty"`
	RelatedEntityType string  `force:",omitempty"`
	Subject           string  `force:",omitempty"`
	SystemModstamp    string  `force:",omitempty"`
	TemplateStyle     string  `force:",omitempty"`
	TemplateType      string  `force:",omitempty"`
	TimesUsed         int     `force:",omitempty"`
	UiType            string  `force:",omitempty"`
}

func (t *EmailTemplate) ApiName() string {
	return "EmailTemplate"
}

func (t *EmailTemplate) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("EmailTemplate #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tApiVersion: %v\n", t.ApiVersion))
	builder.WriteString(fmt.Sprintf("\tBody: %v\n", t.Body))
	builder.WriteString(fmt.Sprintf("\tBrandTemplateId: %v\n", t.BrandTemplateId))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tDescription: %v\n", t.Description))
	builder.WriteString(fmt.Sprintf("\tDeveloperName: %v\n", t.DeveloperName))
	builder.WriteString(fmt.Sprintf("\tEncoding: %v\n", t.Encoding))
	builder.WriteString(fmt.Sprintf("\tFolderId: %v\n", t.FolderId))
	builder.WriteString(fmt.Sprintf("\tHtmlValue: %v\n", t.HtmlValue))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsActive: %v\n", t.IsActive))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tLastUsedDate: %v\n", t.LastUsedDate))
	builder.WriteString(fmt.Sprintf("\tMarkup: %v\n", t.Markup))
	builder.WriteString(fmt.Sprintf("\tName: %v\n", t.Name))
	builder.WriteString(fmt.Sprintf("\tNamespacePrefix: %v\n", t.NamespacePrefix))
	builder.WriteString(fmt.Sprintf("\tOwnerId: %v\n", t.OwnerId))
	builder.WriteString(fmt.Sprintf("\tRelatedEntityType: %v\n", t.RelatedEntityType))
	builder.WriteString(fmt.Sprintf("\tSubject: %v\n", t.Subject))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))
	builder.WriteString(fmt.Sprintf("\tTemplateStyle: %v\n", t.TemplateStyle))
	builder.WriteString(fmt.Sprintf("\tTemplateType: %v\n", t.TemplateType))
	builder.WriteString(fmt.Sprintf("\tTimesUsed: %v\n", t.TimesUsed))
	builder.WriteString(fmt.Sprintf("\tUiType: %v\n", t.UiType))

	return builder.String()
}

type EmailTemplateQueryResponse struct {
	BaseQuery
	Records []EmailTemplate `json:"Records" force:"records"`
}
