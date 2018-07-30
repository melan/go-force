// This file was generated for SObject PicklistValueInfo, API Version v43.0 at 2018-07-30 03:47:33.116861558 -0400 EDT m=+19.460198196

package sobjects

import (
	"fmt"
	"strings"
)

type PicklistValueInfo struct {
	BaseSObject
	DurableId        string `force:",omitempty"`
	EntityParticleId string `force:",omitempty"`
	Id               string `force:",omitempty"`
	IsActive         bool   `force:",omitempty"`
	IsDefaultValue   bool   `force:",omitempty"`
	Label            string `force:",omitempty"`
	ValidFor         string `force:",omitempty"`
	Value            string `force:",omitempty"`
}

func (t *PicklistValueInfo) ApiName() string {
	return "PicklistValueInfo"
}

func (t *PicklistValueInfo) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("PicklistValueInfo #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tDurableId: %v\n", t.DurableId))
	builder.WriteString(fmt.Sprintf("\tEntityParticleId: %v\n", t.EntityParticleId))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsActive: %v\n", t.IsActive))
	builder.WriteString(fmt.Sprintf("\tIsDefaultValue: %v\n", t.IsDefaultValue))
	builder.WriteString(fmt.Sprintf("\tLabel: %v\n", t.Label))
	builder.WriteString(fmt.Sprintf("\tValidFor: %v\n", t.ValidFor))
	builder.WriteString(fmt.Sprintf("\tValue: %v\n", t.Value))

	return builder.String()
}

type PicklistValueInfoQueryResponse struct {
	BaseQuery
	Records []PicklistValueInfo `json:"Records" force:"records"`
}
