// This file was generated for SObject UndecidedEventRelation, API Version v43.0 at 2018-07-30 03:47:59.569882562 -0400 EDT m=+45.914211821

package sobjects

import (
	"fmt"
	"strings"
)

type UndecidedEventRelation struct {
	BaseSObject
	CreatedById      string `force:",omitempty"`
	CreatedDate      string `force:",omitempty"`
	EventId          string `force:",omitempty"`
	Id               string `force:",omitempty"`
	IsDeleted        bool   `force:",omitempty"`
	LastModifiedById string `force:",omitempty"`
	LastModifiedDate string `force:",omitempty"`
	RelationId       string `force:",omitempty"`
	RespondedDate    string `force:",omitempty"`
	Response         string `force:",omitempty"`
	SystemModstamp   string `force:",omitempty"`
	Type             string `force:",omitempty"`
}

func (t *UndecidedEventRelation) ApiName() string {
	return "UndecidedEventRelation"
}

func (t *UndecidedEventRelation) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("UndecidedEventRelation #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tEventId: %v\n", t.EventId))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tRelationId: %v\n", t.RelationId))
	builder.WriteString(fmt.Sprintf("\tRespondedDate: %v\n", t.RespondedDate))
	builder.WriteString(fmt.Sprintf("\tResponse: %v\n", t.Response))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))
	builder.WriteString(fmt.Sprintf("\tType: %v\n", t.Type))

	return builder.String()
}

type UndecidedEventRelationQueryResponse struct {
	BaseQuery
	Records []UndecidedEventRelation `json:"Records" force:"records"`
}
