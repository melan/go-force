// This file was generated for SObject RelationshipDomain, API Version v43.0 at 2018-07-30 03:47:46.482045464 -0400 EDT m=+32.825883616

package sobjects

import (
	"fmt"
	"strings"
)

type RelationshipDomain struct {
	BaseSObject
	ChildSobjectId        string `force:",omitempty"`
	DurableId             string `force:",omitempty"`
	FieldId               string `force:",omitempty"`
	Id                    string `force:",omitempty"`
	IsCascadeDelete       bool   `force:",omitempty"`
	IsDeprecatedAndHidden bool   `force:",omitempty"`
	IsRestrictedDelete    bool   `force:",omitempty"`
	JunctionIdListNames   string `force:",omitempty"`
	ParentSobjectId       string `force:",omitempty"`
	RelationshipInfoId    string `force:",omitempty"`
	RelationshipName      string `force:",omitempty"`
}

func (t *RelationshipDomain) ApiName() string {
	return "RelationshipDomain"
}

func (t *RelationshipDomain) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("RelationshipDomain #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tChildSobjectId: %v\n", t.ChildSobjectId))
	builder.WriteString(fmt.Sprintf("\tDurableId: %v\n", t.DurableId))
	builder.WriteString(fmt.Sprintf("\tFieldId: %v\n", t.FieldId))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsCascadeDelete: %v\n", t.IsCascadeDelete))
	builder.WriteString(fmt.Sprintf("\tIsDeprecatedAndHidden: %v\n", t.IsDeprecatedAndHidden))
	builder.WriteString(fmt.Sprintf("\tIsRestrictedDelete: %v\n", t.IsRestrictedDelete))
	builder.WriteString(fmt.Sprintf("\tJunctionIdListNames: %v\n", t.JunctionIdListNames))
	builder.WriteString(fmt.Sprintf("\tParentSobjectId: %v\n", t.ParentSobjectId))
	builder.WriteString(fmt.Sprintf("\tRelationshipInfoId: %v\n", t.RelationshipInfoId))
	builder.WriteString(fmt.Sprintf("\tRelationshipName: %v\n", t.RelationshipName))

	return builder.String()
}

type RelationshipDomainQueryResponse struct {
	BaseQuery
	Records []RelationshipDomain `json:"Records" force:"records"`
}
