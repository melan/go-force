// This file was generated for SObject RelationshipInfo, API Version v43.0 at 2018-07-30 03:48:04.421676602 -0400 EDT m=+50.766187920

package sobjects

import (
	"fmt"
	"strings"
)

type RelationshipInfo struct {
	BaseSObject
	ChildSobjectId        string `force:",omitempty"`
	DurableId             string `force:",omitempty"`
	FieldId               string `force:",omitempty"`
	Id                    string `force:",omitempty"`
	IsCascadeDelete       bool   `force:",omitempty"`
	IsDeprecatedAndHidden bool   `force:",omitempty"`
	IsRestrictedDelete    bool   `force:",omitempty"`
	JunctionIdListNames   string `force:",omitempty"`
}

func (t *RelationshipInfo) ApiName() string {
	return "RelationshipInfo"
}

func (t *RelationshipInfo) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("RelationshipInfo #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tChildSobjectId: %v\n", t.ChildSobjectId))
	builder.WriteString(fmt.Sprintf("\tDurableId: %v\n", t.DurableId))
	builder.WriteString(fmt.Sprintf("\tFieldId: %v\n", t.FieldId))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsCascadeDelete: %v\n", t.IsCascadeDelete))
	builder.WriteString(fmt.Sprintf("\tIsDeprecatedAndHidden: %v\n", t.IsDeprecatedAndHidden))
	builder.WriteString(fmt.Sprintf("\tIsRestrictedDelete: %v\n", t.IsRestrictedDelete))
	builder.WriteString(fmt.Sprintf("\tJunctionIdListNames: %v\n", t.JunctionIdListNames))

	return builder.String()
}

type RelationshipInfoQueryResponse struct {
	BaseQuery
	Records []RelationshipInfo `json:"Records" force:"records"`
}
