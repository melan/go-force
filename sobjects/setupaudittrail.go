// This file was generated for SObject SetupAuditTrail, API Version v43.0 at 2018-07-30 03:47:59.066989011 -0400 EDT m=+45.411299400

package sobjects

import (
	"fmt"
	"strings"
)

type SetupAuditTrail struct {
	BaseSObject
	Action                     string `force:",omitempty"`
	CreatedById                string `force:",omitempty"`
	CreatedDate                string `force:",omitempty"`
	DelegateUser               string `force:",omitempty"`
	Display                    string `force:",omitempty"`
	Id                         string `force:",omitempty"`
	ResponsibleNamespacePrefix string `force:",omitempty"`
	Section                    string `force:",omitempty"`
}

func (t *SetupAuditTrail) ApiName() string {
	return "SetupAuditTrail"
}

func (t *SetupAuditTrail) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("SetupAuditTrail #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tAction: %v\n", t.Action))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tDelegateUser: %v\n", t.DelegateUser))
	builder.WriteString(fmt.Sprintf("\tDisplay: %v\n", t.Display))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tResponsibleNamespacePrefix: %v\n", t.ResponsibleNamespacePrefix))
	builder.WriteString(fmt.Sprintf("\tSection: %v\n", t.Section))

	return builder.String()
}

type SetupAuditTrailQueryResponse struct {
	BaseQuery
	Records []SetupAuditTrail `json:"Records" force:"records"`
}
