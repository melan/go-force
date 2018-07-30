// This file was generated for SObject AuraDefinitionInfo, API Version v43.0 at 2018-07-30 03:47:18.187255448 -0400 EDT m=+4.530031865

package sobjects

import (
	"fmt"
	"strings"
)

type AuraDefinitionInfo struct {
	BaseSObject
	AuraDefinitionBundleInfoId string `force:",omitempty"`
	AuraDefinitionId           string `force:",omitempty"`
	DefType                    string `force:",omitempty"`
	DeveloperName              string `force:",omitempty"`
	DurableId                  string `force:",omitempty"`
	Format                     string `force:",omitempty"`
	Id                         string `force:",omitempty"`
	LastModifiedDate           string `force:",omitempty"`
	NamespacePrefix            string `force:",omitempty"`
	Source                     string `force:",omitempty"`
}

func (t *AuraDefinitionInfo) ApiName() string {
	return "AuraDefinitionInfo"
}

func (t *AuraDefinitionInfo) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("AuraDefinitionInfo #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tAuraDefinitionBundleInfoId: %v\n", t.AuraDefinitionBundleInfoId))
	builder.WriteString(fmt.Sprintf("\tAuraDefinitionId: %v\n", t.AuraDefinitionId))
	builder.WriteString(fmt.Sprintf("\tDefType: %v\n", t.DefType))
	builder.WriteString(fmt.Sprintf("\tDeveloperName: %v\n", t.DeveloperName))
	builder.WriteString(fmt.Sprintf("\tDurableId: %v\n", t.DurableId))
	builder.WriteString(fmt.Sprintf("\tFormat: %v\n", t.Format))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tNamespacePrefix: %v\n", t.NamespacePrefix))
	builder.WriteString(fmt.Sprintf("\tSource: %v\n", t.Source))

	return builder.String()
}

type AuraDefinitionInfoQueryResponse struct {
	BaseQuery
	Records []AuraDefinitionInfo `json:"Records" force:"records"`
}
