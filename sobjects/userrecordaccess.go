// This file was generated for SObject UserRecordAccess, API Version v43.0 at 2018-07-30 03:47:37.952350631 -0400 EDT m=+24.295868715

package sobjects

import (
	"fmt"
	"strings"
)

type UserRecordAccess struct {
	BaseSObject
	HasAllAccess      bool   `force:",omitempty"`
	HasDeleteAccess   bool   `force:",omitempty"`
	HasEditAccess     bool   `force:",omitempty"`
	HasReadAccess     bool   `force:",omitempty"`
	HasTransferAccess bool   `force:",omitempty"`
	Id                string `force:",omitempty"`
	MaxAccessLevel    string `force:",omitempty"`
	RecordId          string `force:",omitempty"`
	UserId            string `force:",omitempty"`
}

func (t *UserRecordAccess) ApiName() string {
	return "UserRecordAccess"
}

func (t *UserRecordAccess) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("UserRecordAccess #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tHasAllAccess: %v\n", t.HasAllAccess))
	builder.WriteString(fmt.Sprintf("\tHasDeleteAccess: %v\n", t.HasDeleteAccess))
	builder.WriteString(fmt.Sprintf("\tHasEditAccess: %v\n", t.HasEditAccess))
	builder.WriteString(fmt.Sprintf("\tHasReadAccess: %v\n", t.HasReadAccess))
	builder.WriteString(fmt.Sprintf("\tHasTransferAccess: %v\n", t.HasTransferAccess))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tMaxAccessLevel: %v\n", t.MaxAccessLevel))
	builder.WriteString(fmt.Sprintf("\tRecordId: %v\n", t.RecordId))
	builder.WriteString(fmt.Sprintf("\tUserId: %v\n", t.UserId))

	return builder.String()
}

type UserRecordAccessQueryResponse struct {
	BaseQuery
	Records []UserRecordAccess `json:"Records" force:"records"`
}
