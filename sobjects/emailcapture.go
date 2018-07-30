// This file was generated for SObject EmailCapture, API Version v43.0 at 2018-07-30 03:47:44.246316147 -0400 EDT m=+30.590070405

package sobjects

import (
	"fmt"
	"strings"
)

type EmailCapture struct {
	BaseSObject
	CaptureDate      string `force:",omitempty"`
	CreatedById      string `force:",omitempty"`
	CreatedDate      string `force:",omitempty"`
	FromPattern      string `force:",omitempty"`
	Id               string `force:",omitempty"`
	IsActive         bool   `force:",omitempty"`
	IsDeleted        bool   `force:",omitempty"`
	LastModifiedById string `force:",omitempty"`
	LastModifiedDate string `force:",omitempty"`
	RawMessage       string `force:",omitempty"`
	RawMessageLength int    `force:",omitempty"`
	Recipient        string `force:",omitempty"`
	Sender           string `force:",omitempty"`
	SystemModstamp   string `force:",omitempty"`
	ToPattern        string `force:",omitempty"`
}

func (t *EmailCapture) ApiName() string {
	return "EmailCapture"
}

func (t *EmailCapture) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("EmailCapture #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCaptureDate: %v\n", t.CaptureDate))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tFromPattern: %v\n", t.FromPattern))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsActive: %v\n", t.IsActive))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tRawMessage: %v\n", t.RawMessage))
	builder.WriteString(fmt.Sprintf("\tRawMessageLength: %v\n", t.RawMessageLength))
	builder.WriteString(fmt.Sprintf("\tRecipient: %v\n", t.Recipient))
	builder.WriteString(fmt.Sprintf("\tSender: %v\n", t.Sender))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))
	builder.WriteString(fmt.Sprintf("\tToPattern: %v\n", t.ToPattern))

	return builder.String()
}

type EmailCaptureQueryResponse struct {
	BaseQuery
	Records []EmailCapture `json:"Records" force:"records"`
}
