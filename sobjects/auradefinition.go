// This file was generated for SObject AuraDefinition, API Version v43.0 at 2018-07-30 03:47:21.24282258 -0400 EDT m=+7.585713655

package sobjects

import (
	"fmt"
	"strings"
)

type AuraDefinition struct {
	BaseSObject
	AuraDefinitionBundleId string `force:",omitempty"`
	CreatedById            string `force:",omitempty"`
	CreatedDate            string `force:",omitempty"`
	DefType                string `force:",omitempty"`
	Format                 string `force:",omitempty"`
	Id                     string `force:",omitempty"`
	IsDeleted              bool   `force:",omitempty"`
	LastModifiedById       string `force:",omitempty"`
	LastModifiedDate       string `force:",omitempty"`
	Source                 string `force:",omitempty"`
	SystemModstamp         string `force:",omitempty"`
}

func (t *AuraDefinition) ApiName() string {
	return "AuraDefinition"
}

func (t *AuraDefinition) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("AuraDefinition #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tAuraDefinitionBundleId: %v\n", t.AuraDefinitionBundleId))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tDefType: %v\n", t.DefType))
	builder.WriteString(fmt.Sprintf("\tFormat: %v\n", t.Format))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tSource: %v\n", t.Source))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))

	return builder.String()
}

type AuraDefinitionQueryResponse struct {
	BaseQuery
	Records []AuraDefinition `json:"Records" force:"records"`
}
