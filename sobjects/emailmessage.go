// This file was generated for SObject EmailMessage, API Version v43.0 at 2018-07-30 03:47:39.120040184 -0400 EDT m=+25.463602085

package sobjects

import (
	"fmt"
	"strings"
)

type EmailMessage struct {
	BaseSObject
	ActivityId            string `force:",omitempty"`
	BccAddress            string `force:",omitempty"`
	CcAddress             string `force:",omitempty"`
	CreatedById           string `force:",omitempty"`
	CreatedDate           string `force:",omitempty"`
	FromAddress           string `force:",omitempty"`
	FromName              string `force:",omitempty"`
	HasAttachment         bool   `force:",omitempty"`
	Headers               string `force:",omitempty"`
	HtmlBody              string `force:",omitempty"`
	Id                    string `force:",omitempty"`
	Incoming              bool   `force:",omitempty"`
	IsClientManaged       bool   `force:",omitempty"`
	IsDeleted             bool   `force:",omitempty"`
	IsExternallyVisible   bool   `force:",omitempty"`
	LastModifiedById      string `force:",omitempty"`
	LastModifiedDate      string `force:",omitempty"`
	MessageDate           string `force:",omitempty"`
	MessageIdentifier     string `force:",omitempty"`
	ParentId              string `force:",omitempty"`
	RelatedToId           string `force:",omitempty"`
	ReplyToEmailMessageId string `force:",omitempty"`
	Status                string `force:",omitempty"`
	Subject               string `force:",omitempty"`
	SystemModstamp        string `force:",omitempty"`
	TextBody              string `force:",omitempty"`
	ThreadIdentifier      string `force:",omitempty"`
	ToAddress             string `force:",omitempty"`
	ValidatedFromAddress  string `force:",omitempty"`
}

func (t *EmailMessage) ApiName() string {
	return "EmailMessage"
}

func (t *EmailMessage) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("EmailMessage #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tActivityId: %v\n", t.ActivityId))
	builder.WriteString(fmt.Sprintf("\tBccAddress: %v\n", t.BccAddress))
	builder.WriteString(fmt.Sprintf("\tCcAddress: %v\n", t.CcAddress))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tFromAddress: %v\n", t.FromAddress))
	builder.WriteString(fmt.Sprintf("\tFromName: %v\n", t.FromName))
	builder.WriteString(fmt.Sprintf("\tHasAttachment: %v\n", t.HasAttachment))
	builder.WriteString(fmt.Sprintf("\tHeaders: %v\n", t.Headers))
	builder.WriteString(fmt.Sprintf("\tHtmlBody: %v\n", t.HtmlBody))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIncoming: %v\n", t.Incoming))
	builder.WriteString(fmt.Sprintf("\tIsClientManaged: %v\n", t.IsClientManaged))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tIsExternallyVisible: %v\n", t.IsExternallyVisible))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tMessageDate: %v\n", t.MessageDate))
	builder.WriteString(fmt.Sprintf("\tMessageIdentifier: %v\n", t.MessageIdentifier))
	builder.WriteString(fmt.Sprintf("\tParentId: %v\n", t.ParentId))
	builder.WriteString(fmt.Sprintf("\tRelatedToId: %v\n", t.RelatedToId))
	builder.WriteString(fmt.Sprintf("\tReplyToEmailMessageId: %v\n", t.ReplyToEmailMessageId))
	builder.WriteString(fmt.Sprintf("\tStatus: %v\n", t.Status))
	builder.WriteString(fmt.Sprintf("\tSubject: %v\n", t.Subject))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))
	builder.WriteString(fmt.Sprintf("\tTextBody: %v\n", t.TextBody))
	builder.WriteString(fmt.Sprintf("\tThreadIdentifier: %v\n", t.ThreadIdentifier))
	builder.WriteString(fmt.Sprintf("\tToAddress: %v\n", t.ToAddress))
	builder.WriteString(fmt.Sprintf("\tValidatedFromAddress: %v\n", t.ValidatedFromAddress))

	return builder.String()
}

type EmailMessageQueryResponse struct {
	BaseQuery
	Records []EmailMessage `json:"Records" force:"records"`
}
