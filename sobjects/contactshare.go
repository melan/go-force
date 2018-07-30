// This file was generated for SObject ContactShare, API Version v43.0 at 2018-07-30 03:47:52.6460754 -0400 EDT m=+38.990144851

package sobjects

import (
	"fmt"
	"strings"
)

type ContactShare struct {
	BaseSObject
	ContactAccessLevel string `force:",omitempty"`
	ContactId          string `force:",omitempty"`
	Id                 string `force:",omitempty"`
	IsDeleted          bool   `force:",omitempty"`
	LastModifiedById   string `force:",omitempty"`
	LastModifiedDate   string `force:",omitempty"`
	RowCause           string `force:",omitempty"`
	UserOrGroupId      string `force:",omitempty"`
}

func (t *ContactShare) ApiName() string {
	return "ContactShare"
}

func (t *ContactShare) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("ContactShare #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tContactAccessLevel: %v\n", t.ContactAccessLevel))
	builder.WriteString(fmt.Sprintf("\tContactId: %v\n", t.ContactId))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tRowCause: %v\n", t.RowCause))
	builder.WriteString(fmt.Sprintf("\tUserOrGroupId: %v\n", t.UserOrGroupId))

	return builder.String()
}

type ContactShareQueryResponse struct {
	BaseQuery
	Records []ContactShare `json:"Records" force:"records"`
}
