// This file was generated for SObject MailmergeTemplate, API Version v43.0 at 2018-07-30 03:47:48.022520523 -0400 EDT m=+34.366416480

package sobjects

import (
	"fmt"
	"strings"
)

type MailmergeTemplate struct {
	BaseSObject
	Body                                     string `force:",omitempty"`
	BodyLength                               int    `force:",omitempty"`
	CreatedById                              string `force:",omitempty"`
	CreatedDate                              string `force:",omitempty"`
	Description                              string `force:",omitempty"`
	Filename                                 string `force:",omitempty"`
	Id                                       string `force:",omitempty"`
	IsDeleted                                bool   `force:",omitempty"`
	LastModifiedById                         string `force:",omitempty"`
	LastModifiedDate                         string `force:",omitempty"`
	LastUsedDate                             string `force:",omitempty"`
	Name                                     string `force:",omitempty"`
	SecurityOptionsAttachmentHasFlash        bool   `force:",omitempty"`
	SecurityOptionsAttachmentHasXSSThreat    bool   `force:",omitempty"`
	SecurityOptionsAttachmentScannedForXSS   bool   `force:",omitempty"`
	SecurityOptionsAttachmentScannedforFlash bool   `force:",omitempty"`
	SystemModstamp                           string `force:",omitempty"`
}

func (t *MailmergeTemplate) ApiName() string {
	return "MailmergeTemplate"
}

func (t *MailmergeTemplate) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("MailmergeTemplate #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tBody: %v\n", t.Body))
	builder.WriteString(fmt.Sprintf("\tBodyLength: %v\n", t.BodyLength))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tDescription: %v\n", t.Description))
	builder.WriteString(fmt.Sprintf("\tFilename: %v\n", t.Filename))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tLastUsedDate: %v\n", t.LastUsedDate))
	builder.WriteString(fmt.Sprintf("\tName: %v\n", t.Name))
	builder.WriteString(fmt.Sprintf("\tSecurityOptionsAttachmentHasFlash: %v\n", t.SecurityOptionsAttachmentHasFlash))
	builder.WriteString(fmt.Sprintf("\tSecurityOptionsAttachmentHasXSSThreat: %v\n", t.SecurityOptionsAttachmentHasXSSThreat))
	builder.WriteString(fmt.Sprintf("\tSecurityOptionsAttachmentScannedForXSS: %v\n", t.SecurityOptionsAttachmentScannedForXSS))
	builder.WriteString(fmt.Sprintf("\tSecurityOptionsAttachmentScannedforFlash: %v\n", t.SecurityOptionsAttachmentScannedforFlash))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))

	return builder.String()
}

type MailmergeTemplateQueryResponse struct {
	BaseQuery
	Records []MailmergeTemplate `json:"Records" force:"records"`
}
