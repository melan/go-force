// This file was generated for SObject DuplicateRecordItem, API Version v43.0 at 2018-07-30 03:47:26.595040228 -0400 EDT m=+12.938132141

package sobjects

import (
	"fmt"
	"strings"
)

type DuplicateRecordItem struct {
	BaseSObject
	CreatedById          string `force:",omitempty"`
	CreatedDate          string `force:",omitempty"`
	DuplicateRecordSetId string `force:",omitempty"`
	Id                   string `force:",omitempty"`
	IsDeleted            bool   `force:",omitempty"`
	LastModifiedById     string `force:",omitempty"`
	LastModifiedDate     string `force:",omitempty"`
	Name                 string `force:",omitempty"`
	RecordId             string `force:",omitempty"`
	SystemModstamp       string `force:",omitempty"`
}

func (t *DuplicateRecordItem) ApiName() string {
	return "DuplicateRecordItem"
}

func (t *DuplicateRecordItem) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("DuplicateRecordItem #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tDuplicateRecordSetId: %v\n", t.DuplicateRecordSetId))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tName: %v\n", t.Name))
	builder.WriteString(fmt.Sprintf("\tRecordId: %v\n", t.RecordId))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))

	return builder.String()
}

type DuplicateRecordItemQueryResponse struct {
	BaseQuery
	Records []DuplicateRecordItem `json:"Records" force:"records"`
}
