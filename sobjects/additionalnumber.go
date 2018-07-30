// This file was generated for SObject AdditionalNumber, API Version v43.0 at 2018-07-30 03:47:15.766751157 -0400 EDT m=+2.109436747

package sobjects

import (
	"fmt"
	"strings"
)

type AdditionalNumber struct {
	BaseSObject
	CallCenterId     string `force:",omitempty"`
	CreatedById      string `force:",omitempty"`
	CreatedDate      string `force:",omitempty"`
	Description      string `force:",omitempty"`
	Id               string `force:",omitempty"`
	IsDeleted        bool   `force:",omitempty"`
	LastModifiedById string `force:",omitempty"`
	LastModifiedDate string `force:",omitempty"`
	Name             string `force:",omitempty"`
	Phone            string `force:",omitempty"`
	SystemModstamp   string `force:",omitempty"`
}

func (t *AdditionalNumber) ApiName() string {
	return "AdditionalNumber"
}

func (t *AdditionalNumber) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("AdditionalNumber #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCallCenterId: %v\n", t.CallCenterId))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tDescription: %v\n", t.Description))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tName: %v\n", t.Name))
	builder.WriteString(fmt.Sprintf("\tPhone: %v\n", t.Phone))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))

	return builder.String()
}

type AdditionalNumberQueryResponse struct {
	BaseQuery
	Records []AdditionalNumber `json:"Records" force:"records"`
}
