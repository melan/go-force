// This file was generated for SObject CollaborationGroupRecord, API Version v43.0 at 2018-07-30 03:48:07.64202006 -0400 EDT m=+53.986652218

package sobjects

import (
	"fmt"
	"strings"
)

type CollaborationGroupRecord struct {
	BaseSObject
	CollaborationGroupId string `force:",omitempty"`
	CreatedById          string `force:",omitempty"`
	CreatedDate          string `force:",omitempty"`
	Id                   string `force:",omitempty"`
	IsDeleted            bool   `force:",omitempty"`
	LastModifiedById     string `force:",omitempty"`
	LastModifiedDate     string `force:",omitempty"`
	RecordId             string `force:",omitempty"`
	SystemModstamp       string `force:",omitempty"`
}

func (t *CollaborationGroupRecord) ApiName() string {
	return "CollaborationGroupRecord"
}

func (t *CollaborationGroupRecord) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("CollaborationGroupRecord #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCollaborationGroupId: %v\n", t.CollaborationGroupId))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tRecordId: %v\n", t.RecordId))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))

	return builder.String()
}

type CollaborationGroupRecordQueryResponse struct {
	BaseQuery
	Records []CollaborationGroupRecord `json:"Records" force:"records"`
}
