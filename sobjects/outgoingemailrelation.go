// This file was generated for SObject OutgoingEmailRelation, API Version v43.0 at 2018-07-30 03:47:58.932027627 -0400 EDT m=+45.276332952

package sobjects

import (
	"fmt"
	"strings"
)

type OutgoingEmailRelation struct {
	BaseSObject
	ExternalId      string `force:",omitempty"`
	Id              string `force:",omitempty"`
	OutgoingEmailId string `force:",omitempty"`
	RelationAddress string `force:",omitempty"`
	RelationId      string `force:",omitempty"`
}

func (t *OutgoingEmailRelation) ApiName() string {
	return "OutgoingEmailRelation"
}

func (t *OutgoingEmailRelation) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("OutgoingEmailRelation #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tExternalId: %v\n", t.ExternalId))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tOutgoingEmailId: %v\n", t.OutgoingEmailId))
	builder.WriteString(fmt.Sprintf("\tRelationAddress: %v\n", t.RelationAddress))
	builder.WriteString(fmt.Sprintf("\tRelationId: %v\n", t.RelationId))

	return builder.String()
}

type OutgoingEmailRelationQueryResponse struct {
	BaseQuery
	Records []OutgoingEmailRelation `json:"Records" force:"records"`
}
