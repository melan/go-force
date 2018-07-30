// This file was generated for SObject EmailMessageRelation, API Version v43.0 at 2018-07-30 03:47:56.938442736 -0400 EDT m=+43.282673253

package sobjects

import (
	"fmt"
	"strings"
)

type EmailMessageRelation struct {
	BaseSObject
	CreatedById        string `force:",omitempty"`
	CreatedDate        string `force:",omitempty"`
	EmailMessageId     string `force:",omitempty"`
	Id                 string `force:",omitempty"`
	IsDeleted          bool   `force:",omitempty"`
	RelationAddress    string `force:",omitempty"`
	RelationId         string `force:",omitempty"`
	RelationObjectType string `force:",omitempty"`
	RelationType       string `force:",omitempty"`
	SystemModstamp     string `force:",omitempty"`
}

func (t *EmailMessageRelation) ApiName() string {
	return "EmailMessageRelation"
}

func (t *EmailMessageRelation) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("EmailMessageRelation #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tEmailMessageId: %v\n", t.EmailMessageId))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tRelationAddress: %v\n", t.RelationAddress))
	builder.WriteString(fmt.Sprintf("\tRelationId: %v\n", t.RelationId))
	builder.WriteString(fmt.Sprintf("\tRelationObjectType: %v\n", t.RelationObjectType))
	builder.WriteString(fmt.Sprintf("\tRelationType: %v\n", t.RelationType))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))

	return builder.String()
}

type EmailMessageRelationQueryResponse struct {
	BaseQuery
	Records []EmailMessageRelation `json:"Records" force:"records"`
}
