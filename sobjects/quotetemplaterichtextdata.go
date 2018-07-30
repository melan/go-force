// This file was generated for SObject QuoteTemplateRichTextData, API Version v43.0 at 2018-07-30 03:47:21.490067865 -0400 EDT m=+7.832968218

package sobjects

import (
	"fmt"
	"strings"
)

type QuoteTemplateRichTextData struct {
	BaseSObject
	CreatedById      string `force:",omitempty"`
	CreatedDate      string `force:",omitempty"`
	Data             string `force:",omitempty"`
	Id               string `force:",omitempty"`
	IsDeleted        bool   `force:",omitempty"`
	LastModifiedById string `force:",omitempty"`
	LastModifiedDate string `force:",omitempty"`
	Name             string `force:",omitempty"`
	SystemModstamp   string `force:",omitempty"`
}

func (t *QuoteTemplateRichTextData) ApiName() string {
	return "QuoteTemplateRichTextData"
}

func (t *QuoteTemplateRichTextData) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("QuoteTemplateRichTextData #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tData: %v\n", t.Data))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tName: %v\n", t.Name))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))

	return builder.String()
}

type QuoteTemplateRichTextDataQueryResponse struct {
	BaseQuery
	Records []QuoteTemplateRichTextData `json:"Records" force:"records"`
}
