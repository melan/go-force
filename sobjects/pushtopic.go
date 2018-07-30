// This file was generated for SObject PushTopic, API Version v43.0 at 2018-07-30 03:47:27.564160397 -0400 EDT m=+13.907288675

package sobjects

import (
	"fmt"
	"strings"
)

type PushTopic struct {
	BaseSObject
	ApiVersion                 float64 `force:",omitempty"`
	CreatedById                string  `force:",omitempty"`
	CreatedDate                string  `force:",omitempty"`
	Description                string  `force:",omitempty"`
	Id                         string  `force:",omitempty"`
	IsActive                   bool    `force:",omitempty"`
	IsDeleted                  bool    `force:",omitempty"`
	LastModifiedById           string  `force:",omitempty"`
	LastModifiedDate           string  `force:",omitempty"`
	Name                       string  `force:",omitempty"`
	NotifyForFields            string  `force:",omitempty"`
	NotifyForOperationCreate   bool    `force:",omitempty"`
	NotifyForOperationDelete   bool    `force:",omitempty"`
	NotifyForOperationUndelete bool    `force:",omitempty"`
	NotifyForOperationUpdate   bool    `force:",omitempty"`
	NotifyForOperations        string  `force:",omitempty"`
	Query                      string  `force:",omitempty"`
	SystemModstamp             string  `force:",omitempty"`
}

func (t *PushTopic) ApiName() string {
	return "PushTopic"
}

func (t *PushTopic) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("PushTopic #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tApiVersion: %v\n", t.ApiVersion))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tDescription: %v\n", t.Description))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsActive: %v\n", t.IsActive))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tName: %v\n", t.Name))
	builder.WriteString(fmt.Sprintf("\tNotifyForFields: %v\n", t.NotifyForFields))
	builder.WriteString(fmt.Sprintf("\tNotifyForOperationCreate: %v\n", t.NotifyForOperationCreate))
	builder.WriteString(fmt.Sprintf("\tNotifyForOperationDelete: %v\n", t.NotifyForOperationDelete))
	builder.WriteString(fmt.Sprintf("\tNotifyForOperationUndelete: %v\n", t.NotifyForOperationUndelete))
	builder.WriteString(fmt.Sprintf("\tNotifyForOperationUpdate: %v\n", t.NotifyForOperationUpdate))
	builder.WriteString(fmt.Sprintf("\tNotifyForOperations: %v\n", t.NotifyForOperations))
	builder.WriteString(fmt.Sprintf("\tQuery: %v\n", t.Query))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))

	return builder.String()
}

type PushTopicQueryResponse struct {
	BaseQuery
	Records []PushTopic `json:"Records" force:"records"`
}
