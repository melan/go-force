// This file was generated for SObject DuplicateRecordSet, API Version v43.0 at 2018-07-30 03:47:32.341147216 -0400 EDT m=+18.684454745

package sobjects

import (
	"fmt"
	"strings"
)

type DuplicateRecordSet struct {
	BaseSObject
	CreatedById        string `force:",omitempty"`
	CreatedDate        string `force:",omitempty"`
	DuplicateRuleId    string `force:",omitempty"`
	Id                 string `force:",omitempty"`
	IsDeleted          bool   `force:",omitempty"`
	LastModifiedById   string `force:",omitempty"`
	LastModifiedDate   string `force:",omitempty"`
	LastReferencedDate string `force:",omitempty"`
	LastViewedDate     string `force:",omitempty"`
	Name               string `force:",omitempty"`
	RecordCount        int    `force:",omitempty"`
	SystemModstamp     string `force:",omitempty"`
}

func (t *DuplicateRecordSet) ApiName() string {
	return "DuplicateRecordSet"
}

func (t *DuplicateRecordSet) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("DuplicateRecordSet #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tDuplicateRuleId: %v\n", t.DuplicateRuleId))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tLastReferencedDate: %v\n", t.LastReferencedDate))
	builder.WriteString(fmt.Sprintf("\tLastViewedDate: %v\n", t.LastViewedDate))
	builder.WriteString(fmt.Sprintf("\tName: %v\n", t.Name))
	builder.WriteString(fmt.Sprintf("\tRecordCount: %v\n", t.RecordCount))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))

	return builder.String()
}

type DuplicateRecordSetQueryResponse struct {
	BaseQuery
	Records []DuplicateRecordSet `json:"Records" force:"records"`
}
