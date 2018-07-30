// This file was generated for SObject LightningComponentBundle, API Version v43.0 at 2018-07-30 03:47:33.746828357 -0400 EDT m=+20.090188633

package sobjects

import (
	"fmt"
	"strings"
)

type LightningComponentBundle struct {
	BaseSObject
	ApiVersion       float64 `force:",omitempty"`
	AvailableFor     string  `force:",omitempty"`
	CreatedById      string  `force:",omitempty"`
	CreatedDate      string  `force:",omitempty"`
	Description      string  `force:",omitempty"`
	DeveloperName    string  `force:",omitempty"`
	Id               string  `force:",omitempty"`
	IsDeleted        bool    `force:",omitempty"`
	IsExposed        bool    `force:",omitempty"`
	Language         string  `force:",omitempty"`
	LastModifiedById string  `force:",omitempty"`
	LastModifiedDate string  `force:",omitempty"`
	MasterLabel      string  `force:",omitempty"`
	MinVersion       float64 `force:",omitempty"`
	NamespacePrefix  string  `force:",omitempty"`
	SystemModstamp   string  `force:",omitempty"`
	TagConfigs       string  `force:",omitempty"`
}

func (t *LightningComponentBundle) ApiName() string {
	return "LightningComponentBundle"
}

func (t *LightningComponentBundle) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("LightningComponentBundle #%s - %s\n", t.Id, t.Name))
	builder.WriteString(fmt.Sprintf("\tApiVersion: %v\n", t.ApiVersion))
	builder.WriteString(fmt.Sprintf("\tAvailableFor: %v\n", t.AvailableFor))
	builder.WriteString(fmt.Sprintf("\tCreatedById: %v\n", t.CreatedById))
	builder.WriteString(fmt.Sprintf("\tCreatedDate: %v\n", t.CreatedDate))
	builder.WriteString(fmt.Sprintf("\tDescription: %v\n", t.Description))
	builder.WriteString(fmt.Sprintf("\tDeveloperName: %v\n", t.DeveloperName))
	builder.WriteString(fmt.Sprintf("\tId: %v\n", t.Id))
	builder.WriteString(fmt.Sprintf("\tIsDeleted: %v\n", t.IsDeleted))
	builder.WriteString(fmt.Sprintf("\tIsExposed: %v\n", t.IsExposed))
	builder.WriteString(fmt.Sprintf("\tLanguage: %v\n", t.Language))
	builder.WriteString(fmt.Sprintf("\tLastModifiedById: %v\n", t.LastModifiedById))
	builder.WriteString(fmt.Sprintf("\tLastModifiedDate: %v\n", t.LastModifiedDate))
	builder.WriteString(fmt.Sprintf("\tMasterLabel: %v\n", t.MasterLabel))
	builder.WriteString(fmt.Sprintf("\tMinVersion: %v\n", t.MinVersion))
	builder.WriteString(fmt.Sprintf("\tNamespacePrefix: %v\n", t.NamespacePrefix))
	builder.WriteString(fmt.Sprintf("\tSystemModstamp: %v\n", t.SystemModstamp))
	builder.WriteString(fmt.Sprintf("\tTagConfigs: %v\n", t.TagConfigs))

	return builder.String()
}

type LightningComponentBundleQueryResponse struct {
	BaseQuery
	Records []LightningComponentBundle `json:"Records" force:"records"`
}
