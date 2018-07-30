// This file was generated for SObject AuraDefinitionBundleInfo, API Version v43.0 at 2018-07-30 03:47:38.613033821 -0400 EDT m=+24.956576696

package sobjects

import (
	"fmt"
	"strings"
)

type AuraDefinitionBundleInfo struct {
	BaseSObject
	ApiVersion             float64 `force:",omitempty"`
	AuraDefinitionBundleId string  `force:",omitempty"`
	DeveloperName          string  `force:",omitempty"`
	DurableId              string  `force:",omitempty"`
	Id                     string  `force:",omitempty"`
	NamespacePrefix        string  `force:",omitempty"`
}

func (t *AuraDefinitionBundleInfo) ApiName() string {
	return "AuraDefinitionBundleInfo"
}

func (t *AuraDefinitionBundleInfo) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("AuraDefinitionBundleInfo #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tApiVersion: %v\n", t.ApiVersion))
	builder.WriteString(fmt.Sprintf("\tAuraDefinitionBundleId: %v\n", t.AuraDefinitionBundleId))
	builder.WriteString(fmt.Sprintf("\tDeveloperName: %v\n", t.DeveloperName))
	builder.WriteString(fmt.Sprintf("\tDurableId: %v\n", t.DurableId))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tNamespacePrefix: %v\n", t.NamespacePrefix))

	return builder.String()
}

type AuraDefinitionBundleInfoQueryResponse struct {
	BaseQuery
	Records []AuraDefinitionBundleInfo `json:"Records" force:"records"`
}
