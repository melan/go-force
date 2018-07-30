// This file was generated for SObject EmailStatus, API Version v43.0 at 2018-07-30 03:47:16.252557327 -0400 EDT m=+2.595261146

package sobjects

import (
	"fmt"
	"strings"
)

type EmailStatus struct {
	BaseSObject
	CreatedById       string `force:",omitempty"`
	CreatedDate       string `force:",omitempty"`
	EmailTemplateName string `force:",omitempty"`
	FirstOpenDate     string `force:",omitempty"`
	Id                string `force:",omitempty"`
	LastModifiedById  string `force:",omitempty"`
	LastModifiedDate  string `force:",omitempty"`
	LastOpenDate      string `force:",omitempty"`
	TaskId            string `force:",omitempty"`
	TimesOpened       int    `force:",omitempty"`
	WhoId             string `force:",omitempty"`
}

func (t *EmailStatus) ApiName() string {
	return "EmailStatus"
}

func (t *EmailStatus) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("EmailStatus #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tEmailTemplateName: %v\n", t.EmailTemplateName))
	builder.WriteString(fmt.Sprintf("\tFirstOpenDate: %v\n", t.FirstOpenDate))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tLastOpenDate: %v\n", t.LastOpenDate))
	builder.WriteString(fmt.Sprintf("\tTaskId: %v\n", t.TaskId))
	builder.WriteString(fmt.Sprintf("\tTimesOpened: %v\n", t.TimesOpened))
	builder.WriteString(fmt.Sprintf("\tWhoId: %v\n", t.WhoId))

	return builder.String()
}

type EmailStatusQueryResponse struct {
	BaseQuery
	Records []EmailStatus `json:"Records" force:"records"`
}
