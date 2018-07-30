// This file was generated for SObject OutgoingEmail, API Version v43.0 at 2018-07-30 03:47:50.862292153 -0400 EDT m=+37.206294669

package sobjects

import (
	"fmt"
	"strings"
)

type OutgoingEmail struct {
	BaseSObject
	BccAddress           string `force:",omitempty"`
	CcAddress            string `force:",omitempty"`
	EmailTemplateId      string `force:",omitempty"`
	ExternalId           string `force:",omitempty"`
	HtmlBody             string `force:",omitempty"`
	Id                   string `force:",omitempty"`
	RelatedToId          string `force:",omitempty"`
	Subject              string `force:",omitempty"`
	TextBody             string `force:",omitempty"`
	ToAddress            string `force:",omitempty"`
	ValidatedFromAddress string `force:",omitempty"`
	WhoId                string `force:",omitempty"`
}

func (t *OutgoingEmail) ApiName() string {
	return "OutgoingEmail"
}

func (t *OutgoingEmail) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("OutgoingEmail #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tBccAddress: %v\n", t.BccAddress))
	builder.WriteString(fmt.Sprintf("\tCcAddress: %v\n", t.CcAddress))
	builder.WriteString(fmt.Sprintf("\tEmailTemplateId: %v\n", t.EmailTemplateId))
	builder.WriteString(fmt.Sprintf("\tExternalId: %v\n", t.ExternalId))
	builder.WriteString(fmt.Sprintf("\tHtmlBody: %v\n", t.HtmlBody))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tRelatedToId: %v\n", t.RelatedToId))
	builder.WriteString(fmt.Sprintf("\tSubject: %v\n", t.Subject))
	builder.WriteString(fmt.Sprintf("\tTextBody: %v\n", t.TextBody))
	builder.WriteString(fmt.Sprintf("\tToAddress: %v\n", t.ToAddress))
	builder.WriteString(fmt.Sprintf("\tValidatedFromAddress: %v\n", t.ValidatedFromAddress))
	builder.WriteString(fmt.Sprintf("\tWhoId: %v\n", t.WhoId))

	return builder.String()
}

type OutgoingEmailQueryResponse struct {
	BaseQuery
	Records []OutgoingEmail `json:"Records" force:"records"`
}
